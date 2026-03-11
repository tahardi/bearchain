// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// FactorsVerifierMetaData contains all meta data concerning the FactorsVerifier contract.
var FactorsVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_riscZeroVerifier\",\"type\":\"address\",\"internalType\":\"contractIRiscZeroVerifier\"},{\"name\":\"_imageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OWNER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RISC_ZERO_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRiscZeroVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setImageId\",\"inputs\":[{\"name\":\"_imageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"product\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DebugVerification\",\"inputs\":[{\"name\":\"sealSelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"expectedSelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FactorsKnownForProduct\",\"inputs\":[{\"name\":\"product\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageIdUpdated\",\"inputs\":[{\"name\":\"newImageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// FactorsVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use FactorsVerifierMetaData.ABI instead.
var FactorsVerifierABI = FactorsVerifierMetaData.ABI

// FactorsVerifier is an auto generated Go binding around an Ethereum contract.
type FactorsVerifier struct {
	FactorsVerifierCaller     // Read-only binding to the contract
	FactorsVerifierTransactor // Write-only binding to the contract
	FactorsVerifierFilterer   // Log filterer for contract events
}

// FactorsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactorsVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactorsVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactorsVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorsVerifierSession struct {
	Contract     *FactorsVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactorsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactorsVerifierCallerSession struct {
	Contract *FactorsVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FactorsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactorsVerifierTransactorSession struct {
	Contract     *FactorsVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FactorsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactorsVerifierRaw struct {
	Contract *FactorsVerifier // Generic contract binding to access the raw methods on
}

// FactorsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactorsVerifierCallerRaw struct {
	Contract *FactorsVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// FactorsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactorsVerifierTransactorRaw struct {
	Contract *FactorsVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactorsVerifier creates a new instance of FactorsVerifier, bound to a specific deployed contract.
func NewFactorsVerifier(address common.Address, backend bind.ContractBackend) (*FactorsVerifier, error) {
	contract, err := bindFactorsVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifier{FactorsVerifierCaller: FactorsVerifierCaller{contract: contract}, FactorsVerifierTransactor: FactorsVerifierTransactor{contract: contract}, FactorsVerifierFilterer: FactorsVerifierFilterer{contract: contract}}, nil
}

// NewFactorsVerifierCaller creates a new read-only instance of FactorsVerifier, bound to a specific deployed contract.
func NewFactorsVerifierCaller(address common.Address, caller bind.ContractCaller) (*FactorsVerifierCaller, error) {
	contract, err := bindFactorsVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierCaller{contract: contract}, nil
}

// NewFactorsVerifierTransactor creates a new write-only instance of FactorsVerifier, bound to a specific deployed contract.
func NewFactorsVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*FactorsVerifierTransactor, error) {
	contract, err := bindFactorsVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierTransactor{contract: contract}, nil
}

// NewFactorsVerifierFilterer creates a new log filterer instance of FactorsVerifier, bound to a specific deployed contract.
func NewFactorsVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*FactorsVerifierFilterer, error) {
	contract, err := bindFactorsVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierFilterer{contract: contract}, nil
}

