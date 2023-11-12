// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SwapIntentController

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

// CallConfig is an auto generated low-level Go binding around an user-defined struct.
type CallConfig struct {
	Sequenced                        bool
	RequirePreOps                    bool
	TrackPreOpsReturnData            bool
	TrackUserReturnData              bool
	DelegateUser                     bool
	LocalUser                        bool
	PreSolver                        bool
	PostSolver                       bool
	RequirePostOps                   bool
	ZeroSolvers                      bool
	ReuseUserOp                      bool
	UserBundler                      bool
	SolverBundler                    bool
	VerifySolverBundlerCallChainHash bool
	UnknownBundler                   bool
	ForwardReturnData                bool
}

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Antecedent common.Address
	Context    []byte
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
	BidToken   common.Address
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

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
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

// SwapIntentControllerMetaData contains all meta data concerning the SwapIntentController contract.
var SwapIntentControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_escrow\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXPECTED_GAS_USAGE_EX_SOLVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_USER_CONDITIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_CONDITION_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"allocateValueCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"control\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getBidFormat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getBidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"sequenced\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackUserReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"delegateUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"localUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"postSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePostOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"zeroSolvers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reuseUserOp\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userBundler\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"solverBundler\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"verifySolverBundlerCallChainHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unknownBundler\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forwardReturnData\",\"type\":\"bool\"}],\"internalType\":\"structCallConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getDAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppSignatory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"governanceAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"postOpsCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"postSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"preOpsCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"preSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequencedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequenced\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"source\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"solverMustReimburseGas\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"antecedent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"}],\"internalType\":\"structSwapData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delegated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapIntentControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapIntentControllerMetaData.ABI instead.
var SwapIntentControllerABI = SwapIntentControllerMetaData.ABI

// SwapIntentController is an auto generated Go binding around an Ethereum contract.
type SwapIntentController struct {
	SwapIntentControllerCaller     // Read-only binding to the contract
	SwapIntentControllerTransactor // Write-only binding to the contract
	SwapIntentControllerFilterer   // Log filterer for contract events
}

// SwapIntentControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapIntentControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapIntentControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapIntentControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapIntentControllerSession struct {
	Contract     *SwapIntentController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapIntentControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapIntentControllerCallerSession struct {
	Contract *SwapIntentControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SwapIntentControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapIntentControllerTransactorSession struct {
	Contract     *SwapIntentControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SwapIntentControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapIntentControllerRaw struct {
	Contract *SwapIntentController // Generic contract binding to access the raw methods on
}

// SwapIntentControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapIntentControllerCallerRaw struct {
	Contract *SwapIntentControllerCaller // Generic read-only contract binding to access the raw methods on
}

// SwapIntentControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapIntentControllerTransactorRaw struct {
	Contract *SwapIntentControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapIntentController creates a new instance of SwapIntentController, bound to a specific deployed contract.
func NewSwapIntentController(address common.Address, backend bind.ContractBackend) (*SwapIntentController, error) {
	contract, err := bindSwapIntentController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapIntentController{SwapIntentControllerCaller: SwapIntentControllerCaller{contract: contract}, SwapIntentControllerTransactor: SwapIntentControllerTransactor{contract: contract}, SwapIntentControllerFilterer: SwapIntentControllerFilterer{contract: contract}}, nil
}

// NewSwapIntentControllerCaller creates a new read-only instance of SwapIntentController, bound to a specific deployed contract.
func NewSwapIntentControllerCaller(address common.Address, caller bind.ContractCaller) (*SwapIntentControllerCaller, error) {
	contract, err := bindSwapIntentController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerCaller{contract: contract}, nil
}

// NewSwapIntentControllerTransactor creates a new write-only instance of SwapIntentController, bound to a specific deployed contract.
func NewSwapIntentControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapIntentControllerTransactor, error) {
	contract, err := bindSwapIntentController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerTransactor{contract: contract}, nil
}

// NewSwapIntentControllerFilterer creates a new log filterer instance of SwapIntentController, bound to a specific deployed contract.
func NewSwapIntentControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapIntentControllerFilterer, error) {
	contract, err := bindSwapIntentController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerFilterer{contract: contract}, nil
}

