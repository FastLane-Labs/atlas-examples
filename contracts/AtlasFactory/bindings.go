// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AtlasFactory

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

// AtlasFactoryMetaData contains all meta data concerning the AtlasFactory contract.
var AtlasFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"environment\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewExecutionEnvironment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"createExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"getExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"getExecutionEnvironmentCustom\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"}],\"name\":\"getMimicCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"creationCode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AtlasFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlasFactoryMetaData.ABI instead.
var AtlasFactoryABI = AtlasFactoryMetaData.ABI

// AtlasFactory is an auto generated Go binding around an Ethereum contract.
type AtlasFactory struct {
	AtlasFactoryCaller     // Read-only binding to the contract
	AtlasFactoryTransactor // Write-only binding to the contract
	AtlasFactoryFilterer   // Log filterer for contract events
}

// AtlasFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlasFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlasFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlasFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlasFactorySession struct {
	Contract     *AtlasFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlasFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlasFactoryCallerSession struct {
	Contract *AtlasFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AtlasFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlasFactoryTransactorSession struct {
	Contract     *AtlasFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AtlasFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlasFactoryRaw struct {
	Contract *AtlasFactory // Generic contract binding to access the raw methods on
}

// AtlasFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlasFactoryCallerRaw struct {
	Contract *AtlasFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AtlasFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlasFactoryTransactorRaw struct {
	Contract *AtlasFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlasFactory creates a new instance of AtlasFactory, bound to a specific deployed contract.
func NewAtlasFactory(address common.Address, backend bind.ContractBackend) (*AtlasFactory, error) {
	contract, err := bindAtlasFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtlasFactory{AtlasFactoryCaller: AtlasFactoryCaller{contract: contract}, AtlasFactoryTransactor: AtlasFactoryTransactor{contract: contract}, AtlasFactoryFilterer: AtlasFactoryFilterer{contract: contract}}, nil
}

// NewAtlasFactoryCaller creates a new read-only instance of AtlasFactory, bound to a specific deployed contract.
func NewAtlasFactoryCaller(address common.Address, caller bind.ContractCaller) (*AtlasFactoryCaller, error) {
	contract, err := bindAtlasFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasFactoryCaller{contract: contract}, nil
}

// NewAtlasFactoryTransactor creates a new write-only instance of AtlasFactory, bound to a specific deployed contract.
func NewAtlasFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlasFactoryTransactor, error) {
	contract, err := bindAtlasFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasFactoryTransactor{contract: contract}, nil
}

// NewAtlasFactoryFilterer creates a new log filterer instance of AtlasFactory, bound to a specific deployed contract.
func NewAtlasFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlasFactoryFilterer, error) {
	contract, err := bindAtlasFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlasFactoryFilterer{contract: contract}, nil
}

// bindAtlasFactory binds a generic wrapper to an already deployed contract.
func bindAtlasFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtlasFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasFactory *AtlasFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasFactory.Contract.AtlasFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasFactory *AtlasFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasFactory.Contract.AtlasFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasFactory *AtlasFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasFactory.Contract.AtlasFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasFactory *AtlasFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasFactory *AtlasFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasFactory *AtlasFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasFactory.Contract.contract.Transact(opts, method, params...)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_AtlasFactory *AtlasFactoryCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_AtlasFactory *AtlasFactorySession) Atlas() (common.Address, error) {
	return _AtlasFactory.Contract.Atlas(&_AtlasFactory.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_AtlasFactory *AtlasFactoryCallerSession) Atlas() (common.Address, error) {
	return _AtlasFactory.Contract.Atlas(&_AtlasFactory.CallOpts)
}

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_AtlasFactory *AtlasFactoryCaller) ExecutionTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "executionTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_AtlasFactory *AtlasFactorySession) ExecutionTemplate() (common.Address, error) {
	return _AtlasFactory.Contract.ExecutionTemplate(&_AtlasFactory.CallOpts)
}

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_AtlasFactory *AtlasFactoryCallerSession) ExecutionTemplate() (common.Address, error) {
	return _AtlasFactory.Contract.ExecutionTemplate(&_AtlasFactory.CallOpts)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_AtlasFactory *AtlasFactoryCaller) GetExecutionEnvironment(opts *bind.CallOpts, user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "getExecutionEnvironment", user, dAppControl)

	outstruct := new(struct {
		ExecutionEnvironment common.Address
		CallConfig           uint32
		Exists               bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExecutionEnvironment = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallConfig = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_AtlasFactory *AtlasFactorySession) GetExecutionEnvironment(user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _AtlasFactory.Contract.GetExecutionEnvironment(&_AtlasFactory.CallOpts, user, dAppControl)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_AtlasFactory *AtlasFactoryCallerSession) GetExecutionEnvironment(user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _AtlasFactory.Contract.GetExecutionEnvironment(&_AtlasFactory.CallOpts, user, dAppControl)
}

// GetExecutionEnvironmentCustom is a free data retrieval call binding the contract method 0xd0c2c091.
//
// Solidity: function getExecutionEnvironmentCustom(address user, bytes32 controlCodeHash, address controller, uint32 callConfig) view returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactoryCaller) GetExecutionEnvironmentCustom(opts *bind.CallOpts, user common.Address, controlCodeHash [32]byte, controller common.Address, callConfig uint32) (common.Address, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "getExecutionEnvironmentCustom", user, controlCodeHash, controller, callConfig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionEnvironmentCustom is a free data retrieval call binding the contract method 0xd0c2c091.
//
// Solidity: function getExecutionEnvironmentCustom(address user, bytes32 controlCodeHash, address controller, uint32 callConfig) view returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactorySession) GetExecutionEnvironmentCustom(user common.Address, controlCodeHash [32]byte, controller common.Address, callConfig uint32) (common.Address, error) {
	return _AtlasFactory.Contract.GetExecutionEnvironmentCustom(&_AtlasFactory.CallOpts, user, controlCodeHash, controller, callConfig)
}

