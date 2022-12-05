// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oraclesanctions

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
)

// ConsumerContractMetaData contains all meta data concerning the ConsumerContract contract.
var ConsumerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NonSanctionedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SanctionedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"SanctionedAddressesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"SanctionedAddressesRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSanctions\",\"type\":\"address[]\"}],\"name\":\"addToSanctionsList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSanctioned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSanctionedVerbose\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removeSanctions\",\"type\":\"address[]\"}],\"name\":\"removeFromSanctionsList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConsumerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsumerContractMetaData.ABI instead.
var ConsumerContractABI = ConsumerContractMetaData.ABI

// ConsumerContract is an auto generated Go binding around an Ethereum contract.
type ConsumerContract struct {
	ConsumerContractCaller     // Read-only binding to the contract
	ConsumerContractTransactor // Write-only binding to the contract
	ConsumerContractFilterer   // Log filterer for contract events
}

// ConsumerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsumerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsumerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsumerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsumerContractSession struct {
	Contract     *ConsumerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsumerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsumerContractCallerSession struct {
	Contract *ConsumerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ConsumerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsumerContractTransactorSession struct {
	Contract     *ConsumerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConsumerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsumerContractRaw struct {
	Contract *ConsumerContract // Generic contract binding to access the raw methods on
}

// ConsumerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsumerContractCallerRaw struct {
	Contract *ConsumerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ConsumerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsumerContractTransactorRaw struct {
	Contract *ConsumerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsumerContract creates a new instance of ConsumerContract, bound to a specific deployed contract.
func NewConsumerContract(address common.Address, backend bind.ContractBackend) (*ConsumerContract, error) {
	contract, err := bindConsumerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsumerContract{ConsumerContractCaller: ConsumerContractCaller{contract: contract}, ConsumerContractTransactor: ConsumerContractTransactor{contract: contract}, ConsumerContractFilterer: ConsumerContractFilterer{contract: contract}}, nil
}

// NewConsumerContractCaller creates a new read-only instance of ConsumerContract, bound to a specific deployed contract.
func NewConsumerContractCaller(address common.Address, caller bind.ContractCaller) (*ConsumerContractCaller, error) {
	contract, err := bindConsumerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractCaller{contract: contract}, nil
}

// NewConsumerContractTransactor creates a new write-only instance of ConsumerContract, bound to a specific deployed contract.
func NewConsumerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsumerContractTransactor, error) {
	contract, err := bindConsumerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractTransactor{contract: contract}, nil
}

// NewConsumerContractFilterer creates a new log filterer instance of ConsumerContract, bound to a specific deployed contract.
func NewConsumerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsumerContractFilterer, error) {
	contract, err := bindConsumerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractFilterer{contract: contract}, nil
}

// bindConsumerContract binds a generic wrapper to an already deployed contract.
func bindConsumerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsumerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumerContract *ConsumerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsumerContract.Contract.ConsumerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumerContract *ConsumerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumerContract.Contract.ConsumerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumerContract *ConsumerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumerContract.Contract.ConsumerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumerContract *ConsumerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsumerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumerContract *ConsumerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumerContract *ConsumerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumerContract.Contract.contract.Transact(opts, method, params...)
}