// bindSwapIntentController binds a generic wrapper to an already deployed contract.
func bindSwapIntentController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapIntentControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapIntentController *SwapIntentControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapIntentController.Contract.SwapIntentControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapIntentController *SwapIntentControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentController.Contract.SwapIntentControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapIntentController *SwapIntentControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapIntentController.Contract.SwapIntentControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapIntentController *SwapIntentControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapIntentController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapIntentController *SwapIntentControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapIntentController *SwapIntentControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapIntentController.Contract.contract.Transact(opts, method, params...)
}

// EXPECTEDGASUSAGEEXSOLVER is a free data retrieval call binding the contract method 0x7775c3e1.
//
// Solidity: function EXPECTED_GAS_USAGE_EX_SOLVER() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCaller) EXPECTEDGASUSAGEEXSOLVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "EXPECTED_GAS_USAGE_EX_SOLVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPECTEDGASUSAGEEXSOLVER is a free data retrieval call binding the contract method 0x7775c3e1.
//
// Solidity: function EXPECTED_GAS_USAGE_EX_SOLVER() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerSession) EXPECTEDGASUSAGEEXSOLVER() (*big.Int, error) {
	return _SwapIntentController.Contract.EXPECTEDGASUSAGEEXSOLVER(&_SwapIntentController.CallOpts)
}

