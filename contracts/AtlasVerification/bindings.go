// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AtlasVerification

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

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
	BidToken   common.Address
}

// DAppOperation is an auto generated low-level Go binding around an user-defined struct.
type DAppOperation struct {
	From          common.Address
	To            common.Address
	Value         *big.Int
	Gas           *big.Int
	MaxFeePerGas  *big.Int
	Nonce         *big.Int
	Deadline      *big.Int
	Control       common.Address
	UserOpHash    [32]byte
	CallChainHash [32]byte
	Signature     []byte
}

// EscrowAccountData is an auto generated low-level Go binding around an user-defined struct.
type EscrowAccountData struct {
	Balance      *big.Int
	Nonce        uint32
	LastAccessed uint64
}

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	Data         []byte
	Signature    []byte
}

// AtlasVerificationMetaData contains all meta data concerning the AtlasVerification contract.
var AtlasVerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DAppNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDAppControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatoryActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewDAppSignatory\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"addSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"asyncNonceBitIndex\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"LowestEmptyBitmap\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"HighestFullBitmap\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"asyncNonceBitmap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"highestUsedNonce\",\"type\":\"uint8\"},{\"internalType\":\"uint240\",\"name\":\"bitmap\",\"type\":\"uint240\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dapps\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"disableDApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"getDAppOperationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"getGovFromControl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"governanceAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getSolverPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOperationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastUpdate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"initializeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"initializeNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"integrateDApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"removeSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signatories\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"verifyDApp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastAccessed\",\"type\":\"uint64\"}],\"internalType\":\"structEscrowAccountData\",\"name\":\"solverEscrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasWaterMark\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auctionAlreadyComplete\",\"type\":\"bool\"}],\"name\":\"verifySolverOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastAccessed\",\"type\":\"uint64\"}],\"internalType\":\"structEscrowAccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AtlasVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlasVerificationMetaData.ABI instead.
var AtlasVerificationABI = AtlasVerificationMetaData.ABI

// AtlasVerification is an auto generated Go binding around an Ethereum contract.
type AtlasVerification struct {
	AtlasVerificationCaller     // Read-only binding to the contract
	AtlasVerificationTransactor // Write-only binding to the contract
	AtlasVerificationFilterer   // Log filterer for contract events
}

// AtlasVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlasVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlasVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlasVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlasVerificationSession struct {
	Contract     *AtlasVerification // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AtlasVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlasVerificationCallerSession struct {
	Contract *AtlasVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AtlasVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlasVerificationTransactorSession struct {
	Contract     *AtlasVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AtlasVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlasVerificationRaw struct {
	Contract *AtlasVerification // Generic contract binding to access the raw methods on
}

// AtlasVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlasVerificationCallerRaw struct {
	Contract *AtlasVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// AtlasVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlasVerificationTransactorRaw struct {
	Contract *AtlasVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlasVerification creates a new instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerification(address common.Address, backend bind.ContractBackend) (*AtlasVerification, error) {
	contract, err := bindAtlasVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtlasVerification{AtlasVerificationCaller: AtlasVerificationCaller{contract: contract}, AtlasVerificationTransactor: AtlasVerificationTransactor{contract: contract}, AtlasVerificationFilterer: AtlasVerificationFilterer{contract: contract}}, nil
}

// NewAtlasVerificationCaller creates a new read-only instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationCaller(address common.Address, caller bind.ContractCaller) (*AtlasVerificationCaller, error) {
	contract, err := bindAtlasVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationCaller{contract: contract}, nil
}

// NewAtlasVerificationTransactor creates a new write-only instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlasVerificationTransactor, error) {
	contract, err := bindAtlasVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationTransactor{contract: contract}, nil
}

// NewAtlasVerificationFilterer creates a new log filterer instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlasVerificationFilterer, error) {
	contract, err := bindAtlasVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationFilterer{contract: contract}, nil
}