// bindFactorsVerifier binds a generic wrapper to an already deployed contract.
func bindFactorsVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactorsVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactorsVerifier *FactorsVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactorsVerifier.Contract.FactorsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactorsVerifier *FactorsVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.FactorsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactorsVerifier *FactorsVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.FactorsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactorsVerifier *FactorsVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactorsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactorsVerifier *FactorsVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactorsVerifier *FactorsVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.contract.Transact(opts, method, params...)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_FactorsVerifier *FactorsVerifierCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactorsVerifier.contract.Call(opts, &out, "OWNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_FactorsVerifier *FactorsVerifierSession) OWNER() (common.Address, error) {
	return _FactorsVerifier.Contract.OWNER(&_FactorsVerifier.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_FactorsVerifier *FactorsVerifierCallerSession) OWNER() (common.Address, error) {
	return _FactorsVerifier.Contract.OWNER(&_FactorsVerifier.CallOpts)
}

// RISCZEROVERIFIER is a free data retrieval call binding the contract method 0xf3c0ce22.
//
// Solidity: function RISC_ZERO_VERIFIER() view returns(address)
func (_FactorsVerifier *FactorsVerifierCaller) RISCZEROVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactorsVerifier.contract.Call(opts, &out, "RISC_ZERO_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RISCZEROVERIFIER is a free data retrieval call binding the contract method 0xf3c0ce22.
//
// Solidity: function RISC_ZERO_VERIFIER() view returns(address)
func (_FactorsVerifier *FactorsVerifierSession) RISCZEROVERIFIER() (common.Address, error) {
	return _FactorsVerifier.Contract.RISCZEROVERIFIER(&_FactorsVerifier.CallOpts)
}

// RISCZEROVERIFIER is a free data retrieval call binding the contract method 0xf3c0ce22.
//
// Solidity: function RISC_ZERO_VERIFIER() view returns(address)
func (_FactorsVerifier *FactorsVerifierCallerSession) RISCZEROVERIFIER() (common.Address, error) {
	return _FactorsVerifier.Contract.RISCZEROVERIFIER(&_FactorsVerifier.CallOpts)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_FactorsVerifier *FactorsVerifierCaller) ImageId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FactorsVerifier.contract.Call(opts, &out, "imageId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_FactorsVerifier *FactorsVerifierSession) ImageId() ([32]byte, error) {
	return _FactorsVerifier.Contract.ImageId(&_FactorsVerifier.CallOpts)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_FactorsVerifier *FactorsVerifierCallerSession) ImageId() ([32]byte, error) {
	return _FactorsVerifier.Contract.ImageId(&_FactorsVerifier.CallOpts)
}

// SetImageId is a paid mutator transaction binding the contract method 0x901129c2.
//
// Solidity: function setImageId(bytes32 _imageId) returns()
func (_FactorsVerifier *FactorsVerifierTransactor) SetImageId(opts *bind.TransactOpts, _imageId [32]byte) (*types.Transaction, error) {
	return _FactorsVerifier.contract.Transact(opts, "setImageId", _imageId)
}

// SetImageId is a paid mutator transaction binding the contract method 0x901129c2.
//
// Solidity: function setImageId(bytes32 _imageId) returns()
func (_FactorsVerifier *FactorsVerifierSession) SetImageId(_imageId [32]byte) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.SetImageId(&_FactorsVerifier.TransactOpts, _imageId)
}

// SetImageId is a paid mutator transaction binding the contract method 0x901129c2.
//
// Solidity: function setImageId(bytes32 _imageId) returns()
func (_FactorsVerifier *FactorsVerifierTransactorSession) SetImageId(_imageId [32]byte) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.SetImageId(&_FactorsVerifier.TransactOpts, _imageId)
}

// Verify is a paid mutator transaction binding the contract method 0xf7ddea5a.
//
// Solidity: function verify(uint64 product, bytes seal) returns()
func (_FactorsVerifier *FactorsVerifierTransactor) Verify(opts *bind.TransactOpts, product uint64, seal []byte) (*types.Transaction, error) {
	return _FactorsVerifier.contract.Transact(opts, "verify", product, seal)
}

// Verify is a paid mutator transaction binding the contract method 0xf7ddea5a.
//
// Solidity: function verify(uint64 product, bytes seal) returns()
func (_FactorsVerifier *FactorsVerifierSession) Verify(product uint64, seal []byte) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.Verify(&_FactorsVerifier.TransactOpts, product, seal)
}

// Verify is a paid mutator transaction binding the contract method 0xf7ddea5a.
//
// Solidity: function verify(uint64 product, bytes seal) returns()
func (_FactorsVerifier *FactorsVerifierTransactorSession) Verify(product uint64, seal []byte) (*types.Transaction, error) {
	return _FactorsVerifier.Contract.Verify(&_FactorsVerifier.TransactOpts, product, seal)
}