// EXPECTEDGASUSAGEEXSOLVER is a free data retrieval call binding the contract method 0x7775c3e1.
//
// Solidity: function EXPECTED_GAS_USAGE_EX_SOLVER() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCallerSession) EXPECTEDGASUSAGEEXSOLVER() (*big.Int, error) {
	return _SwapIntentController.Contract.EXPECTEDGASUSAGEEXSOLVER(&_SwapIntentController.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SwapIntentController *SwapIntentControllerCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SwapIntentController *SwapIntentControllerSession) ISTEST() (bool, error) {
	return _SwapIntentController.Contract.ISTEST(&_SwapIntentController.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SwapIntentController *SwapIntentControllerCallerSession) ISTEST() (bool, error) {
	return _SwapIntentController.Contract.ISTEST(&_SwapIntentController.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCaller) MAXUSERCONDITIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "MAX_USER_CONDITIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _SwapIntentController.Contract.MAXUSERCONDITIONS(&_SwapIntentController.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCallerSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _SwapIntentController.Contract.MAXUSERCONDITIONS(&_SwapIntentController.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCaller) USERCONDITIONGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "USER_CONDITION_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _SwapIntentController.Contract.USERCONDITIONGASLIMIT(&_SwapIntentController.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentController *SwapIntentControllerCallerSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _SwapIntentController.Contract.USERCONDITIONGASLIMIT(&_SwapIntentController.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_SwapIntentController *SwapIntentControllerCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_SwapIntentController *SwapIntentControllerSession) Atlas() (common.Address, error) {
	return _SwapIntentController.Contract.Atlas(&_SwapIntentController.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_SwapIntentController *SwapIntentControllerCallerSession) Atlas() (common.Address, error) {
	return _SwapIntentController.Contract.Atlas(&_SwapIntentController.CallOpts)
}

// CallConfig is a free data retrieval call binding the contract method 0xb6834538.
//
// Solidity: function callConfig() view returns(uint32)
func (_SwapIntentController *SwapIntentControllerCaller) CallConfig(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "callConfig")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CallConfig is a free data retrieval call binding the contract method 0xb6834538.
//
// Solidity: function callConfig() view returns(uint32)
func (_SwapIntentController *SwapIntentControllerSession) CallConfig() (uint32, error) {
	return _SwapIntentController.Contract.CallConfig(&_SwapIntentController.CallOpts)
}

// CallConfig is a free data retrieval call binding the contract method 0xb6834538.
//
// Solidity: function callConfig() view returns(uint32)
func (_SwapIntentController *SwapIntentControllerCallerSession) CallConfig() (uint32, error) {
	return _SwapIntentController.Contract.CallConfig(&_SwapIntentController.CallOpts)
}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_SwapIntentController *SwapIntentControllerCaller) Control(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "control")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_SwapIntentController *SwapIntentControllerSession) Control() (common.Address, error) {
	return _SwapIntentController.Contract.Control(&_SwapIntentController.CallOpts)
}

// Control is a free data retrieval call binding the contract method 0xd8de6587.
//
// Solidity: function control() view returns(address)
func (_SwapIntentController *SwapIntentControllerCallerSession) Control() (common.Address, error) {
	return _SwapIntentController.Contract.Control(&_SwapIntentController.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_SwapIntentController *SwapIntentControllerCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_SwapIntentController *SwapIntentControllerSession) Escrow() (common.Address, error) {
	return _SwapIntentController.Contract.Escrow(&_SwapIntentController.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_SwapIntentController *SwapIntentControllerCallerSession) Escrow() (common.Address, error) {
	return _SwapIntentController.Contract.Escrow(&_SwapIntentController.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x09dbeb08.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentController *SwapIntentControllerCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x09dbeb08.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentController *SwapIntentControllerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _SwapIntentController.Contract.GetBidFormat(&_SwapIntentController.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x09dbeb08.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentController *SwapIntentControllerCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _SwapIntentController.Contract.GetBidFormat(&_SwapIntentController.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0xc14b782d.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentController *SwapIntentControllerCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0xc14b782d.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentController *SwapIntentControllerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _SwapIntentController.Contract.GetBidValue(&_SwapIntentController.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0xc14b782d.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentController *SwapIntentControllerCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _SwapIntentController.Contract.GetBidValue(&_SwapIntentController.CallOpts, solverOp)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentController *SwapIntentControllerCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentController *SwapIntentControllerSession) GetCallConfig() (CallConfig, error) {
	return _SwapIntentController.Contract.GetCallConfig(&_SwapIntentController.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentController *SwapIntentControllerCallerSession) GetCallConfig() (CallConfig, error) {
	return _SwapIntentController.Contract.GetCallConfig(&_SwapIntentController.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x18fca270.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_SwapIntentController *SwapIntentControllerCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x18fca270.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_SwapIntentController *SwapIntentControllerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _SwapIntentController.Contract.GetDAppConfig(&_SwapIntentController.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x18fca270.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_SwapIntentController *SwapIntentControllerCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _SwapIntentController.Contract.GetDAppConfig(&_SwapIntentController.CallOpts, userOp)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address governanceAddress)
func (_SwapIntentController *SwapIntentControllerCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address governanceAddress)
func (_SwapIntentController *SwapIntentControllerSession) GetDAppSignatory() (common.Address, error) {
	return _SwapIntentController.Contract.GetDAppSignatory(&_SwapIntentController.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address governanceAddress)
func (_SwapIntentController *SwapIntentControllerCallerSession) GetDAppSignatory() (common.Address, error) {
	return _SwapIntentController.Contract.GetDAppSignatory(&_SwapIntentController.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentController *SwapIntentControllerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentController *SwapIntentControllerSession) Governance() (common.Address, error) {
	return _SwapIntentController.Contract.Governance(&_SwapIntentController.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentController *SwapIntentControllerCallerSession) Governance() (common.Address, error) {
	return _SwapIntentController.Contract.Governance(&_SwapIntentController.CallOpts)
}

// RequireSequencedNonces is a free data retrieval call binding the contract method 0xc255cf22.
//
// Solidity: function requireSequencedNonces() view returns(bool isSequenced)
func (_SwapIntentController *SwapIntentControllerCaller) RequireSequencedNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "requireSequencedNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequencedNonces is a free data retrieval call binding the contract method 0xc255cf22.
//
// Solidity: function requireSequencedNonces() view returns(bool isSequenced)
func (_SwapIntentController *SwapIntentControllerSession) RequireSequencedNonces() (bool, error) {
	return _SwapIntentController.Contract.RequireSequencedNonces(&_SwapIntentController.CallOpts)
}

// RequireSequencedNonces is a free data retrieval call binding the contract method 0xc255cf22.
//
// Solidity: function requireSequencedNonces() view returns(bool isSequenced)
func (_SwapIntentController *SwapIntentControllerCallerSession) RequireSequencedNonces() (bool, error) {
	return _SwapIntentController.Contract.RequireSequencedNonces(&_SwapIntentController.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_SwapIntentController *SwapIntentControllerCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_SwapIntentController *SwapIntentControllerSession) Salt() ([32]byte, error) {
	return _SwapIntentController.Contract.Salt(&_SwapIntentController.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_SwapIntentController *SwapIntentControllerCallerSession) Salt() ([32]byte, error) {
	return _SwapIntentController.Contract.Salt(&_SwapIntentController.CallOpts)
}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_SwapIntentController *SwapIntentControllerCaller) Source(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "source")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_SwapIntentController *SwapIntentControllerSession) Source() (common.Address, error) {
	return _SwapIntentController.Contract.Source(&_SwapIntentController.CallOpts)
}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_SwapIntentController *SwapIntentControllerCallerSession) Source() (common.Address, error) {
	return _SwapIntentController.Contract.Source(&_SwapIntentController.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentController *SwapIntentControllerCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentController.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentController *SwapIntentControllerSession) UserDelegated() (bool, error) {
	return _SwapIntentController.Contract.UserDelegated(&_SwapIntentController.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentController *SwapIntentControllerCallerSession) UserDelegated() (bool, error) {
	return _SwapIntentController.Contract.UserDelegated(&_SwapIntentController.CallOpts)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentController *SwapIntentControllerTransactor) AllocateValueCall(opts *bind.TransactOpts, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "allocateValueCall", bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentController *SwapIntentControllerSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.AllocateValueCall(&_SwapIntentController.TransactOpts, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentController *SwapIntentControllerTransactorSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.AllocateValueCall(&_SwapIntentController.TransactOpts, bidToken, bidAmount, data)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SwapIntentController *SwapIntentControllerSession) Failed() (*types.Transaction, error) {
	return _SwapIntentController.Contract.Failed(&_SwapIntentController.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactorSession) Failed() (*types.Transaction, error) {
	return _SwapIntentController.Contract.Failed(&_SwapIntentController.TransactOpts)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x6585e19a.
//
// Solidity: function postOpsCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactor) PostOpsCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "postOpsCall", data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x6585e19a.
//
// Solidity: function postOpsCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerSession) PostOpsCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PostOpsCall(&_SwapIntentController.TransactOpts, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x6585e19a.
//
// Solidity: function postOpsCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactorSession) PostOpsCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PostOpsCall(&_SwapIntentController.TransactOpts, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x0f523813.
//
// Solidity: function postSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactor) PostSolverCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "postSolverCall", data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x0f523813.
//
// Solidity: function postSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerSession) PostSolverCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PostSolverCall(&_SwapIntentController.TransactOpts, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x0f523813.
//
// Solidity: function postSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactorSession) PostSolverCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PostSolverCall(&_SwapIntentController.TransactOpts, data)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xf1f4e7d5.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentController *SwapIntentControllerTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xf1f4e7d5.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentController *SwapIntentControllerSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PreOpsCall(&_SwapIntentController.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xf1f4e7d5.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentController *SwapIntentControllerTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PreOpsCall(&_SwapIntentController.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x59244761.
//
// Solidity: function preSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactor) PreSolverCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "preSolverCall", data)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x59244761.
//
// Solidity: function preSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerSession) PreSolverCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PreSolverCall(&_SwapIntentController.TransactOpts, data)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x59244761.
//
// Solidity: function preSolverCall(bytes data) payable returns(bool)
func (_SwapIntentController *SwapIntentControllerTransactorSession) PreSolverCall(data []byte) (*types.Transaction, error) {
	return _SwapIntentController.Contract.PreSolverCall(&_SwapIntentController.TransactOpts, data)
}

// Swap is a paid mutator transaction binding the contract method 0x83a6992a.
//
// Solidity: function swap((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentController *SwapIntentControllerTransactor) Swap(opts *bind.TransactOpts, swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentController.contract.Transact(opts, "swap", swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x83a6992a.
//
// Solidity: function swap((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentController *SwapIntentControllerSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentController.Contract.Swap(&_SwapIntentController.TransactOpts, swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x83a6992a.
//
// Solidity: function swap((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentController *SwapIntentControllerTransactorSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentController.Contract.Swap(&_SwapIntentController.TransactOpts, swapIntent)
}

// SwapIntentControllerLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the SwapIntentController contract.
type SwapIntentControllerLogIterator struct {
	Event *SwapIntentControllerLog // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLog)
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
		it.Event = new(SwapIntentControllerLog)
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
func (it *SwapIntentControllerLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLog represents a Log event raised by the SwapIntentController contract.
type SwapIntentControllerLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLog(opts *bind.FilterOpts) (*SwapIntentControllerLogIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogIterator{contract: _SwapIntentController.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLog) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLog)
				if err := _SwapIntentController.contract.UnpackLog(event, "log", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLog(log types.Log) (*SwapIntentControllerLog, error) {
	event := new(SwapIntentControllerLog)
	if err := _SwapIntentController.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the SwapIntentController contract.
type SwapIntentControllerLogAddressIterator struct {
	Event *SwapIntentControllerLogAddress // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogAddress)
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
		it.Event = new(SwapIntentControllerLogAddress)
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
func (it *SwapIntentControllerLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogAddress represents a LogAddress event raised by the SwapIntentController contract.
type SwapIntentControllerLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogAddress(opts *bind.FilterOpts) (*SwapIntentControllerLogAddressIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogAddressIterator{contract: _SwapIntentController.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogAddress) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogAddress)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogAddress(log types.Log) (*SwapIntentControllerLogAddress, error) {
	event := new(SwapIntentControllerLogAddress)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the SwapIntentController contract.
type SwapIntentControllerLogArrayIterator struct {
	Event *SwapIntentControllerLogArray // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogArray)
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
		it.Event = new(SwapIntentControllerLogArray)
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
func (it *SwapIntentControllerLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogArray represents a LogArray event raised by the SwapIntentController contract.
type SwapIntentControllerLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogArray(opts *bind.FilterOpts) (*SwapIntentControllerLogArrayIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogArrayIterator{contract: _SwapIntentController.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogArray) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogArray)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogArray(log types.Log) (*SwapIntentControllerLogArray, error) {
	event := new(SwapIntentControllerLogArray)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the SwapIntentController contract.
type SwapIntentControllerLogArray0Iterator struct {
	Event *SwapIntentControllerLogArray0 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogArray0)
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
		it.Event = new(SwapIntentControllerLogArray0)
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
func (it *SwapIntentControllerLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogArray0 represents a LogArray0 event raised by the SwapIntentController contract.
type SwapIntentControllerLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogArray0(opts *bind.FilterOpts) (*SwapIntentControllerLogArray0Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogArray0Iterator{contract: _SwapIntentController.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogArray0) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogArray0)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogArray0(log types.Log) (*SwapIntentControllerLogArray0, error) {
	event := new(SwapIntentControllerLogArray0)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the SwapIntentController contract.
type SwapIntentControllerLogArray1Iterator struct {
	Event *SwapIntentControllerLogArray1 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogArray1)
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
		it.Event = new(SwapIntentControllerLogArray1)
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
func (it *SwapIntentControllerLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogArray1 represents a LogArray1 event raised by the SwapIntentController contract.
type SwapIntentControllerLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogArray1(opts *bind.FilterOpts) (*SwapIntentControllerLogArray1Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogArray1Iterator{contract: _SwapIntentController.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogArray1) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogArray1)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogArray1(log types.Log) (*SwapIntentControllerLogArray1, error) {
	event := new(SwapIntentControllerLogArray1)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the SwapIntentController contract.
type SwapIntentControllerLogBytesIterator struct {
	Event *SwapIntentControllerLogBytes // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogBytes)
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
		it.Event = new(SwapIntentControllerLogBytes)
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
func (it *SwapIntentControllerLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogBytes represents a LogBytes event raised by the SwapIntentController contract.
type SwapIntentControllerLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogBytes(opts *bind.FilterOpts) (*SwapIntentControllerLogBytesIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogBytesIterator{contract: _SwapIntentController.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogBytes) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogBytes)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogBytes(log types.Log) (*SwapIntentControllerLogBytes, error) {
	event := new(SwapIntentControllerLogBytes)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the SwapIntentController contract.
type SwapIntentControllerLogBytes32Iterator struct {
	Event *SwapIntentControllerLogBytes32 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogBytes32)
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
		it.Event = new(SwapIntentControllerLogBytes32)
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
func (it *SwapIntentControllerLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogBytes32 represents a LogBytes32 event raised by the SwapIntentController contract.
type SwapIntentControllerLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*SwapIntentControllerLogBytes32Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogBytes32Iterator{contract: _SwapIntentController.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogBytes32) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogBytes32)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogBytes32(log types.Log) (*SwapIntentControllerLogBytes32, error) {
	event := new(SwapIntentControllerLogBytes32)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the SwapIntentController contract.
type SwapIntentControllerLogIntIterator struct {
	Event *SwapIntentControllerLogInt // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogInt)
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
		it.Event = new(SwapIntentControllerLogInt)
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
func (it *SwapIntentControllerLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogInt represents a LogInt event raised by the SwapIntentController contract.
type SwapIntentControllerLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogInt(opts *bind.FilterOpts) (*SwapIntentControllerLogIntIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogIntIterator{contract: _SwapIntentController.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogInt) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogInt)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogInt(log types.Log) (*SwapIntentControllerLogInt, error) {
	event := new(SwapIntentControllerLogInt)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedAddressIterator struct {
	Event *SwapIntentControllerLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedAddress)
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
		it.Event = new(SwapIntentControllerLogNamedAddress)
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
func (it *SwapIntentControllerLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedAddress represents a LogNamedAddress event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedAddressIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedAddressIterator{contract: _SwapIntentController.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedAddress)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedAddress(log types.Log) (*SwapIntentControllerLogNamedAddress, error) {
	event := new(SwapIntentControllerLogNamedAddress)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArrayIterator struct {
	Event *SwapIntentControllerLogNamedArray // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedArray)
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
		it.Event = new(SwapIntentControllerLogNamedArray)
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
func (it *SwapIntentControllerLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedArray represents a LogNamedArray event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedArrayIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedArrayIterator{contract: _SwapIntentController.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedArray)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedArray(log types.Log) (*SwapIntentControllerLogNamedArray, error) {
	event := new(SwapIntentControllerLogNamedArray)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArray0Iterator struct {
	Event *SwapIntentControllerLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedArray0)
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
		it.Event = new(SwapIntentControllerLogNamedArray0)
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
func (it *SwapIntentControllerLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedArray0 represents a LogNamedArray0 event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedArray0Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedArray0Iterator{contract: _SwapIntentController.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedArray0)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedArray0(log types.Log) (*SwapIntentControllerLogNamedArray0, error) {
	event := new(SwapIntentControllerLogNamedArray0)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArray1Iterator struct {
	Event *SwapIntentControllerLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedArray1)
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
		it.Event = new(SwapIntentControllerLogNamedArray1)
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
func (it *SwapIntentControllerLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedArray1 represents a LogNamedArray1 event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedArray1Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedArray1Iterator{contract: _SwapIntentController.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedArray1)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedArray1(log types.Log) (*SwapIntentControllerLogNamedArray1, error) {
	event := new(SwapIntentControllerLogNamedArray1)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedBytesIterator struct {
	Event *SwapIntentControllerLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedBytes)
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
		it.Event = new(SwapIntentControllerLogNamedBytes)
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
func (it *SwapIntentControllerLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedBytes represents a LogNamedBytes event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedBytesIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedBytesIterator{contract: _SwapIntentController.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedBytes)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedBytes(log types.Log) (*SwapIntentControllerLogNamedBytes, error) {
	event := new(SwapIntentControllerLogNamedBytes)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedBytes32Iterator struct {
	Event *SwapIntentControllerLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedBytes32)
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
		it.Event = new(SwapIntentControllerLogNamedBytes32)
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
func (it *SwapIntentControllerLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedBytes32 represents a LogNamedBytes32 event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedBytes32Iterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedBytes32Iterator{contract: _SwapIntentController.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedBytes32)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedBytes32(log types.Log) (*SwapIntentControllerLogNamedBytes32, error) {
	event := new(SwapIntentControllerLogNamedBytes32)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedDecimalIntIterator struct {
	Event *SwapIntentControllerLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedDecimalInt)
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
		it.Event = new(SwapIntentControllerLogNamedDecimalInt)
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
func (it *SwapIntentControllerLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedDecimalIntIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedDecimalIntIterator{contract: _SwapIntentController.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedDecimalInt)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedDecimalInt(log types.Log) (*SwapIntentControllerLogNamedDecimalInt, error) {
	event := new(SwapIntentControllerLogNamedDecimalInt)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedDecimalUintIterator struct {
	Event *SwapIntentControllerLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedDecimalUint)
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
		it.Event = new(SwapIntentControllerLogNamedDecimalUint)
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
func (it *SwapIntentControllerLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedDecimalUintIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedDecimalUintIterator{contract: _SwapIntentController.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedDecimalUint)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedDecimalUint(log types.Log) (*SwapIntentControllerLogNamedDecimalUint, error) {
	event := new(SwapIntentControllerLogNamedDecimalUint)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedIntIterator struct {
	Event *SwapIntentControllerLogNamedInt // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedInt)
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
		it.Event = new(SwapIntentControllerLogNamedInt)
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
func (it *SwapIntentControllerLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedInt represents a LogNamedInt event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedIntIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedIntIterator{contract: _SwapIntentController.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedInt)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedInt(log types.Log) (*SwapIntentControllerLogNamedInt, error) {
	event := new(SwapIntentControllerLogNamedInt)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedStringIterator struct {
	Event *SwapIntentControllerLogNamedString // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedString)
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
		it.Event = new(SwapIntentControllerLogNamedString)
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
func (it *SwapIntentControllerLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedString represents a LogNamedString event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedStringIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedStringIterator{contract: _SwapIntentController.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedString) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedString)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedString(log types.Log) (*SwapIntentControllerLogNamedString, error) {
	event := new(SwapIntentControllerLogNamedString)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedUintIterator struct {
	Event *SwapIntentControllerLogNamedUint // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogNamedUint)
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
		it.Event = new(SwapIntentControllerLogNamedUint)
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
func (it *SwapIntentControllerLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogNamedUint represents a LogNamedUint event raised by the SwapIntentController contract.
type SwapIntentControllerLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*SwapIntentControllerLogNamedUintIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogNamedUintIterator{contract: _SwapIntentController.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogNamedUint)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogNamedUint(log types.Log) (*SwapIntentControllerLogNamedUint, error) {
	event := new(SwapIntentControllerLogNamedUint)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the SwapIntentController contract.
type SwapIntentControllerLogStringIterator struct {
	Event *SwapIntentControllerLogString // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogString)
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
		it.Event = new(SwapIntentControllerLogString)
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
func (it *SwapIntentControllerLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogString represents a LogString event raised by the SwapIntentController contract.
type SwapIntentControllerLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogString(opts *bind.FilterOpts) (*SwapIntentControllerLogStringIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogStringIterator{contract: _SwapIntentController.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogString) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogString)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogString(log types.Log) (*SwapIntentControllerLogString, error) {
	event := new(SwapIntentControllerLogString)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the SwapIntentController contract.
type SwapIntentControllerLogUintIterator struct {
	Event *SwapIntentControllerLogUint // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogUint)
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
		it.Event = new(SwapIntentControllerLogUint)
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
func (it *SwapIntentControllerLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogUint represents a LogUint event raised by the SwapIntentController contract.
type SwapIntentControllerLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogUint(opts *bind.FilterOpts) (*SwapIntentControllerLogUintIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogUintIterator{contract: _SwapIntentController.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogUint) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogUint)
				if err := _SwapIntentController.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogUint(log types.Log) (*SwapIntentControllerLogUint, error) {
	event := new(SwapIntentControllerLogUint)
	if err := _SwapIntentController.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentControllerLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the SwapIntentController contract.
type SwapIntentControllerLogsIterator struct {
	Event *SwapIntentControllerLogs // Event containing the contract specifics and raw log

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
func (it *SwapIntentControllerLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentControllerLogs)
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
		it.Event = new(SwapIntentControllerLogs)
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
func (it *SwapIntentControllerLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentControllerLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentControllerLogs represents a Logs event raised by the SwapIntentController contract.
type SwapIntentControllerLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) FilterLogs(opts *bind.FilterOpts) (*SwapIntentControllerLogsIterator, error) {

	logs, sub, err := _SwapIntentController.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &SwapIntentControllerLogsIterator{contract: _SwapIntentController.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SwapIntentController *SwapIntentControllerFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *SwapIntentControllerLogs) (event.Subscription, error) {

	logs, sub, err := _SwapIntentController.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentControllerLogs)
				if err := _SwapIntentController.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_SwapIntentController *SwapIntentControllerFilterer) ParseLogs(log types.Log) (*SwapIntentControllerLogs, error) {
	event := new(SwapIntentControllerLogs)
	if err := _SwapIntentController.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
