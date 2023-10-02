// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TxBuilder

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

// DAppApproval is an auto generated low-level Go binding around an user-defined struct.
type DAppApproval struct {
	From            common.Address
	To              common.Address
	Value           *big.Int
	Gas             *big.Int
	MaxFeePerGas    *big.Int
	Nonce           *big.Int
	Deadline        *big.Int
	ControlCodeHash [32]byte
	UserOpHash      [32]byte
	CallChainHash   [32]byte
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
}

// DAppOperation is an auto generated low-level Go binding around an user-defined struct.
type DAppOperation struct {
	To        common.Address
	Approval  DAppApproval
	Signature []byte
}

// SolverCall is an auto generated low-level Go binding around an user-defined struct.
type SolverCall struct {
	From            common.Address
	To              common.Address
	Value           *big.Int
	Gas             *big.Int
	MaxFeePerGas    *big.Int
	Nonce           *big.Int
	Deadline        *big.Int
	ControlCodeHash [32]byte
	UserOpHash      [32]byte
	BidsHash        [32]byte
	Data            []byte
}

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	To        common.Address
	Call      SolverCall
	Signature []byte
	Bids      []BidData
}

// UserCall is an auto generated low-level Go binding around an user-defined struct.
type UserCall struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Control      common.Address
	Data         []byte
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	To        common.Address
	Call      UserCall
	Signature []byte
}

// TxBuilderMetaData contains all meta data concerning the TxBuilder contract.
var TxBuilderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"atlasAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governanceEOA\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"}],\"name\":\"buildDAppOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppApproval\",\"name\":\"approval\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"solverOpData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"solverEOA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"solverContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"buildSolverOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"buildUserOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"control\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"uCall\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBidData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"getControlCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"governanceNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverSigner\",\"type\":\"address\"}],\"name\":\"solverNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TxBuilderABI is the input ABI used to generate the binding from.
// Deprecated: Use TxBuilderMetaData.ABI instead.
var TxBuilderABI = TxBuilderMetaData.ABI

// TxBuilder is an auto generated Go binding around an Ethereum contract.
type TxBuilder struct {
	TxBuilderCaller     // Read-only binding to the contract
	TxBuilderTransactor // Write-only binding to the contract
	TxBuilderFilterer   // Log filterer for contract events
}

// TxBuilderCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxBuilderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxBuilderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxBuilderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxBuilderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxBuilderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxBuilderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxBuilderSession struct {
	Contract     *TxBuilder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxBuilderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxBuilderCallerSession struct {
	Contract *TxBuilderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TxBuilderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxBuilderTransactorSession struct {
	Contract     *TxBuilderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TxBuilderRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxBuilderRaw struct {
	Contract *TxBuilder // Generic contract binding to access the raw methods on
}

// TxBuilderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxBuilderCallerRaw struct {
	Contract *TxBuilderCaller // Generic read-only contract binding to access the raw methods on
}

// TxBuilderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxBuilderTransactorRaw struct {
	Contract *TxBuilderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxBuilder creates a new instance of TxBuilder, bound to a specific deployed contract.
func NewTxBuilder(address common.Address, backend bind.ContractBackend) (*TxBuilder, error) {
	contract, err := bindTxBuilder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxBuilder{TxBuilderCaller: TxBuilderCaller{contract: contract}, TxBuilderTransactor: TxBuilderTransactor{contract: contract}, TxBuilderFilterer: TxBuilderFilterer{contract: contract}}, nil
}

// NewTxBuilderCaller creates a new read-only instance of TxBuilder, bound to a specific deployed contract.
func NewTxBuilderCaller(address common.Address, caller bind.ContractCaller) (*TxBuilderCaller, error) {
	contract, err := bindTxBuilder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxBuilderCaller{contract: contract}, nil
}

// NewTxBuilderTransactor creates a new write-only instance of TxBuilder, bound to a specific deployed contract.
func NewTxBuilderTransactor(address common.Address, transactor bind.ContractTransactor) (*TxBuilderTransactor, error) {
	contract, err := bindTxBuilder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxBuilderTransactor{contract: contract}, nil
}

// NewTxBuilderFilterer creates a new log filterer instance of TxBuilder, bound to a specific deployed contract.
func NewTxBuilderFilterer(address common.Address, filterer bind.ContractFilterer) (*TxBuilderFilterer, error) {
	contract, err := bindTxBuilder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxBuilderFilterer{contract: contract}, nil
}

// bindTxBuilder binds a generic wrapper to an already deployed contract.
func bindTxBuilder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxBuilderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxBuilder *TxBuilderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxBuilder.Contract.TxBuilderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxBuilder *TxBuilderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxBuilder.Contract.TxBuilderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxBuilder *TxBuilderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxBuilder.Contract.TxBuilderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxBuilder *TxBuilderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxBuilder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxBuilder *TxBuilderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxBuilder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxBuilder *TxBuilderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxBuilder.Contract.contract.Transact(opts, method, params...)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_TxBuilder *TxBuilderCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_TxBuilder *TxBuilderSession) Atlas() (common.Address, error) {
	return _TxBuilder.Contract.Atlas(&_TxBuilder.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_TxBuilder *TxBuilderCallerSession) Atlas() (common.Address, error) {
	return _TxBuilder.Contract.Atlas(&_TxBuilder.CallOpts)
}

// BuildDAppOperation is a free data retrieval call binding the contract method 0x7e680394.
//
// Solidity: function buildDAppOperation(address governanceEOA, (address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp)
func (_TxBuilder *TxBuilderCaller) BuildDAppOperation(opts *bind.CallOpts, governanceEOA common.Address, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation) (DAppOperation, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "buildDAppOperation", governanceEOA, dConfig, userOp, solverOps)

	if err != nil {
		return *new(DAppOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppOperation)).(*DAppOperation)

	return out0, err

}

