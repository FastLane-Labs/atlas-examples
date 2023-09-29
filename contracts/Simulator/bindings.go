// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Simulator

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

// SimulatorMetaData contains all meta data concerning the Simulator contract.
var SimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlteredControlHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredUserHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashChainBroken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentUnfulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolverHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuctionWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverBidUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverEVMError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverFailedCallback\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMsgValueUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverOperationReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerificationSimFail\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structBidData[]\",\"name\":\"winningBids\",\"type\":\"tuple[]\"}],\"name\":\"MEVPaymentFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"environment\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewExecutionEnvironment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"SolverTxResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasRefunded\",\"type\":\"uint256\"}],\"name\":\"UserTxResult\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppApproval\",\"name\":\"approval\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"metacallSimulation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"name\":\"setAtlas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppApproval\",\"name\":\"approval\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"simSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppApproval\",\"name\":\"approval\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"simSolverCalls\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"uCall\",\"type\":\"tuple\"}],\"name\":\"simUserOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simUserOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SimulatorMetaData.ABI instead.
var SimulatorABI = SimulatorMetaData.ABI

// Simulator is an auto generated Go binding around an Ethereum contract.
type Simulator struct {
	SimulatorCaller     // Read-only binding to the contract
	SimulatorTransactor // Write-only binding to the contract
	SimulatorFilterer   // Log filterer for contract events
}

// SimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimulatorSession struct {
	Contract     *Simulator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimulatorCallerSession struct {
	Contract *SimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimulatorTransactorSession struct {
	Contract     *SimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimulatorRaw struct {
	Contract *Simulator // Generic contract binding to access the raw methods on
}

// SimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimulatorCallerRaw struct {
	Contract *SimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// SimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimulatorTransactorRaw struct {
	Contract *SimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimulator creates a new instance of Simulator, bound to a specific deployed contract.
func NewSimulator(address common.Address, backend bind.ContractBackend) (*Simulator, error) {
	contract, err := bindSimulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simulator{SimulatorCaller: SimulatorCaller{contract: contract}, SimulatorTransactor: SimulatorTransactor{contract: contract}, SimulatorFilterer: SimulatorFilterer{contract: contract}}, nil
}

// NewSimulatorCaller creates a new read-only instance of Simulator, bound to a specific deployed contract.
func NewSimulatorCaller(address common.Address, caller bind.ContractCaller) (*SimulatorCaller, error) {
	contract, err := bindSimulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimulatorCaller{contract: contract}, nil
}

// NewSimulatorTransactor creates a new write-only instance of Simulator, bound to a specific deployed contract.
func NewSimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SimulatorTransactor, error) {
	contract, err := bindSimulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimulatorTransactor{contract: contract}, nil
}

// NewSimulatorFilterer creates a new log filterer instance of Simulator, bound to a specific deployed contract.
func NewSimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SimulatorFilterer, error) {
	contract, err := bindSimulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimulatorFilterer{contract: contract}, nil
}

// bindSimulator binds a generic wrapper to an already deployed contract.
func bindSimulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimulatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simulator *SimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simulator.Contract.SimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simulator *SimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.Contract.SimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simulator *SimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simulator.Contract.SimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simulator *SimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simulator *SimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simulator *SimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simulator.Contract.contract.Transact(opts, method, params...)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simulator.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorSession) Atlas() (common.Address, error) {
	return _Simulator.Contract.Atlas(&_Simulator.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorCallerSession) Atlas() (common.Address, error) {
	return _Simulator.Contract.Atlas(&_Simulator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simulator.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorSession) Deployer() (common.Address, error) {
	return _Simulator.Contract.Deployer(&_Simulator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorCallerSession) Deployer() (common.Address, error) {
	return _Simulator.Contract.Deployer(&_Simulator.CallOpts)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x4f906ad0.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns()
func (_Simulator *SimulatorTransactor) MetacallSimulation(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "metacallSimulation", dConfig, userOp, solverOps, dAppOp)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x4f906ad0.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns()
func (_Simulator *SimulatorSession) MetacallSimulation(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.MetacallSimulation(&_Simulator.TransactOpts, dConfig, userOp, solverOps, dAppOp)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x4f906ad0.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns()
func (_Simulator *SimulatorTransactorSession) MetacallSimulation(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.MetacallSimulation(&_Simulator.TransactOpts, dConfig, userOp, solverOps, dAppOp)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorTransactor) SetAtlas(opts *bind.TransactOpts, _atlas common.Address) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "setAtlas", _atlas)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorSession) SetAtlas(_atlas common.Address) (*types.Transaction, error) {
	return _Simulator.Contract.SetAtlas(&_Simulator.TransactOpts, _atlas)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorTransactorSession) SetAtlas(_atlas common.Address) (*types.Transaction, error) {
	return _Simulator.Contract.SetAtlas(&_Simulator.TransactOpts, _atlas)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0x63592f8d.
//
// Solidity: function simSolverCall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimSolverCall(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCall", dConfig, userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0x63592f8d.
//
// Solidity: function simSolverCall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimSolverCall(dConfig DAppConfig, userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, dConfig, userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0x63592f8d.
//
// Solidity: function simSolverCall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[]) solverOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimSolverCall(dConfig DAppConfig, userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, dConfig, userOp, solverOp, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0xe9ee7fc8.
//
// Solidity: function simSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimSolverCalls(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCalls", dConfig, userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0xe9ee7fc8.
//
// Solidity: function simSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimSolverCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, dConfig, userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0xe9ee7fc8.
//
// Solidity: function simSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32),bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimSolverCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, dConfig, userOp, solverOps, dAppOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x52355edf.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimUserOperation(opts *bind.TransactOpts, uCall UserCall) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simUserOperation", uCall)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x52355edf.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) payable returns(bool success)
func (_Simulator *SimulatorSession) SimUserOperation(uCall UserCall) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, uCall)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x52355edf.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimUserOperation(uCall UserCall) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, uCall)
}

// SimUserOperation0 is a paid mutator transaction binding the contract method 0x7303b53c.
//
// Solidity: function simUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimUserOperation0(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simUserOperation0", userOp)
}

