package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("加载 .env 文件失败: %v", err)
	}

	// 从环境变量获取配置
	infuraURL := os.Getenv("INFURA_URL")
	privateKeyStr := os.Getenv("PRIVATE_KEY")
	recipientAddressStr := os.Getenv("RECIPIENT_ADDRESS")

	if infuraURL == "" || privateKeyStr == "" || recipientAddressStr == "" {
		log.Fatal("环境变量 INFURA_URL, PRIVATE_KEY 或 RECIPIENT_ADDRESS 未设置")
	}

	// 连接到 Sepolia 测试网络
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer client.Close()

	// 任务 1.1：查询区块信息
	// 动态获取最新区块号
	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("获取最新区块号失败: %v", err)
	}
	blockNumber := big.NewInt(int64(latestBlockNumber)) // 使用最新区块号

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("查询区块失败: %v", err)
	}

	fmt.Printf("区块号: %d\n", block.Number().Uint64())
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("时间戳: %d\n", block.Time())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))

	// 任务 1.2：发送交易
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥转换为 ECDSA 失败")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}

	value := big.NewInt(1000000000000000) // 0.001 ETH（以 wei 为单位）
	gasLimit := uint64(21000)             // 标准 ETH 转账 gas 限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取建议 gas 价格失败: %v", err)
	}

	toAddress := common.HexToAddress(recipientAddressStr)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取链 ID 失败: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("交易签名失败: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Printf("交易哈希: %s\n", signedTx.Hash().Hex())
}