// BuildDAppOperation is a free data retrieval call binding the contract method 0x7e680394.
//
// Solidity: function buildDAppOperation(address governanceEOA, (address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp)
func (_TxBuilder *TxBuilderSession) BuildDAppOperation(governanceEOA common.Address, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation) (DAppOperation, error) {
	return _TxBuilder.Contract.BuildDAppOperation(&_TxBuilder.CallOpts, governanceEOA, dConfig, userOp, solverOps)
}

// BuildDAppOperation is a free data retrieval call binding the contract method 0x7e680394.
//
// Solidity: function buildDAppOperation(address governanceEOA, (address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp)
func (_TxBuilder *TxBuilderCallerSession) BuildDAppOperation(governanceEOA common.Address, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation) (DAppOperation, error) {
	return _TxBuilder.Contract.BuildDAppOperation(&_TxBuilder.CallOpts, governanceEOA, dConfig, userOp, solverOps)
}

// BuildSolverOperation is a free data retrieval call binding the contract method 0x501b8cb1.
//
// Solidity: function buildSolverOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,uint32) dConfig, bytes solverOpData, address solverEOA, address solverContract, uint256 bidAmount) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp)
func (_TxBuilder *TxBuilderCaller) BuildSolverOperation(opts *bind.CallOpts, userOp UserOperation, dConfig DAppConfig, solverOpData []byte, solverEOA common.Address, solverContract common.Address, bidAmount *big.Int) (SolverOperation, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "buildSolverOperation", userOp, dConfig, solverOpData, solverEOA, solverContract, bidAmount)

	if err != nil {
		return *new(SolverOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(SolverOperation)).(*SolverOperation)

	return out0, err

}

// BuildSolverOperation is a free data retrieval call binding the contract method 0x501b8cb1.
//
// Solidity: function buildSolverOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,uint32) dConfig, bytes solverOpData, address solverEOA, address solverContract, uint256 bidAmount) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp)
func (_TxBuilder *TxBuilderSession) BuildSolverOperation(userOp UserOperation, dConfig DAppConfig, solverOpData []byte, solverEOA common.Address, solverContract common.Address, bidAmount *big.Int) (SolverOperation, error) {
	return _TxBuilder.Contract.BuildSolverOperation(&_TxBuilder.CallOpts, userOp, dConfig, solverOpData, solverEOA, solverContract, bidAmount)
}

// BuildSolverOperation is a free data retrieval call binding the contract method 0x501b8cb1.
//
// Solidity: function buildSolverOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,uint32) dConfig, bytes solverOpData, address solverEOA, address solverContract, uint256 bidAmount) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp)
func (_TxBuilder *TxBuilderCallerSession) BuildSolverOperation(userOp UserOperation, dConfig DAppConfig, solverOpData []byte, solverEOA common.Address, solverContract common.Address, bidAmount *big.Int) (SolverOperation, error) {
	return _TxBuilder.Contract.BuildSolverOperation(&_TxBuilder.CallOpts, userOp, dConfig, solverOpData, solverEOA, solverContract, bidAmount)
}

