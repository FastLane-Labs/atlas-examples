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

// BidData is an auto generated low-level Go binding around an user-defined struct.
type BidData struct {
	Token     common.Address
	BidAmount *big.Int
}

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"WETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"solverOpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraReturnData\",\"type\":\"bytes\"}],\"name\":\"atlasSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"solverMustReimburseGas\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"antecedent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"name\":\"fulfillRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// AtlasSolverCall is a paid mutator transaction binding the contract method 0xa0aa4cb2.
//
// Solidity: function atlasSolverCall(address sender, (address,uint256)[] bids, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, sender common.Address, bids []BidData, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.contract.Transact(opts, "atlasSolverCall", sender, bids, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0xa0aa4cb2.
//
// Solidity: function atlasSolverCall(address sender, (address,uint256)[] bids, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverSession) AtlasSolverCall(sender common.Address, bids []BidData, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.AtlasSolverCall(&_SimpleRFQSolver.TransactOpts, sender, bids, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0xa0aa4cb2.
//
// Solidity: function atlasSolverCall(address sender, (address,uint256)[] bids, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRFQSolver *SimpleRFQSolverTransactorSession) AtlasSolverCall(sender common.Address, bids []BidData, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRFQSolver.Contract.AtlasSolverCall(&_SimpleRFQSolver.TransactOpts, sender, bids, solverOpData, extraReturnData)
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
