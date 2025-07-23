package main // main 包，程序的入口点

import (
	"context"      // 用于操作上下文，控制请求的生命周期
	"crypto/ecdsa" // 用于处理 ECDSA 私钥
	"fmt"          // 格式化输入输出
	"log"          // 记录日志
	"math/big"     // 处理大整数，例如以太坊中的 nonce 和计数
	"os"           // 操作系统接口，用于读取环境变量
	"strings"      // 字符串操作，用于 ABI 解析
	"time"         // 时间相关操作，用于暂停

	ethereum "github.com/ethereum/go-ethereum"          // 以太坊核心类型和接口，如 FilterQuery
	"github.com/ethereum/go-ethereum/accounts/abi"      // 处理合约 ABI
	"github.com/ethereum/go-ethereum/accounts/abi/bind" // Go 语言绑定以太坊合约
	"github.com/ethereum/go-ethereum/common"            // 常用以太坊数据类型，如地址、哈希
	"github.com/ethereum/go-ethereum/core/types"        // 以太坊核心类型，如交易、收据
	"github.com/ethereum/go-ethereum/crypto"            // 加密操作，如私钥推导地址
	"github.com/ethereum/go-ethereum/ethclient"         // 以太坊 RPC 客户端
	"github.com/joho/godotenv"                          // 用于从 .env 文件加载环境变量
)

func main() {
	// 1. 加载 .env 环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("错误：无法加载 .env 文件：%v", err)
	}

	// 2. 从环境变量获取配置信息
	privateKeyHex := os.Getenv("PRIVATE_KEY")                   // 账户私钥（十六进制字符串）
	infuraURL := os.Getenv("INFURA_URL")                        // 以太坊节点 URL (通常是 WSS 地址用于订阅)
	contractAddressHex := os.Getenv("CONTRACT_ADDRESS_COUNTER") // 已部署的计数器合约地址

	// 检查必要环境变量是否已设置
	if privateKeyHex == "" || infuraURL == "" || contractAddressHex == "" {
		log.Fatal("错误：请设置 PRIVATE_KEY, INFURA_URL 和 CONTRACT_ADDRESS_COUNTER 环境变量")
	}

	// 3. 连接到以太坊网络节点
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("错误：无法连接到以太坊网络：%v", err)
	}
	defer client.Close() // 确保在 main 函数结束时关闭客户端连接

	// 4. 解析私钥并获取发送方地址
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("错误：无效的私钥：%v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法断言类型: publicKey 不是 *ecdsa.PublicKey 类型")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) // 从公钥推导出以太坊地址
	fmt.Printf("发件人地址：%s\n", fromAddress.Hex())

	// 5. 获取当前连接网络的链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("错误：无法获取链ID：%v", err)
	}

	// 6. 实例化合约绑定对象
	contractAddress := common.HexToAddress(contractAddressHex)
	// 假设 NewTask2Contracts 函数在当前 main 包中可用
	instance, err := NewTask2Contracts(contractAddress, client)
	if err != nil {
		log.Fatalf("错误：无法实例化合约：%v", err)
	}
	fmt.Printf("已连接到合约地址：%s\n", contractAddress.Hex())

	// 7. 启动事件监听 Goroutine
	eventCtx, cancelEvents := context.WithCancel(context.Background())
	defer cancelEvents() // 确保在 main 函数退出时取消事件监听
	// 假设 Task2Contracts 是 abigen 生成的合约结构体名称
	go listenForCountChangedEvents(eventCtx, client, contractAddress, instance)

	// 8. 查询当前计数器的值 (调用合约的视图函数)
	count, err := instance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("错误：无法调用 getCount：%v", err)
	}
	fmt.Printf("当前计数：%s\n", count.String())

	// 9. 构造交易发送器 (TransactOpts)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("错误：无法创建交易发送器：%v", err)
	}
	auth.GasLimit = uint64(100000) // 设置交易的 Gas 上限

	// 获取账户的最新 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("错误：无法获取账户 nonce：%v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce)) // 设置 increment 交易的 nonce

	// 10. 发送 increment 交易 (增加计数)
	fmt.Println("\n发送 increment 交易...")
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatalf("错误：无法发送 increment 交易：%v", err)
	}
	fmt.Printf("Increment 交易哈希：%s\n", tx.Hash().Hex())

	// 11. 等待 increment 交易在链上被确认
	fmt.Println("正在等待 Increment 交易确认...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("错误：等待 Increment 交易确认失败：%v", err)
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Increment 交易成功确认！")
	} else {
		fmt.Println("Increment 交易失败！")
	}

	// 12. 再次查询计数以确认变更
	time.Sleep(2 * time.Second) // 稍作等待，确保节点同步
	updatedCount, err := instance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("错误：无法再次调用 getCount：%v", err)
	}
	fmt.Printf("Increment 后最新计数：%s\n", updatedCount.String())

	// 13. 调用 decrement 交易 (减少计数)
	auth.Nonce = big.NewInt(int64(nonce + 1)) // nonce 在每次发送交易后需要递增
	fmt.Println("\n发送 decrement 交易...")
	decrementTx, err := instance.Decrement(auth)
	if err != nil {
		log.Fatalf("错误：无法发送 decrement 交易：%v", err)
	}
	fmt.Printf("Decrement 交易哈希：%s\n", decrementTx.Hash().Hex())

	// 14. 等待 decrement 交易在链上被确认
	fmt.Println("正在等待 Decrement 交易确认...")
	decrementReceipt, err := bind.WaitMined(context.Background(), client, decrementTx)
	if err != nil {
		log.Fatalf("错误：等待 decrement 交易确认失败：%v", err)
	}
	if decrementReceipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Decrement 交易成功确认！")
	} else {
		fmt.Println("Decrement 交易失败！")
	}

	// 15. 最终查询计数以确认所有变更
	time.Sleep(2 * time.Second) // 稍作等待
	finalCount, err := instance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("错误：无法获取最终计数：%v", err)
	}
	fmt.Printf("Decrement 后最终计数：%s\n", finalCount.String())

	// 保持 main 函数运行，以便事件监听协程持续工作
	fmt.Println("\n程序运行中，监听合约事件... (按 Ctrl+C 退出)")
	select {} // 阻塞 main goroutine，防止程序退出
}