// BuildUserOperation is a free data retrieval call binding the contract method 0x2c7aec7f.
//
// Solidity: function buildUserOperation(address from, address to, uint256 maxFeePerGas, uint256 value, uint256 deadline, bytes data) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp)
func (_TxBuilder *TxBuilderCaller) BuildUserOperation(opts *bind.CallOpts, from common.Address, to common.Address, maxFeePerGas *big.Int, value *big.Int, deadline *big.Int, data []byte) (UserOperation, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "buildUserOperation", from, to, maxFeePerGas, value, deadline, data)

	if err != nil {
		return *new(UserOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(UserOperation)).(*UserOperation)

	return out0, err

}

// BuildUserOperation is a free data retrieval call binding the contract method 0x2c7aec7f.
//
// Solidity: function buildUserOperation(address from, address to, uint256 maxFeePerGas, uint256 value, uint256 deadline, bytes data) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp)
func (_TxBuilder *TxBuilderSession) BuildUserOperation(from common.Address, to common.Address, maxFeePerGas *big.Int, value *big.Int, deadline *big.Int, data []byte) (UserOperation, error) {
	return _TxBuilder.Contract.BuildUserOperation(&_TxBuilder.CallOpts, from, to, maxFeePerGas, value, deadline, data)
}

// BuildUserOperation is a free data retrieval call binding the contract method 0x2c7aec7f.
//
// Solidity: function buildUserOperation(address from, address to, uint256 maxFeePerGas, uint256 value, uint256 deadline, bytes data) view returns((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp)
func (_TxBuilder *TxBuilderCallerSession) BuildUserOperation(from common.Address, to common.Address, maxFeePerGas *big.Int, value *big.Int, deadline *big.Int, data []byte) (UserOperation, error) {
	return _TxBuilder.Contract.BuildUserOperation(&_TxBuilder.CallOpts, from, to, maxFeePerGas, value, deadline, data)
}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_TxBuilder *TxBuilderCaller) Control(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "control")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_TxBuilder *TxBuilderSession) Control() (common.Address, error) {
	return _TxBuilder.Contract.Control(&_TxBuilder.CallOpts)
}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_TxBuilder *TxBuilderCallerSession) Control() (common.Address, error) {
	return _TxBuilder.Contract.Control(&_TxBuilder.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_TxBuilder *TxBuilderCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_TxBuilder *TxBuilderSession) Escrow() (common.Address, error) {
	return _TxBuilder.Contract.Escrow(&_TxBuilder.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_TxBuilder *TxBuilderCallerSession) Escrow() (common.Address, error) {
	return _TxBuilder.Contract.Escrow(&_TxBuilder.CallOpts)
}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_TxBuilder *TxBuilderCaller) Gas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "gas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_TxBuilder *TxBuilderSession) Gas() (*big.Int, error) {
	return _TxBuilder.Contract.Gas(&_TxBuilder.CallOpts)
}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_TxBuilder *TxBuilderCallerSession) Gas() (*big.Int, error) {
	return _TxBuilder.Contract.Gas(&_TxBuilder.CallOpts)
}

// GetBidData is a free data retrieval call binding the contract method 0xd9d1c9d0.
//
// Solidity: function getBidData((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, uint256 amount) view returns((address,uint256)[] bids)
func (_TxBuilder *TxBuilderCaller) GetBidData(opts *bind.CallOpts, uCall UserCall, amount *big.Int) ([]BidData, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "getBidData", uCall, amount)

	if err != nil {
		return *new([]BidData), err
	}

	out0 := *abi.ConvertType(out[0], new([]BidData)).(*[]BidData)

	return out0, err

}

// GetBidData is a free data retrieval call binding the contract method 0xd9d1c9d0.
//
// Solidity: function getBidData((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, uint256 amount) view returns((address,uint256)[] bids)
func (_TxBuilder *TxBuilderSession) GetBidData(uCall UserCall, amount *big.Int) ([]BidData, error) {
	return _TxBuilder.Contract.GetBidData(&_TxBuilder.CallOpts, uCall, amount)
}

// GetBidData is a free data retrieval call binding the contract method 0xd9d1c9d0.
//
// Solidity: function getBidData((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, uint256 amount) view returns((address,uint256)[] bids)
func (_TxBuilder *TxBuilderCallerSession) GetBidData(uCall UserCall, amount *big.Int) ([]BidData, error) {
	return _TxBuilder.Contract.GetBidData(&_TxBuilder.CallOpts, uCall, amount)
}

// GetControlCodeHash is a free data retrieval call binding the contract method 0x68726f96.
//
// Solidity: function getControlCodeHash(address dAppControl) view returns(bytes32)
func (_TxBuilder *TxBuilderCaller) GetControlCodeHash(opts *bind.CallOpts, dAppControl common.Address) ([32]byte, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "getControlCodeHash", dAppControl)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetControlCodeHash is a free data retrieval call binding the contract method 0x68726f96.
//
// Solidity: function getControlCodeHash(address dAppControl) view returns(bytes32)
func (_TxBuilder *TxBuilderSession) GetControlCodeHash(dAppControl common.Address) ([32]byte, error) {
	return _TxBuilder.Contract.GetControlCodeHash(&_TxBuilder.CallOpts, dAppControl)
}