// GetExecutionEnvironmentCustom is a free data retrieval call binding the contract method 0xd0c2c091.
//
// Solidity: function getExecutionEnvironmentCustom(address user, bytes32 controlCodeHash, address controller, uint32 callConfig) view returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactoryCallerSession) GetExecutionEnvironmentCustom(user common.Address, controlCodeHash [32]byte, controller common.Address, callConfig uint32) (common.Address, error) {
	return _AtlasFactory.Contract.GetExecutionEnvironmentCustom(&_AtlasFactory.CallOpts, user, controlCodeHash, controller, callConfig)
}

// GetMimicCreationCode is a free data retrieval call binding the contract method 0xa2097838.
//
// Solidity: function getMimicCreationCode(address controller, uint32 callConfig, address user, bytes32 controlCodeHash) view returns(bytes creationCode)
func (_AtlasFactory *AtlasFactoryCaller) GetMimicCreationCode(opts *bind.CallOpts, controller common.Address, callConfig uint32, user common.Address, controlCodeHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "getMimicCreationCode", controller, callConfig, user, controlCodeHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMimicCreationCode is a free data retrieval call binding the contract method 0xa2097838.
//
// Solidity: function getMimicCreationCode(address controller, uint32 callConfig, address user, bytes32 controlCodeHash) view returns(bytes creationCode)
func (_AtlasFactory *AtlasFactorySession) GetMimicCreationCode(controller common.Address, callConfig uint32, user common.Address, controlCodeHash [32]byte) ([]byte, error) {
	return _AtlasFactory.Contract.GetMimicCreationCode(&_AtlasFactory.CallOpts, controller, callConfig, user, controlCodeHash)
}

// GetMimicCreationCode is a free data retrieval call binding the contract method 0xa2097838.
//
// Solidity: function getMimicCreationCode(address controller, uint32 callConfig, address user, bytes32 controlCodeHash) view returns(bytes creationCode)
func (_AtlasFactory *AtlasFactoryCallerSession) GetMimicCreationCode(controller common.Address, callConfig uint32, user common.Address, controlCodeHash [32]byte) ([]byte, error) {
	return _AtlasFactory.Contract.GetMimicCreationCode(&_AtlasFactory.CallOpts, controller, callConfig, user, controlCodeHash)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_AtlasFactory *AtlasFactoryCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AtlasFactory.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_AtlasFactory *AtlasFactorySession) Salt() ([32]byte, error) {
	return _AtlasFactory.Contract.Salt(&_AtlasFactory.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_AtlasFactory *AtlasFactoryCallerSession) Salt() ([32]byte, error) {
	return _AtlasFactory.Contract.Salt(&_AtlasFactory.CallOpts)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x7e29c684.
//
// Solidity: function createExecutionEnvironment(address account, address dAppControl) returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactoryTransactor) CreateExecutionEnvironment(opts *bind.TransactOpts, account common.Address, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasFactory.contract.Transact(opts, "createExecutionEnvironment", account, dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x7e29c684.
//
// Solidity: function createExecutionEnvironment(address account, address dAppControl) returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactorySession) CreateExecutionEnvironment(account common.Address, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasFactory.Contract.CreateExecutionEnvironment(&_AtlasFactory.TransactOpts, account, dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x7e29c684.
//
// Solidity: function createExecutionEnvironment(address account, address dAppControl) returns(address executionEnvironment)
func (_AtlasFactory *AtlasFactoryTransactorSession) CreateExecutionEnvironment(account common.Address, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasFactory.Contract.CreateExecutionEnvironment(&_AtlasFactory.TransactOpts, account, dAppControl)
}

// AtlasFactoryNewExecutionEnvironmentIterator is returned from FilterNewExecutionEnvironment and is used to iterate over the raw logs and unpacked data for NewExecutionEnvironment events raised by the AtlasFactory contract.
type AtlasFactoryNewExecutionEnvironmentIterator struct {
	Event *AtlasFactoryNewExecutionEnvironment // Event containing the contract specifics and raw log

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
func (it *AtlasFactoryNewExecutionEnvironmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasFactoryNewExecutionEnvironment)
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
		it.Event = new(AtlasFactoryNewExecutionEnvironment)
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
func (it *AtlasFactoryNewExecutionEnvironmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasFactoryNewExecutionEnvironmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasFactoryNewExecutionEnvironment represents a NewExecutionEnvironment event raised by the AtlasFactory contract.
type AtlasFactoryNewExecutionEnvironment struct {
	Environment common.Address
	User        common.Address
	Controller  common.Address
	CallConfig  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionEnvironment is a free log retrieval operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_AtlasFactory *AtlasFactoryFilterer) FilterNewExecutionEnvironment(opts *bind.FilterOpts, environment []common.Address, user []common.Address, controller []common.Address) (*AtlasFactoryNewExecutionEnvironmentIterator, error) {

	var environmentRule []interface{}
	for _, environmentItem := range environment {
		environmentRule = append(environmentRule, environmentItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _AtlasFactory.contract.FilterLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasFactoryNewExecutionEnvironmentIterator{contract: _AtlasFactory.contract, event: "NewExecutionEnvironment", logs: logs, sub: sub}, nil
}

// WatchNewExecutionEnvironment is a free log subscription operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_AtlasFactory *AtlasFactoryFilterer) WatchNewExecutionEnvironment(opts *bind.WatchOpts, sink chan<- *AtlasFactoryNewExecutionEnvironment, environment []common.Address, user []common.Address, controller []common.Address) (event.Subscription, error) {

	var environmentRule []interface{}
	for _, environmentItem := range environment {
		environmentRule = append(environmentRule, environmentItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _AtlasFactory.contract.WatchLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasFactoryNewExecutionEnvironment)
				if err := _AtlasFactory.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
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

// ParseNewExecutionEnvironment is a log parse operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_AtlasFactory *AtlasFactoryFilterer) ParseNewExecutionEnvironment(log types.Log) (*AtlasFactoryNewExecutionEnvironment, error) {
	event := new(AtlasFactoryNewExecutionEnvironment)
	if err := _AtlasFactory.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