// bindAtlasVerification binds a generic wrapper to an already deployed contract.
func bindAtlasVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtlasVerificationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasVerification *AtlasVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasVerification.Contract.AtlasVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasVerification *AtlasVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AtlasVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasVerification *AtlasVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AtlasVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasVerification *AtlasVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasVerification *AtlasVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasVerification *AtlasVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasVerification.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationSession) ATLAS() (common.Address, error) {
	return _AtlasVerification.Contract.ATLAS(&_AtlasVerification.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationCallerSession) ATLAS() (common.Address, error) {
	return _AtlasVerification.Contract.ATLAS(&_AtlasVerification.CallOpts)
}

// AsyncNonceBitIndex is a free data retrieval call binding the contract method 0xfeba8496.
//
// Solidity: function asyncNonceBitIndex(address ) view returns(uint128 LowestEmptyBitmap, uint128 HighestFullBitmap)
func (_AtlasVerification *AtlasVerificationCaller) AsyncNonceBitIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	LowestEmptyBitmap *big.Int
	HighestFullBitmap *big.Int
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "asyncNonceBitIndex", arg0)

	outstruct := new(struct {
		LowestEmptyBitmap *big.Int
		HighestFullBitmap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LowestEmptyBitmap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HighestFullBitmap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AsyncNonceBitIndex is a free data retrieval call binding the contract method 0xfeba8496.
//
// Solidity: function asyncNonceBitIndex(address ) view returns(uint128 LowestEmptyBitmap, uint128 HighestFullBitmap)
func (_AtlasVerification *AtlasVerificationSession) AsyncNonceBitIndex(arg0 common.Address) (struct {
	LowestEmptyBitmap *big.Int
	HighestFullBitmap *big.Int
}, error) {
	return _AtlasVerification.Contract.AsyncNonceBitIndex(&_AtlasVerification.CallOpts, arg0)
}

// AsyncNonceBitIndex is a free data retrieval call binding the contract method 0xfeba8496.
//
// Solidity: function asyncNonceBitIndex(address ) view returns(uint128 LowestEmptyBitmap, uint128 HighestFullBitmap)
func (_AtlasVerification *AtlasVerificationCallerSession) AsyncNonceBitIndex(arg0 common.Address) (struct {
	LowestEmptyBitmap *big.Int
	HighestFullBitmap *big.Int
}, error) {
	return _AtlasVerification.Contract.AsyncNonceBitIndex(&_AtlasVerification.CallOpts, arg0)
}

// AsyncNonceBitmap is a free data retrieval call binding the contract method 0x5c21ce00.
//
// Solidity: function asyncNonceBitmap(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationCaller) AsyncNonceBitmap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "asyncNonceBitmap", arg0)

	outstruct := new(struct {
		HighestUsedNonce uint8
		Bitmap           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HighestUsedNonce = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Bitmap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AsyncNonceBitmap is a free data retrieval call binding the contract method 0x5c21ce00.
//
// Solidity: function asyncNonceBitmap(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationSession) AsyncNonceBitmap(arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	return _AtlasVerification.Contract.AsyncNonceBitmap(&_AtlasVerification.CallOpts, arg0)
}

// AsyncNonceBitmap is a free data retrieval call binding the contract method 0x5c21ce00.
//
// Solidity: function asyncNonceBitmap(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationCallerSession) AsyncNonceBitmap(arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	return _AtlasVerification.Contract.AsyncNonceBitmap(&_AtlasVerification.CallOpts, arg0)
}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_AtlasVerification *AtlasVerificationCaller) Dapps(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "dapps", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_AtlasVerification *AtlasVerificationSession) Dapps(arg0 [32]byte) ([32]byte, error) {
	return _AtlasVerification.Contract.Dapps(&_AtlasVerification.CallOpts, arg0)
}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_AtlasVerification *AtlasVerificationCallerSession) Dapps(arg0 [32]byte) ([32]byte, error) {
	return _AtlasVerification.Contract.Dapps(&_AtlasVerification.CallOpts, arg0)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AtlasVerification *AtlasVerificationCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AtlasVerification *AtlasVerificationSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AtlasVerification.Contract.Eip712Domain(&_AtlasVerification.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AtlasVerification *AtlasVerificationCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AtlasVerification.Contract.Eip712Domain(&_AtlasVerification.CallOpts)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x8623248e.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetDAppOperationPayload(opts *bind.CallOpts, dAppOp DAppOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDAppOperationPayload", dAppOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x8623248e.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x8623248e.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationSession) GetDomainSeparator() ([32]byte, error) {
	return _AtlasVerification.Contract.GetDomainSeparator(&_AtlasVerification.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _AtlasVerification.Contract.GetDomainSeparator(&_AtlasVerification.CallOpts)
}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address governanceAddress)
func (_AtlasVerification *AtlasVerificationCaller) GetGovFromControl(opts *bind.CallOpts, dAppControl common.Address) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getGovFromControl", dAppControl)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address governanceAddress)
func (_AtlasVerification *AtlasVerificationSession) GetGovFromControl(dAppControl common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, dAppControl)
}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address governanceAddress)
func (_AtlasVerification *AtlasVerificationCallerSession) GetGovFromControl(dAppControl common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, dAppControl)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.GetNextNonce(&_AtlasVerification.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.GetNextNonce(&_AtlasVerification.CallOpts, account)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0xe9244c27.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetSolverPayload(opts *bind.CallOpts, solverOp SolverOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getSolverPayload", solverOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSolverPayload is a free data retrieval call binding the contract method 0xe9244c27.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetSolverPayload(solverOp SolverOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetSolverPayload(&_AtlasVerification.CallOpts, solverOp)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0xe9244c27.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetSolverPayload(solverOp SolverOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetSolverPayload(&_AtlasVerification.CallOpts, solverOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x626099ec.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetUserOperationPayload(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserOperationPayload", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x626099ec.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x626099ec.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// Governance is a free data retrieval call binding the contract method 0x8426e6c1.
//
// Solidity: function governance(address ) view returns(address governance, uint32 callConfig, uint64 lastUpdate)
func (_AtlasVerification *AtlasVerificationCaller) Governance(opts *bind.CallOpts, arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "governance", arg0)

	outstruct := new(struct {
		Governance common.Address
		CallConfig uint32
		LastUpdate uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Governance = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallConfig = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LastUpdate = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Governance is a free data retrieval call binding the contract method 0x8426e6c1.
//
// Solidity: function governance(address ) view returns(address governance, uint32 callConfig, uint64 lastUpdate)
func (_AtlasVerification *AtlasVerificationSession) Governance(arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	return _AtlasVerification.Contract.Governance(&_AtlasVerification.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x8426e6c1.
//
// Solidity: function governance(address ) view returns(address governance, uint32 callConfig, uint64 lastUpdate)
func (_AtlasVerification *AtlasVerificationCallerSession) Governance(arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	return _AtlasVerification.Contract.Governance(&_AtlasVerification.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationCaller) Signatories(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "signatories", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationSession) Signatories(arg0 [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationCallerSession) Signatories(arg0 [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, arg0)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0xf2d66102.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (uint128,uint32,uint64) solverEscrow, uint256 gasWaterMark, bool auctionAlreadyComplete) view returns(uint256 result, uint256 gasLimit, (uint128,uint32,uint64))
func (_AtlasVerification *AtlasVerificationCaller) VerifySolverOp(opts *bind.CallOpts, solverOp SolverOperation, solverEscrow EscrowAccountData, gasWaterMark *big.Int, auctionAlreadyComplete bool) (*big.Int, *big.Int, EscrowAccountData, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "verifySolverOp", solverOp, solverEscrow, gasWaterMark, auctionAlreadyComplete)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(EscrowAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(EscrowAccountData)).(*EscrowAccountData)

	return out0, out1, out2, err

}

// VerifySolverOp is a free data retrieval call binding the contract method 0xf2d66102.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (uint128,uint32,uint64) solverEscrow, uint256 gasWaterMark, bool auctionAlreadyComplete) view returns(uint256 result, uint256 gasLimit, (uint128,uint32,uint64))
func (_AtlasVerification *AtlasVerificationSession) VerifySolverOp(solverOp SolverOperation, solverEscrow EscrowAccountData, gasWaterMark *big.Int, auctionAlreadyComplete bool) (*big.Int, *big.Int, EscrowAccountData, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, solverEscrow, gasWaterMark, auctionAlreadyComplete)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0xf2d66102.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (uint128,uint32,uint64) solverEscrow, uint256 gasWaterMark, bool auctionAlreadyComplete) view returns(uint256 result, uint256 gasLimit, (uint128,uint32,uint64))
func (_AtlasVerification *AtlasVerificationCallerSession) VerifySolverOp(solverOp SolverOperation, solverEscrow EscrowAccountData, gasWaterMark *big.Int, auctionAlreadyComplete bool) (*big.Int, *big.Int, EscrowAccountData, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, solverEscrow, gasWaterMark, auctionAlreadyComplete)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) AddSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "addSignatory", controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactor) DisableDApp(opts *bind.TransactOpts, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "disableDApp", dAppControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationSession) DisableDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) DisableDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationTransactor) InitializeGovernance(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "initializeGovernance", controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, controller)
}

// InitializeNonce is a paid mutator transaction binding the contract method 0xc8cf7971.
//
// Solidity: function initializeNonce(address account) returns()
func (_AtlasVerification *AtlasVerificationTransactor) InitializeNonce(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "initializeNonce", account)
}