// FactorsVerifierDebugVerificationIterator is returned from FilterDebugVerification and is used to iterate over the raw logs and unpacked data for DebugVerification events raised by the FactorsVerifier contract.
type FactorsVerifierDebugVerificationIterator struct {
	Event *FactorsVerifierDebugVerification // Event containing the contract specifics and raw log

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
func (it *FactorsVerifierDebugVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactorsVerifierDebugVerification)
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
		it.Event = new(FactorsVerifierDebugVerification)
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
func (it *FactorsVerifierDebugVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactorsVerifierDebugVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactorsVerifierDebugVerification represents a DebugVerification event raised by the FactorsVerifier contract.
type FactorsVerifierDebugVerification struct {
	SealSelector     [4]byte
	ExpectedSelector [4]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDebugVerification is a free log retrieval operation binding the contract event 0x3c48bde270adf6a2395093bec32b741e175ee4b363acbb991b79b579a97fd15b.
//
// Solidity: event DebugVerification(bytes4 sealSelector, bytes4 expectedSelector)
func (_FactorsVerifier *FactorsVerifierFilterer) FilterDebugVerification(opts *bind.FilterOpts) (*FactorsVerifierDebugVerificationIterator, error) {

	logs, sub, err := _FactorsVerifier.contract.FilterLogs(opts, "DebugVerification")
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierDebugVerificationIterator{contract: _FactorsVerifier.contract, event: "DebugVerification", logs: logs, sub: sub}, nil
}

// WatchDebugVerification is a free log subscription operation binding the contract event 0x3c48bde270adf6a2395093bec32b741e175ee4b363acbb991b79b579a97fd15b.
//
// Solidity: event DebugVerification(bytes4 sealSelector, bytes4 expectedSelector)
func (_FactorsVerifier *FactorsVerifierFilterer) WatchDebugVerification(opts *bind.WatchOpts, sink chan<- *FactorsVerifierDebugVerification) (event.Subscription, error) {

	logs, sub, err := _FactorsVerifier.contract.WatchLogs(opts, "DebugVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactorsVerifierDebugVerification)
				if err := _FactorsVerifier.contract.UnpackLog(event, "DebugVerification", log); err != nil {
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

// ParseDebugVerification is a log parse operation binding the contract event 0x3c48bde270adf6a2395093bec32b741e175ee4b363acbb991b79b579a97fd15b.
//
// Solidity: event DebugVerification(bytes4 sealSelector, bytes4 expectedSelector)
func (_FactorsVerifier *FactorsVerifierFilterer) ParseDebugVerification(log types.Log) (*FactorsVerifierDebugVerification, error) {
	event := new(FactorsVerifierDebugVerification)
	if err := _FactorsVerifier.contract.UnpackLog(event, "DebugVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactorsVerifierFactorsKnownForProductIterator is returned from FilterFactorsKnownForProduct and is used to iterate over the raw logs and unpacked data for FactorsKnownForProduct events raised by the FactorsVerifier contract.
type FactorsVerifierFactorsKnownForProductIterator struct {
	Event *FactorsVerifierFactorsKnownForProduct // Event containing the contract specifics and raw log

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
func (it *FactorsVerifierFactorsKnownForProductIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactorsVerifierFactorsKnownForProduct)
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
		it.Event = new(FactorsVerifierFactorsKnownForProduct)
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
func (it *FactorsVerifierFactorsKnownForProductIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactorsVerifierFactorsKnownForProductIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactorsVerifierFactorsKnownForProduct represents a FactorsKnownForProduct event raised by the FactorsVerifier contract.
type FactorsVerifierFactorsKnownForProduct struct {
	Product uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFactorsKnownForProduct is a free log retrieval operation binding the contract event 0x08b620c727b9819c2bd9420a0bd3cdc6a9c4ef4cde2d5730fae682a346248b72.
//
// Solidity: event FactorsKnownForProduct(uint64 indexed product)
func (_FactorsVerifier *FactorsVerifierFilterer) FilterFactorsKnownForProduct(opts *bind.FilterOpts, product []uint64) (*FactorsVerifierFactorsKnownForProductIterator, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _FactorsVerifier.contract.FilterLogs(opts, "FactorsKnownForProduct", productRule)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierFactorsKnownForProductIterator{contract: _FactorsVerifier.contract, event: "FactorsKnownForProduct", logs: logs, sub: sub}, nil
}

// WatchFactorsKnownForProduct is a free log subscription operation binding the contract event 0x08b620c727b9819c2bd9420a0bd3cdc6a9c4ef4cde2d5730fae682a346248b72.
//
// Solidity: event FactorsKnownForProduct(uint64 indexed product)
func (_FactorsVerifier *FactorsVerifierFilterer) WatchFactorsKnownForProduct(opts *bind.WatchOpts, sink chan<- *FactorsVerifierFactorsKnownForProduct, product []uint64) (event.Subscription, error) {

	var productRule []interface{}
	for _, productItem := range product {
		productRule = append(productRule, productItem)
	}

	logs, sub, err := _FactorsVerifier.contract.WatchLogs(opts, "FactorsKnownForProduct", productRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactorsVerifierFactorsKnownForProduct)
				if err := _FactorsVerifier.contract.UnpackLog(event, "FactorsKnownForProduct", log); err != nil {
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

// ParseFactorsKnownForProduct is a log parse operation binding the contract event 0x08b620c727b9819c2bd9420a0bd3cdc6a9c4ef4cde2d5730fae682a346248b72.
//
// Solidity: event FactorsKnownForProduct(uint64 indexed product)
func (_FactorsVerifier *FactorsVerifierFilterer) ParseFactorsKnownForProduct(log types.Log) (*FactorsVerifierFactorsKnownForProduct, error) {
	event := new(FactorsVerifierFactorsKnownForProduct)
	if err := _FactorsVerifier.contract.UnpackLog(event, "FactorsKnownForProduct", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactorsVerifierImageIdUpdatedIterator is returned from FilterImageIdUpdated and is used to iterate over the raw logs and unpacked data for ImageIdUpdated events raised by the FactorsVerifier contract.
type FactorsVerifierImageIdUpdatedIterator struct {
	Event *FactorsVerifierImageIdUpdated // Event containing the contract specifics and raw log

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
func (it *FactorsVerifierImageIdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactorsVerifierImageIdUpdated)
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
		it.Event = new(FactorsVerifierImageIdUpdated)
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
func (it *FactorsVerifierImageIdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactorsVerifierImageIdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactorsVerifierImageIdUpdated represents a ImageIdUpdated event raised by the FactorsVerifier contract.
type FactorsVerifierImageIdUpdated struct {
	NewImageId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterImageIdUpdated is a free log retrieval operation binding the contract event 0xb4e3b50ca9dccac181336b4e74d2f5293b4b69adcc057d39456e8d889efc48dd.
//
// Solidity: event ImageIdUpdated(bytes32 indexed newImageId)
func (_FactorsVerifier *FactorsVerifierFilterer) FilterImageIdUpdated(opts *bind.FilterOpts, newImageId [][32]byte) (*FactorsVerifierImageIdUpdatedIterator, error) {

	var newImageIdRule []interface{}
	for _, newImageIdItem := range newImageId {
		newImageIdRule = append(newImageIdRule, newImageIdItem)
	}

	logs, sub, err := _FactorsVerifier.contract.FilterLogs(opts, "ImageIdUpdated", newImageIdRule)
	if err != nil {
		return nil, err
	}
	return &FactorsVerifierImageIdUpdatedIterator{contract: _FactorsVerifier.contract, event: "ImageIdUpdated", logs: logs, sub: sub}, nil
}

// WatchImageIdUpdated is a free log subscription operation binding the contract event 0xb4e3b50ca9dccac181336b4e74d2f5293b4b69adcc057d39456e8d889efc48dd.
//
// Solidity: event ImageIdUpdated(bytes32 indexed newImageId)
func (_FactorsVerifier *FactorsVerifierFilterer) WatchImageIdUpdated(opts *bind.WatchOpts, sink chan<- *FactorsVerifierImageIdUpdated, newImageId [][32]byte) (event.Subscription, error) {

	var newImageIdRule []interface{}
	for _, newImageIdItem := range newImageId {
		newImageIdRule = append(newImageIdRule, newImageIdItem)
	}

	logs, sub, err := _FactorsVerifier.contract.WatchLogs(opts, "ImageIdUpdated", newImageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactorsVerifierImageIdUpdated)
				if err := _FactorsVerifier.contract.UnpackLog(event, "ImageIdUpdated", log); err != nil {
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

// ParseImageIdUpdated is a log parse operation binding the contract event 0xb4e3b50ca9dccac181336b4e74d2f5293b4b69adcc057d39456e8d889efc48dd.
//
// Solidity: event ImageIdUpdated(bytes32 indexed newImageId)
func (_FactorsVerifier *FactorsVerifierFilterer) ParseImageIdUpdated(log types.Log) (*FactorsVerifierImageIdUpdated, error) {
	event := new(FactorsVerifierImageIdUpdated)
	if err := _FactorsVerifier.contract.UnpackLog(event, "ImageIdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
