// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SimpleRFQSolver

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

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Antecedent common.Address
	Context    []byte
}

// SwapIntent is an auto generated low-level Go binding around an user-defined struct.
type SwapIntent struct {
	TokenUserBuys          common.Address
	AmountUserBuys         *big.Int
	TokenUserSells         common.Address
	AmountUserSells        *big.Int
	AuctionBaseCurrency    common.Address
	SolverMustReimburseGas bool
	Conditions             []Condition
}

// SimpleRFQSolverMetaData contains all meta data concerning the SimpleRFQSolver contract.
var SimpleRFQSolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"solverOpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraReturnData\",\"type\":\"bytes\"}],\"name\":\"atlasSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"solverMustReimburseGas\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"antecedent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"name\":\"fulfillRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SimpleRFQSolverABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleRFQSolverMetaData.ABI instead.
var SimpleRFQSolverABI = SimpleRFQSolverMetaData.ABI

// SimpleRFQSolver is an auto generated Go binding around an Ethereum contract.
type SimpleRFQSolver struct {
	SimpleRFQSolverCaller     // Read-only binding to the contract
	SimpleRFQSolverTransactor // Write-only binding to the contract
	SimpleRFQSolverFilterer   // Log filterer for contract events
}

// SimpleRFQSolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleRFQSolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRFQSolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleRFQSolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRFQSolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleRFQSolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRFQSolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleRFQSolverSession struct {
	Contract     *SimpleRFQSolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRFQSolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleRFQSolverCallerSession struct {
	Contract *SimpleRFQSolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SimpleRFQSolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleRFQSolverTransactorSession struct {
	Contract     *SimpleRFQSolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SimpleRFQSolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRFQSolverRaw struct {
	Contract *SimpleRFQSolver // Generic contract binding to access the raw methods on
}

// SimpleRFQSolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleRFQSolverCallerRaw struct {
	Contract *SimpleRFQSolverCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleRFQSolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleRFQSolverTransactorRaw struct {
	Contract *SimpleRFQSolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleRFQSolver creates a new instance of SimpleRFQSolver, bound to a specific deployed contract.
func NewSimpleRFQSolver(address common.Address, backend bind.ContractBackend) (*SimpleRFQSolver, error) {
	contract, err := bindSimpleRFQSolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolver{SimpleRFQSolverCaller: SimpleRFQSolverCaller{contract: contract}, SimpleRFQSolverTransactor: SimpleRFQSolverTransactor{contract: contract}, SimpleRFQSolverFilterer: SimpleRFQSolverFilterer{contract: contract}}, nil
}

// NewSimpleRFQSolverCaller creates a new read-only instance of SimpleRFQSolver, bound to a specific deployed contract.
func NewSimpleRFQSolverCaller(address common.Address, caller bind.ContractCaller) (*SimpleRFQSolverCaller, error) {
	contract, err := bindSimpleRFQSolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverCaller{contract: contract}, nil
}

// NewSimpleRFQSolverTransactor creates a new write-only instance of SimpleRFQSolver, bound to a specific deployed contract.
func NewSimpleRFQSolverTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleRFQSolverTransactor, error) {
	contract, err := bindSimpleRFQSolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverTransactor{contract: contract}, nil
}

// NewSimpleRFQSolverFilterer creates a new log filterer instance of SimpleRFQSolver, bound to a specific deployed contract.
func NewSimpleRFQSolverFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleRFQSolverFilterer, error) {
	contract, err := bindSimpleRFQSolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverFilterer{contract: contract}, nil
}

// bindSimpleRFQSolver binds a generic wrapper to an already deployed contract.
func bindSimpleRFQSolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleRFQSolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRFQSolver *SimpleRFQSolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRFQSolver.Contract.SimpleRFQSolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRFQSolver *SimpleRFQSolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.SimpleRFQSolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRFQSolver *SimpleRFQSolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.SimpleRFQSolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRFQSolver *SimpleRFQSolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRFQSolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRFQSolver *SimpleRFQSolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRFQSolver *SimpleRFQSolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SimpleRFQSolver.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverSession) ISTEST() (bool, error) {
	return _SimpleRFQSolver.Contract.ISTEST(&_SimpleRFQSolver.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverCallerSession) ISTEST() (bool, error) {
	return _SimpleRFQSolver.Contract.ISTEST(&_SimpleRFQSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRFQSolver *SimpleRFQSolverCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRFQSolver.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRFQSolver *SimpleRFQSolverSession) WETHADDRESS() (common.Address, error) {
	return _SimpleRFQSolver.Contract.WETHADDRESS(&_SimpleRFQSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRFQSolver *SimpleRFQSolverCallerSession) WETHADDRESS() (common.Address, error) {
	return _SimpleRFQSolver.Contract.WETHADDRESS(&_SimpleRFQSolver.CallOpts)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.Transact(opts, "atlasSolverCall", sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverSession) AtlasSolverCall(sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.AtlasSolverCall(&_SimpleRFQSolver.TransactOpts, sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) AtlasSolverCall(sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.AtlasSolverCall(&_SimpleRFQSolver.TransactOpts, sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverSession) Failed() (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Failed(&_SimpleRFQSolver.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) Failed() (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Failed(&_SimpleRFQSolver.TransactOpts)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) FulfillRFQ(opts *bind.TransactOpts, swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.Transact(opts, "fulfillRFQ", swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRFQSolver *SimpleRFQSolverSession) FulfillRFQ(swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.FulfillRFQ(&_SimpleRFQSolver.TransactOpts, swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) FulfillRFQ(swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.FulfillRFQ(&_SimpleRFQSolver.TransactOpts, swapIntent, executionEnvironment)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Fallback(&_SimpleRFQSolver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Fallback(&_SimpleRFQSolver.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverSession) Receive() (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Receive(&_SimpleRFQSolver.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.Receive(&_SimpleRFQSolver.TransactOpts)
}

// SimpleRFQSolverLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogIterator struct {
	Event *SimpleRFQSolverLog // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLog)
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
		it.Event = new(SimpleRFQSolverLog)
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
func (it *SimpleRFQSolverLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLog represents a Log event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLog(opts *bind.FilterOpts) (*SimpleRFQSolverLogIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogIterator{contract: _SimpleRFQSolver.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLog) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLog)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLog(log types.Log) (*SimpleRFQSolverLog, error) {
	event := new(SimpleRFQSolverLog)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogAddressIterator struct {
	Event *SimpleRFQSolverLogAddress // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogAddress)
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
		it.Event = new(SimpleRFQSolverLogAddress)
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
func (it *SimpleRFQSolverLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogAddress represents a LogAddress event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogAddress(opts *bind.FilterOpts) (*SimpleRFQSolverLogAddressIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogAddressIterator{contract: _SimpleRFQSolver.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogAddress) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogAddress)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogAddress(log types.Log) (*SimpleRFQSolverLogAddress, error) {
	event := new(SimpleRFQSolverLogAddress)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArrayIterator struct {
	Event *SimpleRFQSolverLogArray // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogArray)
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
		it.Event = new(SimpleRFQSolverLogArray)
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
func (it *SimpleRFQSolverLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogArray represents a LogArray event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogArray(opts *bind.FilterOpts) (*SimpleRFQSolverLogArrayIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogArrayIterator{contract: _SimpleRFQSolver.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogArray) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogArray)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogArray(log types.Log) (*SimpleRFQSolverLogArray, error) {
	event := new(SimpleRFQSolverLogArray)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArray0Iterator struct {
	Event *SimpleRFQSolverLogArray0 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogArray0)
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
		it.Event = new(SimpleRFQSolverLogArray0)
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
func (it *SimpleRFQSolverLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogArray0 represents a LogArray0 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogArray0(opts *bind.FilterOpts) (*SimpleRFQSolverLogArray0Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogArray0Iterator{contract: _SimpleRFQSolver.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogArray0) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogArray0)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogArray0(log types.Log) (*SimpleRFQSolverLogArray0, error) {
	event := new(SimpleRFQSolverLogArray0)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArray1Iterator struct {
	Event *SimpleRFQSolverLogArray1 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogArray1)
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
		it.Event = new(SimpleRFQSolverLogArray1)
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
func (it *SimpleRFQSolverLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogArray1 represents a LogArray1 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogArray1(opts *bind.FilterOpts) (*SimpleRFQSolverLogArray1Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogArray1Iterator{contract: _SimpleRFQSolver.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogArray1) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogArray1)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogArray1(log types.Log) (*SimpleRFQSolverLogArray1, error) {
	event := new(SimpleRFQSolverLogArray1)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogBytesIterator struct {
	Event *SimpleRFQSolverLogBytes // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogBytes)
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
		it.Event = new(SimpleRFQSolverLogBytes)
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
func (it *SimpleRFQSolverLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogBytes represents a LogBytes event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogBytes(opts *bind.FilterOpts) (*SimpleRFQSolverLogBytesIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogBytesIterator{contract: _SimpleRFQSolver.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogBytes) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogBytes)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogBytes(log types.Log) (*SimpleRFQSolverLogBytes, error) {
	event := new(SimpleRFQSolverLogBytes)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogBytes32Iterator struct {
	Event *SimpleRFQSolverLogBytes32 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogBytes32)
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
		it.Event = new(SimpleRFQSolverLogBytes32)
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
func (it *SimpleRFQSolverLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogBytes32 represents a LogBytes32 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*SimpleRFQSolverLogBytes32Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogBytes32Iterator{contract: _SimpleRFQSolver.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogBytes32) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogBytes32)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogBytes32(log types.Log) (*SimpleRFQSolverLogBytes32, error) {
	event := new(SimpleRFQSolverLogBytes32)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogIntIterator struct {
	Event *SimpleRFQSolverLogInt // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogInt)
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
		it.Event = new(SimpleRFQSolverLogInt)
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
func (it *SimpleRFQSolverLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogInt represents a LogInt event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogInt(opts *bind.FilterOpts) (*SimpleRFQSolverLogIntIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogIntIterator{contract: _SimpleRFQSolver.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogInt)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogInt(log types.Log) (*SimpleRFQSolverLogInt, error) {
	event := new(SimpleRFQSolverLogInt)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedAddressIterator struct {
	Event *SimpleRFQSolverLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedAddress)
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
		it.Event = new(SimpleRFQSolverLogNamedAddress)
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
func (it *SimpleRFQSolverLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedAddress represents a LogNamedAddress event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedAddressIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedAddressIterator{contract: _SimpleRFQSolver.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedAddress)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedAddress(log types.Log) (*SimpleRFQSolverLogNamedAddress, error) {
	event := new(SimpleRFQSolverLogNamedAddress)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArrayIterator struct {
	Event *SimpleRFQSolverLogNamedArray // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedArray)
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
		it.Event = new(SimpleRFQSolverLogNamedArray)
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
func (it *SimpleRFQSolverLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedArray represents a LogNamedArray event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedArrayIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedArrayIterator{contract: _SimpleRFQSolver.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedArray)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedArray(log types.Log) (*SimpleRFQSolverLogNamedArray, error) {
	event := new(SimpleRFQSolverLogNamedArray)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArray0Iterator struct {
	Event *SimpleRFQSolverLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedArray0)
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
		it.Event = new(SimpleRFQSolverLogNamedArray0)
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
func (it *SimpleRFQSolverLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedArray0 represents a LogNamedArray0 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedArray0Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedArray0Iterator{contract: _SimpleRFQSolver.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedArray0)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedArray0(log types.Log) (*SimpleRFQSolverLogNamedArray0, error) {
	event := new(SimpleRFQSolverLogNamedArray0)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArray1Iterator struct {
	Event *SimpleRFQSolverLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedArray1)
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
		it.Event = new(SimpleRFQSolverLogNamedArray1)
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
func (it *SimpleRFQSolverLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedArray1 represents a LogNamedArray1 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedArray1Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedArray1Iterator{contract: _SimpleRFQSolver.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedArray1)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedArray1(log types.Log) (*SimpleRFQSolverLogNamedArray1, error) {
	event := new(SimpleRFQSolverLogNamedArray1)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedBytesIterator struct {
	Event *SimpleRFQSolverLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedBytes)
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
		it.Event = new(SimpleRFQSolverLogNamedBytes)
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
func (it *SimpleRFQSolverLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedBytes represents a LogNamedBytes event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedBytesIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedBytesIterator{contract: _SimpleRFQSolver.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedBytes)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedBytes(log types.Log) (*SimpleRFQSolverLogNamedBytes, error) {
	event := new(SimpleRFQSolverLogNamedBytes)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedBytes32Iterator struct {
	Event *SimpleRFQSolverLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedBytes32)
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
		it.Event = new(SimpleRFQSolverLogNamedBytes32)
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
func (it *SimpleRFQSolverLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedBytes32 represents a LogNamedBytes32 event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedBytes32Iterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedBytes32Iterator{contract: _SimpleRFQSolver.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedBytes32)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedBytes32(log types.Log) (*SimpleRFQSolverLogNamedBytes32, error) {
	event := new(SimpleRFQSolverLogNamedBytes32)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedDecimalIntIterator struct {
	Event *SimpleRFQSolverLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedDecimalInt)
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
		it.Event = new(SimpleRFQSolverLogNamedDecimalInt)
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
func (it *SimpleRFQSolverLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedDecimalIntIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedDecimalIntIterator{contract: _SimpleRFQSolver.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedDecimalInt)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedDecimalInt(log types.Log) (*SimpleRFQSolverLogNamedDecimalInt, error) {
	event := new(SimpleRFQSolverLogNamedDecimalInt)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedDecimalUintIterator struct {
	Event *SimpleRFQSolverLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedDecimalUint)
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
		it.Event = new(SimpleRFQSolverLogNamedDecimalUint)
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
func (it *SimpleRFQSolverLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedDecimalUintIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedDecimalUintIterator{contract: _SimpleRFQSolver.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedDecimalUint)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedDecimalUint(log types.Log) (*SimpleRFQSolverLogNamedDecimalUint, error) {
	event := new(SimpleRFQSolverLogNamedDecimalUint)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedIntIterator struct {
	Event *SimpleRFQSolverLogNamedInt // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedInt)
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
		it.Event = new(SimpleRFQSolverLogNamedInt)
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
func (it *SimpleRFQSolverLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedInt represents a LogNamedInt event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedIntIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedIntIterator{contract: _SimpleRFQSolver.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedInt)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedInt(log types.Log) (*SimpleRFQSolverLogNamedInt, error) {
	event := new(SimpleRFQSolverLogNamedInt)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedStringIterator struct {
	Event *SimpleRFQSolverLogNamedString // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedString)
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
		it.Event = new(SimpleRFQSolverLogNamedString)
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
func (it *SimpleRFQSolverLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedString represents a LogNamedString event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedStringIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedStringIterator{contract: _SimpleRFQSolver.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedString) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedString)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedString(log types.Log) (*SimpleRFQSolverLogNamedString, error) {
	event := new(SimpleRFQSolverLogNamedString)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedUintIterator struct {
	Event *SimpleRFQSolverLogNamedUint // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogNamedUint)
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
		it.Event = new(SimpleRFQSolverLogNamedUint)
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
func (it *SimpleRFQSolverLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogNamedUint represents a LogNamedUint event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*SimpleRFQSolverLogNamedUintIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogNamedUintIterator{contract: _SimpleRFQSolver.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogNamedUint)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogNamedUint(log types.Log) (*SimpleRFQSolverLogNamedUint, error) {
	event := new(SimpleRFQSolverLogNamedUint)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogStringIterator struct {
	Event *SimpleRFQSolverLogString // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogString)
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
		it.Event = new(SimpleRFQSolverLogString)
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
func (it *SimpleRFQSolverLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogString represents a LogString event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogString(opts *bind.FilterOpts) (*SimpleRFQSolverLogStringIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogStringIterator{contract: _SimpleRFQSolver.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogString) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogString)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogString(log types.Log) (*SimpleRFQSolverLogString, error) {
	event := new(SimpleRFQSolverLogString)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogUintIterator struct {
	Event *SimpleRFQSolverLogUint // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogUint)
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
		it.Event = new(SimpleRFQSolverLogUint)
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
func (it *SimpleRFQSolverLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogUint represents a LogUint event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogUint(opts *bind.FilterOpts) (*SimpleRFQSolverLogUintIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogUintIterator{contract: _SimpleRFQSolver.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogUint)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogUint(log types.Log) (*SimpleRFQSolverLogUint, error) {
	event := new(SimpleRFQSolverLogUint)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRFQSolverLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogsIterator struct {
	Event *SimpleRFQSolverLogs // Event containing the contract specifics and raw log

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
func (it *SimpleRFQSolverLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRFQSolverLogs)
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
		it.Event = new(SimpleRFQSolverLogs)
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
func (it *SimpleRFQSolverLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRFQSolverLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRFQSolverLogs represents a Logs event raised by the SimpleRFQSolver contract.
type SimpleRFQSolverLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) FilterLogs(opts *bind.FilterOpts) (*SimpleRFQSolverLogsIterator, error) {

	logs, sub, err := _SimpleRFQSolver.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &SimpleRFQSolverLogsIterator{contract: _SimpleRFQSolver.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *SimpleRFQSolverLogs) (event.Subscription, error) {

	logs, sub, err := _SimpleRFQSolver.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRFQSolverLogs)
				if err := _SimpleRFQSolver.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRFQSolver *SimpleRFQSolverFilterer) ParseLogs(log types.Log) (*SimpleRFQSolverLogs, error) {
	event := new(SimpleRFQSolverLogs)
	if err := _SimpleRFQSolver.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
