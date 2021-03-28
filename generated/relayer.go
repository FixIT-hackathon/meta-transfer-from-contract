// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GeneratedABI is the input ABI used to generate the binding from.
const GeneratedABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_FROM_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"transferFromBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"name\":\"list\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"name\":\"unlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Generated is an auto generated Go binding around an Ethereum contract.
type Generated struct {
	GeneratedCaller     // Read-only binding to the contract
	GeneratedTransactor // Write-only binding to the contract
	GeneratedFilterer   // Log filterer for contract events
}

// GeneratedCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneratedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneratedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneratedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneratedSession struct {
	Contract     *Generated        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GeneratedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneratedCallerSession struct {
	Contract *GeneratedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GeneratedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneratedTransactorSession struct {
	Contract     *GeneratedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GeneratedRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneratedRaw struct {
	Contract *Generated // Generic contract binding to access the raw methods on
}

// GeneratedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneratedCallerRaw struct {
	Contract *GeneratedCaller // Generic read-only contract binding to access the raw methods on
}

// GeneratedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneratedTransactorRaw struct {
	Contract *GeneratedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenerated creates a new instance of Generated, bound to a specific deployed contract.
func NewGenerated(address common.Address, backend bind.ContractBackend) (*Generated, error) {
	contract, err := bindGenerated(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Generated{GeneratedCaller: GeneratedCaller{contract: contract}, GeneratedTransactor: GeneratedTransactor{contract: contract}, GeneratedFilterer: GeneratedFilterer{contract: contract}}, nil
}

// NewGeneratedCaller creates a new read-only instance of Generated, bound to a specific deployed contract.
func NewGeneratedCaller(address common.Address, caller bind.ContractCaller) (*GeneratedCaller, error) {
	contract, err := bindGenerated(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneratedCaller{contract: contract}, nil
}

// NewGeneratedTransactor creates a new write-only instance of Generated, bound to a specific deployed contract.
func NewGeneratedTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneratedTransactor, error) {
	contract, err := bindGenerated(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneratedTransactor{contract: contract}, nil
}

// NewGeneratedFilterer creates a new log filterer instance of Generated, bound to a specific deployed contract.
func NewGeneratedFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneratedFilterer, error) {
	contract, err := bindGenerated(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneratedFilterer{contract: contract}, nil
}

// bindGenerated binds a generic wrapper to an already deployed contract.
func bindGenerated(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneratedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Generated *GeneratedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Generated.Contract.GeneratedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Generated *GeneratedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Generated.Contract.GeneratedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Generated *GeneratedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Generated.Contract.GeneratedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Generated *GeneratedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Generated.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Generated *GeneratedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Generated.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Generated *GeneratedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Generated.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "DOMAIN_TYPEHASH")
	return *ret0, err
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Generated.Contract.DOMAINTYPEHASH(&_Generated.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Generated.Contract.DOMAINTYPEHASH(&_Generated.CallOpts)
}

// TRANSFERFROMTYPEHASH is a free data retrieval call binding the contract method 0x8175db48.
//
// Solidity: function TRANSFER_FROM_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedCaller) TRANSFERFROMTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "TRANSFER_FROM_TYPEHASH")
	return *ret0, err
}

// TRANSFERFROMTYPEHASH is a free data retrieval call binding the contract method 0x8175db48.
//
// Solidity: function TRANSFER_FROM_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedSession) TRANSFERFROMTYPEHASH() ([32]byte, error) {
	return _Generated.Contract.TRANSFERFROMTYPEHASH(&_Generated.CallOpts)
}

// TRANSFERFROMTYPEHASH is a free data retrieval call binding the contract method 0x8175db48.
//
// Solidity: function TRANSFER_FROM_TYPEHASH() view returns(bytes32)
func (_Generated *GeneratedCallerSession) TRANSFERFROMTYPEHASH() ([32]byte, error) {
	return _Generated.Contract.TRANSFERFROMTYPEHASH(&_Generated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Generated *GeneratedCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Generated *GeneratedSession) Name() (string, error) {
	return _Generated.Contract.Name(&_Generated.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Generated *GeneratedCallerSession) Name() (string, error) {
	return _Generated.Contract.Name(&_Generated.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Generated *GeneratedCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Generated *GeneratedSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Generated.Contract.Nonces(&_Generated.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Generated *GeneratedCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Generated.Contract.Nonces(&_Generated.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Generated *GeneratedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Generated *GeneratedSession) Owner() (common.Address, error) {
	return _Generated.Contract.Owner(&_Generated.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Generated *GeneratedCallerSession) Owner() (common.Address, error) {
	return _Generated.Contract.Owner(&_Generated.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Generated *GeneratedCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Generated.contract.Call(opts, out, "whitelisted", arg0)
	return *ret0, err
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Generated *GeneratedSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Generated.Contract.Whitelisted(&_Generated.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Generated *GeneratedCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Generated.Contract.Whitelisted(&_Generated.CallOpts, arg0)
}

// List is a paid mutator transaction binding the contract method 0xbb032a66.
//
// Solidity: function list(address erc20) returns()
func (_Generated *GeneratedTransactor) List(opts *bind.TransactOpts, erc20 common.Address) (*types.Transaction, error) {
	return _Generated.contract.Transact(opts, "list", erc20)
}

// List is a paid mutator transaction binding the contract method 0xbb032a66.
//
// Solidity: function list(address erc20) returns()
func (_Generated *GeneratedSession) List(erc20 common.Address) (*types.Transaction, error) {
	return _Generated.Contract.List(&_Generated.TransactOpts, erc20)
}

// List is a paid mutator transaction binding the contract method 0xbb032a66.
//
// Solidity: function list(address erc20) returns()
func (_Generated *GeneratedTransactorSession) List(erc20 common.Address) (*types.Transaction, error) {
	return _Generated.Contract.List(&_Generated.TransactOpts, erc20)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Generated *GeneratedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Generated.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Generated *GeneratedSession) RenounceOwnership() (*types.Transaction, error) {
	return _Generated.Contract.RenounceOwnership(&_Generated.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Generated *GeneratedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Generated.Contract.RenounceOwnership(&_Generated.TransactOpts)
}

// TransferFromBySig is a paid mutator transaction binding the contract method 0x22124fbd.
//
// Solidity: function transferFromBySig(address receiver, address erc20, uint256 amount, uint256 fee, uint256 nonce, bytes32 r, bytes32 s, uint8 v) returns()
func (_Generated *GeneratedTransactor) TransferFromBySig(opts *bind.TransactOpts, receiver common.Address, erc20 common.Address, amount *big.Int, fee *big.Int, nonce *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Generated.contract.Transact(opts, "transferFromBySig", receiver, erc20, amount, fee, nonce, r, s, v)
}

// TransferFromBySig is a paid mutator transaction binding the contract method 0x22124fbd.
//
// Solidity: function transferFromBySig(address receiver, address erc20, uint256 amount, uint256 fee, uint256 nonce, bytes32 r, bytes32 s, uint8 v) returns()
func (_Generated *GeneratedSession) TransferFromBySig(receiver common.Address, erc20 common.Address, amount *big.Int, fee *big.Int, nonce *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Generated.Contract.TransferFromBySig(&_Generated.TransactOpts, receiver, erc20, amount, fee, nonce, r, s, v)
}

// TransferFromBySig is a paid mutator transaction binding the contract method 0x22124fbd.
//
// Solidity: function transferFromBySig(address receiver, address erc20, uint256 amount, uint256 fee, uint256 nonce, bytes32 r, bytes32 s, uint8 v) returns()
func (_Generated *GeneratedTransactorSession) TransferFromBySig(receiver common.Address, erc20 common.Address, amount *big.Int, fee *big.Int, nonce *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Generated.Contract.TransferFromBySig(&_Generated.TransactOpts, receiver, erc20, amount, fee, nonce, r, s, v)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Generated *GeneratedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Generated.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Generated *GeneratedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Generated.Contract.TransferOwnership(&_Generated.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Generated *GeneratedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Generated.Contract.TransferOwnership(&_Generated.TransactOpts, newOwner)
}

// Unlist is a paid mutator transaction binding the contract method 0xbde3672d.
//
// Solidity: function unlist(address erc20) returns()
func (_Generated *GeneratedTransactor) Unlist(opts *bind.TransactOpts, erc20 common.Address) (*types.Transaction, error) {
	return _Generated.contract.Transact(opts, "unlist", erc20)
}

// Unlist is a paid mutator transaction binding the contract method 0xbde3672d.
//
// Solidity: function unlist(address erc20) returns()
func (_Generated *GeneratedSession) Unlist(erc20 common.Address) (*types.Transaction, error) {
	return _Generated.Contract.Unlist(&_Generated.TransactOpts, erc20)
}

// Unlist is a paid mutator transaction binding the contract method 0xbde3672d.
//
// Solidity: function unlist(address erc20) returns()
func (_Generated *GeneratedTransactorSession) Unlist(erc20 common.Address) (*types.Transaction, error) {
	return _Generated.Contract.Unlist(&_Generated.TransactOpts, erc20)
}

// GeneratedExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Generated contract.
type GeneratedExecutedIterator struct {
	Event *GeneratedExecuted // Event containing the contract specifics and raw log

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
func (it *GeneratedExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneratedExecuted)
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
		it.Event = new(GeneratedExecuted)
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
func (it *GeneratedExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneratedExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneratedExecuted represents a Executed event raised by the Generated contract.
type GeneratedExecuted struct {
	Nonce  *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x7271f8b3d48d26ccc839c0f424fe9e12fe4e058108b58c096c078659c7fbadf6.
//
// Solidity: event Executed(uint256 nonce, address signer)
func (_Generated *GeneratedFilterer) FilterExecuted(opts *bind.FilterOpts) (*GeneratedExecutedIterator, error) {

	logs, sub, err := _Generated.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &GeneratedExecutedIterator{contract: _Generated.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x7271f8b3d48d26ccc839c0f424fe9e12fe4e058108b58c096c078659c7fbadf6.
//
// Solidity: event Executed(uint256 nonce, address signer)
func (_Generated *GeneratedFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GeneratedExecuted) (event.Subscription, error) {

	logs, sub, err := _Generated.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneratedExecuted)
				if err := _Generated.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x7271f8b3d48d26ccc839c0f424fe9e12fe4e058108b58c096c078659c7fbadf6.
//
// Solidity: event Executed(uint256 nonce, address signer)
func (_Generated *GeneratedFilterer) ParseExecuted(log types.Log) (*GeneratedExecuted, error) {
	event := new(GeneratedExecuted)
	if err := _Generated.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GeneratedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Generated contract.
type GeneratedOwnershipTransferredIterator struct {
	Event *GeneratedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GeneratedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneratedOwnershipTransferred)
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
		it.Event = new(GeneratedOwnershipTransferred)
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
func (it *GeneratedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneratedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneratedOwnershipTransferred represents a OwnershipTransferred event raised by the Generated contract.
type GeneratedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Generated *GeneratedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GeneratedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Generated.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GeneratedOwnershipTransferredIterator{contract: _Generated.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Generated *GeneratedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GeneratedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Generated.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneratedOwnershipTransferred)
				if err := _Generated.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Generated *GeneratedFilterer) ParseOwnershipTransferred(log types.Log) (*GeneratedOwnershipTransferred, error) {
	event := new(GeneratedOwnershipTransferred)
	if err := _Generated.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