// SimUserOperation0 is a paid mutator transaction binding the contract method 0x7303b53c.
//
// Solidity: function simUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimUserOperation0(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation0(&_Simulator.TransactOpts, userOp)
}

// SimUserOperation0 is a paid mutator transaction binding the contract method 0x7303b53c.
//
// Solidity: function simUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimUserOperation0(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation0(&_Simulator.TransactOpts, userOp)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Simulator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Simulator.Contract.Fallback(&_Simulator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Simulator.Contract.Fallback(&_Simulator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorSession) Receive() (*types.Transaction, error) {
	return _Simulator.Contract.Receive(&_Simulator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Simulator.Contract.Receive(&_Simulator.TransactOpts)
}

// SimulatorMEVPaymentFailureIterator is returned from FilterMEVPaymentFailure and is used to iterate over the raw logs and unpacked data for MEVPaymentFailure events raised by the Simulator contract.
type SimulatorMEVPaymentFailureIterator struct {
	Event *SimulatorMEVPaymentFailure // Event containing the contract specifics and raw log

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
func (it *SimulatorMEVPaymentFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimulatorMEVPaymentFailure)
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
		it.Event = new(SimulatorMEVPaymentFailure)
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
func (it *SimulatorMEVPaymentFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimulatorMEVPaymentFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimulatorMEVPaymentFailure represents a MEVPaymentFailure event raised by the Simulator contract.
type SimulatorMEVPaymentFailure struct {
	Controller  common.Address
	CallConfig  uint32
	WinningBids []BidData
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMEVPaymentFailure is a free log retrieval operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
func (_Simulator *SimulatorFilterer) FilterMEVPaymentFailure(opts *bind.FilterOpts, controller []common.Address) (*SimulatorMEVPaymentFailureIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Simulator.contract.FilterLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return &SimulatorMEVPaymentFailureIterator{contract: _Simulator.contract, event: "MEVPaymentFailure", logs: logs, sub: sub}, nil
}

// WatchMEVPaymentFailure is a free log subscription operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
func (_Simulator *SimulatorFilterer) WatchMEVPaymentFailure(opts *bind.WatchOpts, sink chan<- *SimulatorMEVPaymentFailure, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Simulator.contract.WatchLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimulatorMEVPaymentFailure)
				if err := _Simulator.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
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

// ParseMEVPaymentFailure is a log parse operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
func (_Simulator *SimulatorFilterer) ParseMEVPaymentFailure(log types.Log) (*SimulatorMEVPaymentFailure, error) {
	event := new(SimulatorMEVPaymentFailure)
	if err := _Simulator.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimulatorNewExecutionEnvironmentIterator is returned from FilterNewExecutionEnvironment and is used to iterate over the raw logs and unpacked data for NewExecutionEnvironment events raised by the Simulator contract.
type SimulatorNewExecutionEnvironmentIterator struct {
	Event *SimulatorNewExecutionEnvironment // Event containing the contract specifics and raw log

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
func (it *SimulatorNewExecutionEnvironmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimulatorNewExecutionEnvironment)
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
		it.Event = new(SimulatorNewExecutionEnvironment)
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
func (it *SimulatorNewExecutionEnvironmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimulatorNewExecutionEnvironmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimulatorNewExecutionEnvironment represents a NewExecutionEnvironment event raised by the Simulator contract.
type SimulatorNewExecutionEnvironment struct {
	Environment common.Address
	User        common.Address
	Controller  common.Address
	CallConfig  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionEnvironment is a free log retrieval operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_Simulator *SimulatorFilterer) FilterNewExecutionEnvironment(opts *bind.FilterOpts, environment []common.Address, user []common.Address, controller []common.Address) (*SimulatorNewExecutionEnvironmentIterator, error) {

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

	logs, sub, err := _Simulator.contract.FilterLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &SimulatorNewExecutionEnvironmentIterator{contract: _Simulator.contract, event: "NewExecutionEnvironment", logs: logs, sub: sub}, nil
}

// WatchNewExecutionEnvironment is a free log subscription operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_Simulator *SimulatorFilterer) WatchNewExecutionEnvironment(opts *bind.WatchOpts, sink chan<- *SimulatorNewExecutionEnvironment, environment []common.Address, user []common.Address, controller []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Simulator.contract.WatchLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimulatorNewExecutionEnvironment)
				if err := _Simulator.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
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
func (_Simulator *SimulatorFilterer) ParseNewExecutionEnvironment(log types.Log) (*SimulatorNewExecutionEnvironment, error) {
	event := new(SimulatorNewExecutionEnvironment)
	if err := _Simulator.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimulatorSolverTxResultIterator is returned from FilterSolverTxResult and is used to iterate over the raw logs and unpacked data for SolverTxResult events raised by the Simulator contract.
type SimulatorSolverTxResultIterator struct {
	Event *SimulatorSolverTxResult // Event containing the contract specifics and raw log

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
func (it *SimulatorSolverTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimulatorSolverTxResult)
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
		it.Event = new(SimulatorSolverTxResult)
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
func (it *SimulatorSolverTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimulatorSolverTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimulatorSolverTxResult represents a SolverTxResult event raised by the Simulator contract.
type SimulatorSolverTxResult struct {
	SolverTo   common.Address
	SolverFrom common.Address
	Executed   bool
	Success    bool
	Nonce      *big.Int
	Result     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSolverTxResult is a free log retrieval operation binding the contract event 0xeea2753019c05c54d8dc4ff90699ee5034f9ac243b7c2f9e3a0db08f6597a838.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 nonce, uint256 result)
func (_Simulator *SimulatorFilterer) FilterSolverTxResult(opts *bind.FilterOpts, solverTo []common.Address, solverFrom []common.Address) (*SimulatorSolverTxResultIterator, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Simulator.contract.FilterLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return &SimulatorSolverTxResultIterator{contract: _Simulator.contract, event: "SolverTxResult", logs: logs, sub: sub}, nil
}

// WatchSolverTxResult is a free log subscription operation binding the contract event 0xeea2753019c05c54d8dc4ff90699ee5034f9ac243b7c2f9e3a0db08f6597a838.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 nonce, uint256 result)
func (_Simulator *SimulatorFilterer) WatchSolverTxResult(opts *bind.WatchOpts, sink chan<- *SimulatorSolverTxResult, solverTo []common.Address, solverFrom []common.Address) (event.Subscription, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Simulator.contract.WatchLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimulatorSolverTxResult)
				if err := _Simulator.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
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

// ParseSolverTxResult is a log parse operation binding the contract event 0xeea2753019c05c54d8dc4ff90699ee5034f9ac243b7c2f9e3a0db08f6597a838.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 nonce, uint256 result)
func (_Simulator *SimulatorFilterer) ParseSolverTxResult(log types.Log) (*SimulatorSolverTxResult, error) {
	event := new(SimulatorSolverTxResult)
	if err := _Simulator.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimulatorUserTxResultIterator is returned from FilterUserTxResult and is used to iterate over the raw logs and unpacked data for UserTxResult events raised by the Simulator contract.
type SimulatorUserTxResultIterator struct {
	Event *SimulatorUserTxResult // Event containing the contract specifics and raw log

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
func (it *SimulatorUserTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimulatorUserTxResult)
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
		it.Event = new(SimulatorUserTxResult)
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
func (it *SimulatorUserTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimulatorUserTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimulatorUserTxResult represents a UserTxResult event raised by the Simulator contract.
type SimulatorUserTxResult struct {
	User          common.Address
	ValueReturned *big.Int
	GasRefunded   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserTxResult is a free log retrieval operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Simulator *SimulatorFilterer) FilterUserTxResult(opts *bind.FilterOpts, user []common.Address) (*SimulatorUserTxResultIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Simulator.contract.FilterLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return &SimulatorUserTxResultIterator{contract: _Simulator.contract, event: "UserTxResult", logs: logs, sub: sub}, nil
}

// WatchUserTxResult is a free log subscription operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Simulator *SimulatorFilterer) WatchUserTxResult(opts *bind.WatchOpts, sink chan<- *SimulatorUserTxResult, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Simulator.contract.WatchLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimulatorUserTxResult)
				if err := _Simulator.contract.UnpackLog(event, "UserTxResult", log); err != nil {
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

// ParseUserTxResult is a log parse operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Simulator *SimulatorFilterer) ParseUserTxResult(log types.Log) (*SimulatorUserTxResult, error) {
	event := new(SimulatorUserTxResult)
	if err := _Simulator.contract.UnpackLog(event, "UserTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