// InitializeNonce is a paid mutator transaction binding the contract method 0xc8cf7971.
//
// Solidity: function initializeNonce(address account) returns()
func (_AtlasVerification *AtlasVerificationSession) InitializeNonce(account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeNonce(&_AtlasVerification.TransactOpts, account)
}

// InitializeNonce is a paid mutator transaction binding the contract method 0xc8cf7971.
//
// Solidity: function initializeNonce(address account) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) InitializeNonce(account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeNonce(&_AtlasVerification.TransactOpts, account)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x86ebbbdf.
//
// Solidity: function integrateDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactor) IntegrateDApp(opts *bind.TransactOpts, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "integrateDApp", dAppControl)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x86ebbbdf.
//
// Solidity: function integrateDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationSession) IntegrateDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.IntegrateDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x86ebbbdf.
//
// Solidity: function integrateDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) IntegrateDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.IntegrateDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) RemoveSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "removeSignatory", controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// VerifyDApp is a paid mutator transaction binding the contract method 0x55de74c7.
//
// Solidity: function verifyDApp((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) returns(bool)
func (_AtlasVerification *AtlasVerificationTransactor) VerifyDApp(opts *bind.TransactOpts, dConfig DAppConfig, dAppOp DAppOperation) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "verifyDApp", dConfig, dAppOp)
}