// GetControlCodeHash is a free data retrieval call binding the contract method 0x68726f96.
//
// Solidity: function getControlCodeHash(address dAppControl) view returns(bytes32)
func (_TxBuilder *TxBuilderCallerSession) GetControlCodeHash(dAppControl common.Address) ([32]byte, error) {
	return _TxBuilder.Contract.GetControlCodeHash(&_TxBuilder.CallOpts, dAppControl)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0xac4aab6e.
//
// Solidity: function getDAppConfig() view returns((address,uint32))
func (_TxBuilder *TxBuilderCaller) GetDAppConfig(opts *bind.CallOpts) (DAppConfig, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "getDAppConfig")

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0xac4aab6e.
//
// Solidity: function getDAppConfig() view returns((address,uint32))
func (_TxBuilder *TxBuilderSession) GetDAppConfig() (DAppConfig, error) {
	return _TxBuilder.Contract.GetDAppConfig(&_TxBuilder.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0xac4aab6e.
//
// Solidity: function getDAppConfig() view returns((address,uint32))
func (_TxBuilder *TxBuilderCallerSession) GetDAppConfig() (DAppConfig, error) {
	return _TxBuilder.Contract.GetDAppConfig(&_TxBuilder.CallOpts)
}

// GovernanceNextNonce is a free data retrieval call binding the contract method 0x9dc3ce1c.
//
// Solidity: function governanceNextNonce(address signatory) view returns(uint256)
func (_TxBuilder *TxBuilderCaller) GovernanceNextNonce(opts *bind.CallOpts, signatory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "governanceNextNonce", signatory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GovernanceNextNonce is a free data retrieval call binding the contract method 0x9dc3ce1c.
//
// Solidity: function governanceNextNonce(address signatory) view returns(uint256)
func (_TxBuilder *TxBuilderSession) GovernanceNextNonce(signatory common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.GovernanceNextNonce(&_TxBuilder.CallOpts, signatory)
}

// GovernanceNextNonce is a free data retrieval call binding the contract method 0x9dc3ce1c.
//
// Solidity: function governanceNextNonce(address signatory) view returns(uint256)
func (_TxBuilder *TxBuilderCallerSession) GovernanceNextNonce(signatory common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.GovernanceNextNonce(&_TxBuilder.CallOpts, signatory)
}

// SolverNextNonce is a free data retrieval call binding the contract method 0x7e9e6de2.
//
// Solidity: function solverNextNonce(address solverSigner) view returns(uint256)
func (_TxBuilder *TxBuilderCaller) SolverNextNonce(opts *bind.CallOpts, solverSigner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "solverNextNonce", solverSigner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SolverNextNonce is a free data retrieval call binding the contract method 0x7e9e6de2.
//
// Solidity: function solverNextNonce(address solverSigner) view returns(uint256)
func (_TxBuilder *TxBuilderSession) SolverNextNonce(solverSigner common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.SolverNextNonce(&_TxBuilder.CallOpts, solverSigner)
}

// SolverNextNonce is a free data retrieval call binding the contract method 0x7e9e6de2.
//
// Solidity: function solverNextNonce(address solverSigner) view returns(uint256)
func (_TxBuilder *TxBuilderCallerSession) SolverNextNonce(solverSigner common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.SolverNextNonce(&_TxBuilder.CallOpts, solverSigner)
}

// UserNextNonce is a free data retrieval call binding the contract method 0x0610ada4.
//
// Solidity: function userNextNonce(address user) view returns(uint256)
func (_TxBuilder *TxBuilderCaller) UserNextNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TxBuilder.contract.Call(opts, &out, "userNextNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNextNonce is a free data retrieval call binding the contract method 0x0610ada4.
//
// Solidity: function userNextNonce(address user) view returns(uint256)
func (_TxBuilder *TxBuilderSession) UserNextNonce(user common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.UserNextNonce(&_TxBuilder.CallOpts, user)
}

// UserNextNonce is a free data retrieval call binding the contract method 0x0610ada4.
//
// Solidity: function userNextNonce(address user) view returns(uint256)
func (_TxBuilder *TxBuilderCallerSession) UserNextNonce(user common.Address) (*big.Int, error) {
	return _TxBuilder.Contract.UserNextNonce(&_TxBuilder.CallOpts, user)
}