// listenForCountChangedEvents 函数：持续监听 CountChanged 事件
// ctx: 用于控制协程生命周期
// client: 以太坊客户端
// contractAddress: 要监听的合约地址
// instance: 合约绑定实例，用于解析事件日志
func listenForCountChangedEvents(ctx context.Context, client *ethclient.Client, contractAddress common.Address, instance *Task2Contracts) {
	logs := make(chan types.Log)

	// 从 Task2ContractsMetaData.ABI 动态解析 CountChanged 事件的主题 (Topic)
	// 假设 Task2ContractsMetaData 在同一个 package main 下
	parsedABI, err := abi.JSON(strings.NewReader(Task2ContractsMetaData.ABI))
	if err != nil {
		log.Printf("错误：无法解析 ABI：%v", err)
		return
	}
	eventSig := parsedABI.Events["CountChanged"].ID // 获取 CountChanged 事件的签名哈希

	// 构建事件过滤查询，只监听指定合约和 CountChanged 事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventSig}},
	}

	// 订阅事件日志流
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Printf("错误：无法订阅 CountChanged 事件日志：%v", err)
		return
	}
	defer sub.Unsubscribe() // 确保在函数退出时取消订阅

	fmt.Println("开始监听 CountChanged 事件...")
	for {
		select {
		case <-ctx.Done(): // 接收到取消信号时退出
			fmt.Println("事件监听协程已停止。")
			return
		case err := <-sub.Err(): // 订阅错误时退出
			log.Printf("事件订阅错误：%v", err)
			return
		case vLog := <-logs: // 接收到新的事件日志
			// 使用合约绑定实例的 Filterer 方法来解析原始日志
			event, err := instance.Task2ContractsFilterer.ParseCountChanged(vLog)
			if err != nil {
				log.Printf("错误：无法解析事件日志：%v", err)
				continue
			}

			// 打印解析后的事件详情，包括中文消息
			fmt.Printf("收到 CountChanged 事件: 新计数 = %s, 变更者 = %s, 消息 = \"%s\", 区块号 = %d\n",
				event.NewCount.String(), event.Changer.Hex(), event.Message, vLog.BlockNumber)
		}
	}
}