// VerifyDApp is a paid mutator transaction binding the contract method 0x55de74c7.
//
// Solidity: function verifyDApp((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) returns(bool)
func (_AtlasVerification *AtlasVerificationSession) VerifyDApp(dConfig DAppConfig, dAppOp DAppOperation) (*types.Transaction, error) {
	return _AtlasVerification.Contract.VerifyDApp(&_AtlasVerification.TransactOpts, dConfig, dAppOp)
}

// VerifyDApp is a paid mutator transaction binding the contract method 0x55de74c7.
//
// Solidity: function verifyDApp((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) returns(bool)
func (_AtlasVerification *AtlasVerificationTransactorSession) VerifyDApp(dConfig DAppConfig, dAppOp DAppOperation) (*types.Transaction, error) {
	return _AtlasVerification.Contract.VerifyDApp(&_AtlasVerification.TransactOpts, dConfig, dAppOp)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x364e7c78.
//
// Solidity: function verifyUser((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) returns(bool)
func (_AtlasVerification *AtlasVerificationTransactor) VerifyUser(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "verifyUser", dConfig, userOp)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x364e7c78.
//
// Solidity: function verifyUser((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) returns(bool)
func (_AtlasVerification *AtlasVerificationSession) VerifyUser(dConfig DAppConfig, userOp UserOperation) (*types.Transaction, error) {
	return _AtlasVerification.Contract.VerifyUser(&_AtlasVerification.TransactOpts, dConfig, userOp)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x364e7c78.
//
// Solidity: function verifyUser((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) returns(bool)
func (_AtlasVerification *AtlasVerificationTransactorSession) VerifyUser(dConfig DAppConfig, userOp UserOperation) (*types.Transaction, error) {
	return _AtlasVerification.Contract.VerifyUser(&_AtlasVerification.TransactOpts, dConfig, userOp)
}

// AtlasVerificationEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AtlasVerification contract.
type AtlasVerificationEIP712DomainChangedIterator struct {
	Event *AtlasVerificationEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationEIP712DomainChanged)
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
		it.Event = new(AtlasVerificationEIP712DomainChanged)
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
func (it *AtlasVerificationEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationEIP712DomainChanged represents a EIP712DomainChanged event raised by the AtlasVerification contract.
type AtlasVerificationEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AtlasVerification *AtlasVerificationFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AtlasVerificationEIP712DomainChangedIterator, error) {

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationEIP712DomainChangedIterator{contract: _AtlasVerification.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AtlasVerification *AtlasVerificationFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AtlasVerificationEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationEIP712DomainChanged)
				if err := _AtlasVerification.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AtlasVerification *AtlasVerificationFilterer) ParseEIP712DomainChanged(log types.Log) (*AtlasVerificationEIP712DomainChanged, error) {
	event := new(AtlasVerificationEIP712DomainChanged)
	if err := _AtlasVerification.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasVerificationNewDAppSignatoryIterator is returned from FilterNewDAppSignatory and is used to iterate over the raw logs and unpacked data for NewDAppSignatory events raised by the AtlasVerification contract.
type AtlasVerificationNewDAppSignatoryIterator struct {
	Event *AtlasVerificationNewDAppSignatory // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationNewDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationNewDAppSignatory)
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
		it.Event = new(AtlasVerificationNewDAppSignatory)
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
func (it *AtlasVerificationNewDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationNewDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationNewDAppSignatory represents a NewDAppSignatory event raised by the AtlasVerification contract.
type AtlasVerificationNewDAppSignatory struct {
	Controller common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewDAppSignatory is a free log retrieval operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterNewDAppSignatory(opts *bind.FilterOpts, controller []common.Address, governance []common.Address, signatory []common.Address) (*AtlasVerificationNewDAppSignatoryIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationNewDAppSignatoryIterator{contract: _AtlasVerification.contract, event: "NewDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchNewDAppSignatory is a free log subscription operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchNewDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasVerificationNewDAppSignatory, controller []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationNewDAppSignatory)
				if err := _AtlasVerification.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
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

// ParseNewDAppSignatory is a log parse operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) ParseNewDAppSignatory(log types.Log) (*AtlasVerificationNewDAppSignatory, error) {
	event := new(AtlasVerificationNewDAppSignatory)
	if err := _AtlasVerification.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
