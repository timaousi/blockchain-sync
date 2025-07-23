// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Task2ContractsMetaData contains all meta data concerning the Task2Contracts contract.
var Task2ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f5f81905550610381806100225f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806306661abd1461004e5780632baeceb71461006c578063a87d942c14610076578063d09de08a14610094575b5f5ffd5b61005661009e565b604051610063919061019b565b60405180910390f35b6100746100a3565b005b61007e610114565b60405161008b919061019b565b60405180910390f35b61009c61011c565b005b5f5481565b5f5f541115610112575f5f8154809291906100bd906101e1565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f89b3507ed3792bd341f5c1001be8efeac401e77d678ab2d92e1841120d14b7455f546040516101099190610262565b60405180910390a25b565b5f5f54905090565b5f5f81548092919061012d9061028e565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f89b3507ed3792bd341f5c1001be8efeac401e77d678ab2d92e1841120d14b7455f54604051610179919061031f565b60405180910390a2565b5f819050919050565b61019581610183565b82525050565b5f6020820190506101ae5f83018461018c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101eb82610183565b91505f82036101fd576101fc6101b4565b5b600182039050919050565b5f82825260208201905092915050565b7fe8aea1e695b0e5b7b2e5878fe5b09120f09f98a50000000000000000000000005f82015250565b5f61024c601483610208565b915061025782610218565b602082019050919050565b5f6040820190506102755f83018461018c565b818103602083015261028681610240565b905092915050565b5f61029882610183565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036102ca576102c96101b4565b5b600182019050919050565b7fe8aea1e695b0e5b7b2e5a29ee58aa020f09f98840000000000000000000000005f82015250565b5f610309601483610208565b9150610314826102d5565b602082019050919050565b5f6040820190506103325f83018461018c565b8181036020830152610343816102fd565b90509291505056fea26469706673582212205854f7cdc67390ac14d3c5e96cdb17f7fe81e4981261bfdb0deaaaa95c80122d64736f6c634300081e0033",
}

// Task2ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use Task2ContractsMetaData.ABI instead.
var Task2ContractsABI = Task2ContractsMetaData.ABI

// Task2ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Task2ContractsMetaData.Bin instead.
var Task2ContractsBin = Task2ContractsMetaData.Bin

// DeployTask2Contracts deploys a new Ethereum contract, binding an instance of Task2Contracts to it.
func DeployTask2Contracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Task2Contracts, error) {
	parsed, err := Task2ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Task2ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Task2Contracts{Task2ContractsCaller: Task2ContractsCaller{contract: contract}, Task2ContractsTransactor: Task2ContractsTransactor{contract: contract}, Task2ContractsFilterer: Task2ContractsFilterer{contract: contract}}, nil
}

// Task2Contracts is an auto generated Go binding around an Ethereum contract.
type Task2Contracts struct {
	Task2ContractsCaller     // Read-only binding to the contract
	Task2ContractsTransactor // Write-only binding to the contract
	Task2ContractsFilterer   // Log filterer for contract events
}

// Task2ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Task2ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Task2ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Task2ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Task2ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Task2ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Task2ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Task2ContractsSession struct {
	Contract     *Task2Contracts   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Task2ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Task2ContractsCallerSession struct {
	Contract *Task2ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Task2ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Task2ContractsTransactorSession struct {
	Contract     *Task2ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Task2ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Task2ContractsRaw struct {
	Contract *Task2Contracts // Generic contract binding to access the raw methods on
}

// Task2ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Task2ContractsCallerRaw struct {
	Contract *Task2ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// Task2ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Task2ContractsTransactorRaw struct {
	Contract *Task2ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask2Contracts creates a new instance of Task2Contracts, bound to a specific deployed contract.
func NewTask2Contracts(address common.Address, backend bind.ContractBackend) (*Task2Contracts, error) {
	contract, err := bindTask2Contracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task2Contracts{Task2ContractsCaller: Task2ContractsCaller{contract: contract}, Task2ContractsTransactor: Task2ContractsTransactor{contract: contract}, Task2ContractsFilterer: Task2ContractsFilterer{contract: contract}}, nil
}

// NewTask2ContractsCaller creates a new read-only instance of Task2Contracts, bound to a specific deployed contract.
func NewTask2ContractsCaller(address common.Address, caller bind.ContractCaller) (*Task2ContractsCaller, error) {
	contract, err := bindTask2Contracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Task2ContractsCaller{contract: contract}, nil
}

// NewTask2ContractsTransactor creates a new write-only instance of Task2Contracts, bound to a specific deployed contract.
func NewTask2ContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*Task2ContractsTransactor, error) {
	contract, err := bindTask2Contracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Task2ContractsTransactor{contract: contract}, nil
}

// NewTask2ContractsFilterer creates a new log filterer instance of Task2Contracts, bound to a specific deployed contract.
func NewTask2ContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*Task2ContractsFilterer, error) {
	contract, err := bindTask2Contracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Task2ContractsFilterer{contract: contract}, nil
}