// IsSanctioned is a free data retrieval call binding the contract method 0xdf592f7d.
//
// Solidity: function isSanctioned(address addr) view returns(bool)
func (_ConsumerContract *ConsumerContractCaller) IsSanctioned(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ConsumerContract.contract.Call(opts, &out, "isSanctioned", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSanctioned is a free data retrieval call binding the contract method 0xdf592f7d.
//
// Solidity: function isSanctioned(address addr) view returns(bool)
func (_ConsumerContract *ConsumerContractSession) IsSanctioned(addr common.Address) (bool, error) {
	return _ConsumerContract.Contract.IsSanctioned(&_ConsumerContract.CallOpts, addr)
}

// IsSanctioned is a free data retrieval call binding the contract method 0xdf592f7d.
//
// Solidity: function isSanctioned(address addr) view returns(bool)
func (_ConsumerContract *ConsumerContractCallerSession) IsSanctioned(addr common.Address) (bool, error) {
	return _ConsumerContract.Contract.IsSanctioned(&_ConsumerContract.CallOpts, addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ConsumerContract *ConsumerContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConsumerContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ConsumerContract *ConsumerContractSession) Name() (string, error) {
	return _ConsumerContract.Contract.Name(&_ConsumerContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ConsumerContract *ConsumerContractCallerSession) Name() (string, error) {
	return _ConsumerContract.Contract.Name(&_ConsumerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConsumerContract *ConsumerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConsumerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConsumerContract *ConsumerContractSession) Owner() (common.Address, error) {
	return _ConsumerContract.Contract.Owner(&_ConsumerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConsumerContract *ConsumerContractCallerSession) Owner() (common.Address, error) {
	return _ConsumerContract.Contract.Owner(&_ConsumerContract.CallOpts)
}

// AddToSanctionsList is a paid mutator transaction binding the contract method 0xb972dfcc.
//
// Solidity: function addToSanctionsList(address[] newSanctions) returns()
func (_ConsumerContract *ConsumerContractTransactor) AddToSanctionsList(opts *bind.TransactOpts, newSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.contract.Transact(opts, "addToSanctionsList", newSanctions)
}

// AddToSanctionsList is a paid mutator transaction binding the contract method 0xb972dfcc.
//
// Solidity: function addToSanctionsList(address[] newSanctions) returns()
func (_ConsumerContract *ConsumerContractSession) AddToSanctionsList(newSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.AddToSanctionsList(&_ConsumerContract.TransactOpts, newSanctions)
}

// AddToSanctionsList is a paid mutator transaction binding the contract method 0xb972dfcc.
//
// Solidity: function addToSanctionsList(address[] newSanctions) returns()
func (_ConsumerContract *ConsumerContractTransactorSession) AddToSanctionsList(newSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.AddToSanctionsList(&_ConsumerContract.TransactOpts, newSanctions)
}

// IsSanctionedVerbose is a paid mutator transaction binding the contract method 0xa2a6bbd8.
//
// Solidity: function isSanctionedVerbose(address addr) returns(bool)
func (_ConsumerContract *ConsumerContractTransactor) IsSanctionedVerbose(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ConsumerContract.contract.Transact(opts, "isSanctionedVerbose", addr)
}

// IsSanctionedVerbose is a paid mutator transaction binding the contract method 0xa2a6bbd8.
//
// Solidity: function isSanctionedVerbose(address addr) returns(bool)
func (_ConsumerContract *ConsumerContractSession) IsSanctionedVerbose(addr common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.IsSanctionedVerbose(&_ConsumerContract.TransactOpts, addr)
}

// IsSanctionedVerbose is a paid mutator transaction binding the contract method 0xa2a6bbd8.
//
// Solidity: function isSanctionedVerbose(address addr) returns(bool)
func (_ConsumerContract *ConsumerContractTransactorSession) IsSanctionedVerbose(addr common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.IsSanctionedVerbose(&_ConsumerContract.TransactOpts, addr)
}

// RemoveFromSanctionsList is a paid mutator transaction binding the contract method 0xef782431.
//
// Solidity: function removeFromSanctionsList(address[] removeSanctions) returns()
func (_ConsumerContract *ConsumerContractTransactor) RemoveFromSanctionsList(opts *bind.TransactOpts, removeSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.contract.Transact(opts, "removeFromSanctionsList", removeSanctions)
}

// RemoveFromSanctionsList is a paid mutator transaction binding the contract method 0xef782431.
//
// Solidity: function removeFromSanctionsList(address[] removeSanctions) returns()
func (_ConsumerContract *ConsumerContractSession) RemoveFromSanctionsList(removeSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.RemoveFromSanctionsList(&_ConsumerContract.TransactOpts, removeSanctions)
}

// RemoveFromSanctionsList is a paid mutator transaction binding the contract method 0xef782431.
//
// Solidity: function removeFromSanctionsList(address[] removeSanctions) returns()
func (_ConsumerContract *ConsumerContractTransactorSession) RemoveFromSanctionsList(removeSanctions []common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.RemoveFromSanctionsList(&_ConsumerContract.TransactOpts, removeSanctions)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConsumerContract *ConsumerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConsumerContract *ConsumerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConsumerContract.Contract.RenounceOwnership(&_ConsumerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConsumerContract *ConsumerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConsumerContract.Contract.RenounceOwnership(&_ConsumerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConsumerContract *ConsumerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConsumerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConsumerContract *ConsumerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.TransferOwnership(&_ConsumerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConsumerContract *ConsumerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConsumerContract.Contract.TransferOwnership(&_ConsumerContract.TransactOpts, newOwner)
}

// ConsumerContractNonSanctionedAddressIterator is returned from FilterNonSanctionedAddress and is used to iterate over the raw logs and unpacked data for NonSanctionedAddress events raised by the ConsumerContract contract.
type ConsumerContractNonSanctionedAddressIterator struct {
	Event *ConsumerContractNonSanctionedAddress // Event containing the contract specifics and raw log

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
func (it *ConsumerContractNonSanctionedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumerContractNonSanctionedAddress)
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
		it.Event = new(ConsumerContractNonSanctionedAddress)
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
func (it *ConsumerContractNonSanctionedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumerContractNonSanctionedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumerContractNonSanctionedAddress represents a NonSanctionedAddress event raised by the ConsumerContract contract.
type ConsumerContractNonSanctionedAddress struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNonSanctionedAddress is a free log retrieval operation binding the contract event 0xd595018321fcb8c2bcbf5bfe4b27d74bea505825f7d195abe8517f94a065539c.
//
// Solidity: event NonSanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) FilterNonSanctionedAddress(opts *bind.FilterOpts, addr []common.Address) (*ConsumerContractNonSanctionedAddressIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ConsumerContract.contract.FilterLogs(opts, "NonSanctionedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractNonSanctionedAddressIterator{contract: _ConsumerContract.contract, event: "NonSanctionedAddress", logs: logs, sub: sub}, nil
}

// WatchNonSanctionedAddress is a free log subscription operation binding the contract event 0xd595018321fcb8c2bcbf5bfe4b27d74bea505825f7d195abe8517f94a065539c.
//
// Solidity: event NonSanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) WatchNonSanctionedAddress(opts *bind.WatchOpts, sink chan<- *ConsumerContractNonSanctionedAddress, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ConsumerContract.contract.WatchLogs(opts, "NonSanctionedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumerContractNonSanctionedAddress)
				if err := _ConsumerContract.contract.UnpackLog(event, "NonSanctionedAddress", log); err != nil {
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

// ParseNonSanctionedAddress is a log parse operation binding the contract event 0xd595018321fcb8c2bcbf5bfe4b27d74bea505825f7d195abe8517f94a065539c.
//
// Solidity: event NonSanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) ParseNonSanctionedAddress(log types.Log) (*ConsumerContractNonSanctionedAddress, error) {
	event := new(ConsumerContractNonSanctionedAddress)
	if err := _ConsumerContract.contract.UnpackLog(event, "NonSanctionedAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsumerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConsumerContract contract.
type ConsumerContractOwnershipTransferredIterator struct {
	Event *ConsumerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConsumerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumerContractOwnershipTransferred)
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
		it.Event = new(ConsumerContractOwnershipTransferred)
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
func (it *ConsumerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumerContractOwnershipTransferred represents a OwnershipTransferred event raised by the ConsumerContract contract.
type ConsumerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConsumerContract *ConsumerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConsumerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConsumerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractOwnershipTransferredIterator{contract: _ConsumerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConsumerContract *ConsumerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConsumerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConsumerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumerContractOwnershipTransferred)
				if err := _ConsumerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConsumerContract *ConsumerContractFilterer) ParseOwnershipTransferred(log types.Log) (*ConsumerContractOwnershipTransferred, error) {
	event := new(ConsumerContractOwnershipTransferred)
	if err := _ConsumerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsumerContractSanctionedAddressIterator is returned from FilterSanctionedAddress and is used to iterate over the raw logs and unpacked data for SanctionedAddress events raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddressIterator struct {
	Event *ConsumerContractSanctionedAddress // Event containing the contract specifics and raw log

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
func (it *ConsumerContractSanctionedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumerContractSanctionedAddress)
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
		it.Event = new(ConsumerContractSanctionedAddress)
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
func (it *ConsumerContractSanctionedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumerContractSanctionedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumerContractSanctionedAddress represents a SanctionedAddress event raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddress struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSanctionedAddress is a free log retrieval operation binding the contract event 0x8027911123971054d93579ebea046c8461473fa4d2e510b9b49eed3bed3270e0.
//
// Solidity: event SanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) FilterSanctionedAddress(opts *bind.FilterOpts, addr []common.Address) (*ConsumerContractSanctionedAddressIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ConsumerContract.contract.FilterLogs(opts, "SanctionedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return &ConsumerContractSanctionedAddressIterator{contract: _ConsumerContract.contract, event: "SanctionedAddress", logs: logs, sub: sub}, nil
}

// WatchSanctionedAddress is a free log subscription operation binding the contract event 0x8027911123971054d93579ebea046c8461473fa4d2e510b9b49eed3bed3270e0.
//
// Solidity: event SanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) WatchSanctionedAddress(opts *bind.WatchOpts, sink chan<- *ConsumerContractSanctionedAddress, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ConsumerContract.contract.WatchLogs(opts, "SanctionedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumerContractSanctionedAddress)
				if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddress", log); err != nil {
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

// ParseSanctionedAddress is a log parse operation binding the contract event 0x8027911123971054d93579ebea046c8461473fa4d2e510b9b49eed3bed3270e0.
//
// Solidity: event SanctionedAddress(address indexed addr)
func (_ConsumerContract *ConsumerContractFilterer) ParseSanctionedAddress(log types.Log) (*ConsumerContractSanctionedAddress, error) {
	event := new(ConsumerContractSanctionedAddress)
	if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsumerContractSanctionedAddressesAddedIterator is returned from FilterSanctionedAddressesAdded and is used to iterate over the raw logs and unpacked data for SanctionedAddressesAdded events raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddressesAddedIterator struct {
	Event *ConsumerContractSanctionedAddressesAdded // Event containing the contract specifics and raw log

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
func (it *ConsumerContractSanctionedAddressesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumerContractSanctionedAddressesAdded)
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
		it.Event = new(ConsumerContractSanctionedAddressesAdded)
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
func (it *ConsumerContractSanctionedAddressesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumerContractSanctionedAddressesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumerContractSanctionedAddressesAdded represents a SanctionedAddressesAdded event raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddressesAdded struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSanctionedAddressesAdded is a free log retrieval operation binding the contract event 0x2596d7dd6966c5673f9c06ddb0564c4f0e6d8d206ea075b83ad9ddd71a4fb927.
//
// Solidity: event SanctionedAddressesAdded(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) FilterSanctionedAddressesAdded(opts *bind.FilterOpts) (*ConsumerContractSanctionedAddressesAddedIterator, error) {

	logs, sub, err := _ConsumerContract.contract.FilterLogs(opts, "SanctionedAddressesAdded")
	if err != nil {
		return nil, err
	}
	return &ConsumerContractSanctionedAddressesAddedIterator{contract: _ConsumerContract.contract, event: "SanctionedAddressesAdded", logs: logs, sub: sub}, nil
}

// WatchSanctionedAddressesAdded is a free log subscription operation binding the contract event 0x2596d7dd6966c5673f9c06ddb0564c4f0e6d8d206ea075b83ad9ddd71a4fb927.
//
// Solidity: event SanctionedAddressesAdded(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) WatchSanctionedAddressesAdded(opts *bind.WatchOpts, sink chan<- *ConsumerContractSanctionedAddressesAdded) (event.Subscription, error) {

	logs, sub, err := _ConsumerContract.contract.WatchLogs(opts, "SanctionedAddressesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumerContractSanctionedAddressesAdded)
				if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddressesAdded", log); err != nil {
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

// ParseSanctionedAddressesAdded is a log parse operation binding the contract event 0x2596d7dd6966c5673f9c06ddb0564c4f0e6d8d206ea075b83ad9ddd71a4fb927.
//
// Solidity: event SanctionedAddressesAdded(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) ParseSanctionedAddressesAdded(log types.Log) (*ConsumerContractSanctionedAddressesAdded, error) {
	event := new(ConsumerContractSanctionedAddressesAdded)
	if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddressesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsumerContractSanctionedAddressesRemovedIterator is returned from FilterSanctionedAddressesRemoved and is used to iterate over the raw logs and unpacked data for SanctionedAddressesRemoved events raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddressesRemovedIterator struct {
	Event *ConsumerContractSanctionedAddressesRemoved // Event containing the contract specifics and raw log

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
func (it *ConsumerContractSanctionedAddressesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumerContractSanctionedAddressesRemoved)
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
		it.Event = new(ConsumerContractSanctionedAddressesRemoved)
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
func (it *ConsumerContractSanctionedAddressesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumerContractSanctionedAddressesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumerContractSanctionedAddressesRemoved represents a SanctionedAddressesRemoved event raised by the ConsumerContract contract.
type ConsumerContractSanctionedAddressesRemoved struct {
	Addrs []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSanctionedAddressesRemoved is a free log retrieval operation binding the contract event 0x32aab684eee99db715515d1a9987a8fe33bb6341b0e35e60db7eab48a08f9a3a.
//
// Solidity: event SanctionedAddressesRemoved(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) FilterSanctionedAddressesRemoved(opts *bind.FilterOpts) (*ConsumerContractSanctionedAddressesRemovedIterator, error) {

	logs, sub, err := _ConsumerContract.contract.FilterLogs(opts, "SanctionedAddressesRemoved")
	if err != nil {
		return nil, err
	}
	return &ConsumerContractSanctionedAddressesRemovedIterator{contract: _ConsumerContract.contract, event: "SanctionedAddressesRemoved", logs: logs, sub: sub}, nil
}

// WatchSanctionedAddressesRemoved is a free log subscription operation binding the contract event 0x32aab684eee99db715515d1a9987a8fe33bb6341b0e35e60db7eab48a08f9a3a.
//
// Solidity: event SanctionedAddressesRemoved(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) WatchSanctionedAddressesRemoved(opts *bind.WatchOpts, sink chan<- *ConsumerContractSanctionedAddressesRemoved) (event.Subscription, error) {

	logs, sub, err := _ConsumerContract.contract.WatchLogs(opts, "SanctionedAddressesRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumerContractSanctionedAddressesRemoved)
				if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddressesRemoved", log); err != nil {
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

// ParseSanctionedAddressesRemoved is a log parse operation binding the contract event 0x32aab684eee99db715515d1a9987a8fe33bb6341b0e35e60db7eab48a08f9a3a.
//
// Solidity: event SanctionedAddressesRemoved(address[] addrs)
func (_ConsumerContract *ConsumerContractFilterer) ParseSanctionedAddressesRemoved(log types.Log) (*ConsumerContractSanctionedAddressesRemoved, error) {
	event := new(ConsumerContractSanctionedAddressesRemoved)
	if err := _ConsumerContract.contract.UnpackLog(event, "SanctionedAddressesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