// bindTask2Contracts binds a generic wrapper to an already deployed contract.
func bindTask2Contracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Task2ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task2Contracts *Task2ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task2Contracts.Contract.Task2ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task2Contracts *Task2ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task2Contracts.Contract.Task2ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task2Contracts *Task2ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task2Contracts.Contract.Task2ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task2Contracts *Task2ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task2Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task2Contracts *Task2ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task2Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task2Contracts *Task2ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task2Contracts.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Task2Contracts *Task2ContractsCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task2Contracts.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Task2Contracts *Task2ContractsSession) Count() (*big.Int, error) {
	return _Task2Contracts.Contract.Count(&_Task2Contracts.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Task2Contracts *Task2ContractsCallerSession) Count() (*big.Int, error) {
	return _Task2Contracts.Contract.Count(&_Task2Contracts.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Task2Contracts *Task2ContractsCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task2Contracts.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Task2Contracts *Task2ContractsSession) GetCount() (*big.Int, error) {
	return _Task2Contracts.Contract.GetCount(&_Task2Contracts.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Task2Contracts *Task2ContractsCallerSession) GetCount() (*big.Int, error) {
	return _Task2Contracts.Contract.GetCount(&_Task2Contracts.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Task2Contracts *Task2ContractsTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task2Contracts.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Task2Contracts *Task2ContractsSession) Decrement() (*types.Transaction, error) {
	return _Task2Contracts.Contract.Decrement(&_Task2Contracts.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Task2Contracts *Task2ContractsTransactorSession) Decrement() (*types.Transaction, error) {
	return _Task2Contracts.Contract.Decrement(&_Task2Contracts.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Task2Contracts *Task2ContractsTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task2Contracts.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Task2Contracts *Task2ContractsSession) Increment() (*types.Transaction, error) {
	return _Task2Contracts.Contract.Increment(&_Task2Contracts.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Task2Contracts *Task2ContractsTransactorSession) Increment() (*types.Transaction, error) {
	return _Task2Contracts.Contract.Increment(&_Task2Contracts.TransactOpts)
}

// Task2ContractsCountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the Task2Contracts contract.
type Task2ContractsCountChangedIterator struct {
	Event *Task2ContractsCountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Task2ContractsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Task2ContractsCountChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Task2ContractsCountChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Task2ContractsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Task2ContractsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Task2ContractsCountChanged represents a CountChanged event raised by the Task2Contracts contract.
type Task2ContractsCountChanged struct {
	NewCount *big.Int
	Changer  common.Address
	Message  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0x89b3507ed3792bd341f5c1001be8efeac401e77d678ab2d92e1841120d14b745.
//
// Solidity: event CountChanged(uint256 newCount, address indexed changer, string message)
func (_Task2Contracts *Task2ContractsFilterer) FilterCountChanged(opts *bind.FilterOpts, changer []common.Address) (*Task2ContractsCountChangedIterator, error) {

	var changerRule []interface{}
	for _, changerItem := range changer {
		changerRule = append(changerRule, changerItem)
	}

	logs, sub, err := _Task2Contracts.contract.FilterLogs(opts, "CountChanged", changerRule)
	if err != nil {
		return nil, err
	}
	return &Task2ContractsCountChangedIterator{contract: _Task2Contracts.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0x89b3507ed3792bd341f5c1001be8efeac401e77d678ab2d92e1841120d14b745.
//
// Solidity: event CountChanged(uint256 newCount, address indexed changer, string message)
func (_Task2Contracts *Task2ContractsFilterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *Task2ContractsCountChanged, changer []common.Address) (event.Subscription, error) {

	var changerRule []interface{}
	for _, changerItem := range changer {
		changerRule = append(changerRule, changerItem)
	}

	logs, sub, err := _Task2Contracts.contract.WatchLogs(opts, "CountChanged", changerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Task2ContractsCountChanged)
				if err := _Task2Contracts.contract.UnpackLog(event, "CountChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCountChanged is a log parse operation binding the contract event 0x89b3507ed3792bd341f5c1001be8efeac401e77d678ab2d92e1841120d14b745.
//
// Solidity: event CountChanged(uint256 newCount, address indexed changer, string message)
func (_Task2Contracts *Task2ContractsFilterer) ParseCountChanged(log types.Log) (*Task2ContractsCountChanged, error) {
	event := new(Task2ContractsCountChanged)
	if err := _Task2Contracts.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
