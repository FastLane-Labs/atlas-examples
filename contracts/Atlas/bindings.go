// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Atlas

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

// AtlasMetaData contains all meta data concerning the Atlas contract.
var AtlasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_escrowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verification\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasAccLib\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safetyLocksLib\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_simulator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredControlHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredUserHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DAppNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnvironmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasAccountingLibError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashChainBroken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentUnfulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDAppControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEnvironment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLockState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolverHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerBalancing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"MissingFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuctionWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnfilledRequests\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevertToReuse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafetyLocksLibError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatoryActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverBidUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverEVMError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverFailedCallback\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMsgValueUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMustReconcile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverOperationReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UncoveredResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpSimFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ValidCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerificationSimFail\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"MEVPaymentFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"environment\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewExecutionEnvironment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"SolverTxResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasRefunded\",\"type\":\"uint256\"}],\"name\":\"UserTxResult\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ESCROW_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_ACC_LIB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAFETY_LOCKS_LIB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIMULATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFICATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountLastActiveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"recipient\",\"type\":\"uint8\"}],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"donor\",\"type\":\"uint8\"},{\"internalType\":\"enumParty\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"contributeGasTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"donor\",\"type\":\"uint8\"},{\"internalType\":\"enumParty\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"contributeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dAppControl\",\"type\":\"address\"}],\"name\":\"createExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"party\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"auctionWon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"accruedGasRebate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winningSearcherIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"party\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"partyAddress\",\"type\":\"address\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ledgers\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"balance\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"contributed\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"requested\",\"type\":\"int64\"},{\"internalType\":\"enumLedgerStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumParty\",\"name\":\"proxy\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"activeEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"activeParties\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"startingBalance\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"metacall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"auctionWon\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nextAccountNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"environment\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"searcherFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxApprovedGasSpend\",\"type\":\"uint256\"}],\"name\":\"reconcile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"donor\",\"type\":\"uint8\"},{\"internalType\":\"enumParty\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"requestFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumParty\",\"name\":\"donor\",\"type\":\"uint8\"},{\"internalType\":\"enumParty\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"requestGasFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"transferDAppERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"transferUserERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validateBalances\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AtlasABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlasMetaData.ABI instead.
var AtlasABI = AtlasMetaData.ABI

// Atlas is an auto generated Go binding around an Ethereum contract.
type Atlas struct {
	AtlasCaller     // Read-only binding to the contract
	AtlasTransactor // Write-only binding to the contract
	AtlasFilterer   // Log filterer for contract events
}

// AtlasCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlasSession struct {
	Contract     *Atlas            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlasCallerSession struct {
	Contract *AtlasCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtlasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlasTransactorSession struct {
	Contract     *AtlasTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlasRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlasRaw struct {
	Contract *Atlas // Generic contract binding to access the raw methods on
}

// AtlasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlasCallerRaw struct {
	Contract *AtlasCaller // Generic read-only contract binding to access the raw methods on
}

// AtlasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlasTransactorRaw struct {
	Contract *AtlasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlas creates a new instance of Atlas, bound to a specific deployed contract.
func NewAtlas(address common.Address, backend bind.ContractBackend) (*Atlas, error) {
	contract, err := bindAtlas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atlas{AtlasCaller: AtlasCaller{contract: contract}, AtlasTransactor: AtlasTransactor{contract: contract}, AtlasFilterer: AtlasFilterer{contract: contract}}, nil
}

// NewAtlasCaller creates a new read-only instance of Atlas, bound to a specific deployed contract.
func NewAtlasCaller(address common.Address, caller bind.ContractCaller) (*AtlasCaller, error) {
	contract, err := bindAtlas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasCaller{contract: contract}, nil
}

// NewAtlasTransactor creates a new write-only instance of Atlas, bound to a specific deployed contract.
func NewAtlasTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlasTransactor, error) {
	contract, err := bindAtlas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasTransactor{contract: contract}, nil
}

// NewAtlasFilterer creates a new log filterer instance of Atlas, bound to a specific deployed contract.
func NewAtlasFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlasFilterer, error) {
	contract, err := bindAtlas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlasFilterer{contract: contract}, nil
}

// bindAtlas binds a generic wrapper to an already deployed contract.
func bindAtlas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtlasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atlas *AtlasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atlas.Contract.AtlasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atlas *AtlasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.Contract.AtlasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atlas *AtlasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atlas.Contract.AtlasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atlas *AtlasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atlas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atlas *AtlasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atlas *AtlasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atlas.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atlas.Contract.DOMAINSEPARATOR(&_Atlas.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atlas.Contract.DOMAINSEPARATOR(&_Atlas.CallOpts)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasCaller) ESCROWDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "ESCROW_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasSession) ESCROWDURATION() (*big.Int, error) {
	return _Atlas.Contract.ESCROWDURATION(&_Atlas.CallOpts)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasCallerSession) ESCROWDURATION() (*big.Int, error) {
	return _Atlas.Contract.ESCROWDURATION(&_Atlas.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Atlas *AtlasCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Atlas *AtlasSession) FACTORY() (common.Address, error) {
	return _Atlas.Contract.FACTORY(&_Atlas.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Atlas *AtlasCallerSession) FACTORY() (common.Address, error) {
	return _Atlas.Contract.FACTORY(&_Atlas.CallOpts)
}

// GASACCLIB is a free data retrieval call binding the contract method 0x5986ce00.
//
// Solidity: function GAS_ACC_LIB() view returns(address)
func (_Atlas *AtlasCaller) GASACCLIB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "GAS_ACC_LIB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GASACCLIB is a free data retrieval call binding the contract method 0x5986ce00.
//
// Solidity: function GAS_ACC_LIB() view returns(address)
func (_Atlas *AtlasSession) GASACCLIB() (common.Address, error) {
	return _Atlas.Contract.GASACCLIB(&_Atlas.CallOpts)
}

// GASACCLIB is a free data retrieval call binding the contract method 0x5986ce00.
//
// Solidity: function GAS_ACC_LIB() view returns(address)
func (_Atlas *AtlasCallerSession) GASACCLIB() (common.Address, error) {
	return _Atlas.Contract.GASACCLIB(&_Atlas.CallOpts)
}

// SAFETYLOCKSLIB is a free data retrieval call binding the contract method 0xa89f0fc1.
//
// Solidity: function SAFETY_LOCKS_LIB() view returns(address)
func (_Atlas *AtlasCaller) SAFETYLOCKSLIB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SAFETY_LOCKS_LIB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SAFETYLOCKSLIB is a free data retrieval call binding the contract method 0xa89f0fc1.
//
// Solidity: function SAFETY_LOCKS_LIB() view returns(address)
func (_Atlas *AtlasSession) SAFETYLOCKSLIB() (common.Address, error) {
	return _Atlas.Contract.SAFETYLOCKSLIB(&_Atlas.CallOpts)
}

// SAFETYLOCKSLIB is a free data retrieval call binding the contract method 0xa89f0fc1.
//
// Solidity: function SAFETY_LOCKS_LIB() view returns(address)
func (_Atlas *AtlasCallerSession) SAFETYLOCKSLIB() (common.Address, error) {
	return _Atlas.Contract.SAFETYLOCKSLIB(&_Atlas.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasCaller) SIMULATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SIMULATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasSession) SIMULATOR() (common.Address, error) {
	return _Atlas.Contract.SIMULATOR(&_Atlas.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasCallerSession) SIMULATOR() (common.Address, error) {
	return _Atlas.Contract.SIMULATOR(&_Atlas.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasCaller) VERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasSession) VERIFICATION() (common.Address, error) {
	return _Atlas.Contract.VERIFICATION(&_Atlas.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasCallerSession) VERIFICATION() (common.Address, error) {
	return _Atlas.Contract.VERIFICATION(&_Atlas.CallOpts)
}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasCaller) AccountLastActiveBlock(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "accountLastActiveBlock", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasSession) AccountLastActiveBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.AccountLastActiveBlock(&_Atlas.CallOpts, account)
}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) AccountLastActiveBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.AccountLastActiveBlock(&_Atlas.CallOpts, account)
}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasCaller) ActiveEnvironment(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "activeEnvironment")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasSession) ActiveEnvironment() (common.Address, error) {
	return _Atlas.Contract.ActiveEnvironment(&_Atlas.CallOpts)
}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasCallerSession) ActiveEnvironment() (common.Address, error) {
	return _Atlas.Contract.ActiveEnvironment(&_Atlas.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Allowance(&_Atlas.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Allowance(&_Atlas.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOf(&_Atlas.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOf(&_Atlas.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasSession) Decimals() (uint8, error) {
	return _Atlas.Contract.Decimals(&_Atlas.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasCallerSession) Decimals() (uint8, error) {
	return _Atlas.Contract.Decimals(&_Atlas.CallOpts)
}

// Ledgers is a free data retrieval call binding the contract method 0xe9789190.
//
// Solidity: function ledgers(uint256 ) view returns(int64 balance, int64 contributed, int64 requested, uint8 status, uint8 proxy)
func (_Atlas *AtlasCaller) Ledgers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Balance     int64
	Contributed int64
	Requested   int64
	Status      uint8
	Proxy       uint8
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "ledgers", arg0)

	outstruct := new(struct {
		Balance     int64
		Contributed int64
		Requested   int64
		Status      uint8
		Proxy       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.Contributed = *abi.ConvertType(out[1], new(int64)).(*int64)
	outstruct.Requested = *abi.ConvertType(out[2], new(int64)).(*int64)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Proxy = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// Ledgers is a free data retrieval call binding the contract method 0xe9789190.
//
// Solidity: function ledgers(uint256 ) view returns(int64 balance, int64 contributed, int64 requested, uint8 status, uint8 proxy)
func (_Atlas *AtlasSession) Ledgers(arg0 *big.Int) (struct {
	Balance     int64
	Contributed int64
	Requested   int64
	Status      uint8
	Proxy       uint8
}, error) {
	return _Atlas.Contract.Ledgers(&_Atlas.CallOpts, arg0)
}

// Ledgers is a free data retrieval call binding the contract method 0xe9789190.
//
// Solidity: function ledgers(uint256 ) view returns(int64 balance, int64 contributed, int64 requested, uint8 status, uint8 proxy)
func (_Atlas *AtlasCallerSession) Ledgers(arg0 *big.Int) (struct {
	Balance     int64
	Contributed int64
	Requested   int64
	Status      uint8
	Proxy       uint8
}, error) {
	return _Atlas.Contract.Ledgers(&_Atlas.CallOpts, arg0)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint16 activeParties, uint64 startingBalance)
func (_Atlas *AtlasCaller) Lock(opts *bind.CallOpts) (struct {
	ActiveEnvironment common.Address
	ActiveParties     uint16
	StartingBalance   uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "lock")

	outstruct := new(struct {
		ActiveEnvironment common.Address
		ActiveParties     uint16
		StartingBalance   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActiveEnvironment = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ActiveParties = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.StartingBalance = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint16 activeParties, uint64 startingBalance)
func (_Atlas *AtlasSession) Lock() (struct {
	ActiveEnvironment common.Address
	ActiveParties     uint16
	StartingBalance   uint64
}, error) {
	return _Atlas.Contract.Lock(&_Atlas.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint16 activeParties, uint64 startingBalance)
func (_Atlas *AtlasCallerSession) Lock() (struct {
	ActiveEnvironment common.Address
	ActiveParties     uint16
	StartingBalance   uint64
}, error) {
	return _Atlas.Contract.Lock(&_Atlas.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasSession) Name() (string, error) {
	return _Atlas.Contract.Name(&_Atlas.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasCallerSession) Name() (string, error) {
	return _Atlas.Contract.Name(&_Atlas.CallOpts)
}

// NextAccountNonce is a free data retrieval call binding the contract method 0x061d3ac6.
//
// Solidity: function nextAccountNonce(address account) view returns(uint256)
func (_Atlas *AtlasCaller) NextAccountNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nextAccountNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAccountNonce is a free data retrieval call binding the contract method 0x061d3ac6.
//
// Solidity: function nextAccountNonce(address account) view returns(uint256)
func (_Atlas *AtlasSession) NextAccountNonce(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextAccountNonce(&_Atlas.CallOpts, account)
}

// NextAccountNonce is a free data retrieval call binding the contract method 0x061d3ac6.
//
// Solidity: function nextAccountNonce(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) NextAccountNonce(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextAccountNonce(&_Atlas.CallOpts, account)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Nonces(&_Atlas.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Nonces(&_Atlas.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasSession) Symbol() (string, error) {
	return _Atlas.Contract.Symbol(&_Atlas.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasCallerSession) Symbol() (string, error) {
	return _Atlas.Contract.Symbol(&_Atlas.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasSession) TotalSupply() (*big.Int, error) {
	return _Atlas.Contract.TotalSupply(&_Atlas.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasCallerSession) TotalSupply() (*big.Int, error) {
	return _Atlas.Contract.TotalSupply(&_Atlas.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Approve(&_Atlas.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Approve(&_Atlas.TransactOpts, spender, amount)
}

// Contribute is a paid mutator transaction binding the contract method 0x64c2c3ee.
//
// Solidity: function contribute(uint8 recipient) payable returns()
func (_Atlas *AtlasTransactor) Contribute(opts *bind.TransactOpts, recipient uint8) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "contribute", recipient)
}

// Contribute is a paid mutator transaction binding the contract method 0x64c2c3ee.
//
// Solidity: function contribute(uint8 recipient) payable returns()
func (_Atlas *AtlasSession) Contribute(recipient uint8) (*types.Transaction, error) {
	return _Atlas.Contract.Contribute(&_Atlas.TransactOpts, recipient)
}

// Contribute is a paid mutator transaction binding the contract method 0x64c2c3ee.
//
// Solidity: function contribute(uint8 recipient) payable returns()
func (_Atlas *AtlasTransactorSession) Contribute(recipient uint8) (*types.Transaction, error) {
	return _Atlas.Contract.Contribute(&_Atlas.TransactOpts, recipient)
}

// ContributeGasTo is a paid mutator transaction binding the contract method 0x6f22cf99.
//
// Solidity: function contributeGasTo(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) ContributeGasTo(opts *bind.TransactOpts, donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "contributeGasTo", donor, recipient, amt, lockState)
}

// ContributeGasTo is a paid mutator transaction binding the contract method 0x6f22cf99.
//
// Solidity: function contributeGasTo(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasSession) ContributeGasTo(donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.ContributeGasTo(&_Atlas.TransactOpts, donor, recipient, amt, lockState)
}

// ContributeGasTo is a paid mutator transaction binding the contract method 0x6f22cf99.
//
// Solidity: function contributeGasTo(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) ContributeGasTo(donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.ContributeGasTo(&_Atlas.TransactOpts, donor, recipient, amt, lockState)
}

// ContributeTo is a paid mutator transaction binding the contract method 0x7a031c6c.
//
// Solidity: function contributeTo(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasTransactor) ContributeTo(opts *bind.TransactOpts, donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "contributeTo", donor, recipient, amt)
}

// ContributeTo is a paid mutator transaction binding the contract method 0x7a031c6c.
//
// Solidity: function contributeTo(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasSession) ContributeTo(donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.ContributeTo(&_Atlas.TransactOpts, donor, recipient, amt)
}

// ContributeTo is a paid mutator transaction binding the contract method 0x7a031c6c.
//
// Solidity: function contributeTo(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasTransactorSession) ContributeTo(donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.ContributeTo(&_Atlas.TransactOpts, donor, recipient, amt)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasTransactor) CreateExecutionEnvironment(opts *bind.TransactOpts, dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "createExecutionEnvironment", dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasSession) CreateExecutionEnvironment(dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasTransactorSession) CreateExecutionEnvironment(dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dAppControl)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasSession) Deposit() (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasTransactorSession) Deposit() (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xebe41b73.
//
// Solidity: function deposit(uint8 party) payable returns()
func (_Atlas *AtlasTransactor) Deposit0(opts *bind.TransactOpts, party uint8) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "deposit0", party)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xebe41b73.
//
// Solidity: function deposit(uint8 party) payable returns()
func (_Atlas *AtlasSession) Deposit0(party uint8) (*types.Transaction, error) {
	return _Atlas.Contract.Deposit0(&_Atlas.TransactOpts, party)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xebe41b73.
//
// Solidity: function deposit(uint8 party) payable returns()
func (_Atlas *AtlasTransactorSession) Deposit0(party uint8) (*types.Transaction, error) {
	return _Atlas.Contract.Deposit0(&_Atlas.TransactOpts, party)
}

// Execute is a paid mutator transaction binding the contract method 0x06df7940.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler) payable returns(bool auctionWon, uint256 accruedGasRebate, uint256 winningSearcherIndex)
func (_Atlas *AtlasTransactor) Execute(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "execute", dConfig, userOp, solverOps, executionEnvironment, bundler)
}

// Execute is a paid mutator transaction binding the contract method 0x06df7940.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler) payable returns(bool auctionWon, uint256 accruedGasRebate, uint256 winningSearcherIndex)
func (_Atlas *AtlasSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler)
}

// Execute is a paid mutator transaction binding the contract method 0x06df7940.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler) payable returns(bool auctionWon, uint256 accruedGasRebate, uint256 winningSearcherIndex)
func (_Atlas *AtlasTransactorSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler)
}

// Finalize is a paid mutator transaction binding the contract method 0xcd960226.
//
// Solidity: function finalize(uint8 party, address partyAddress) returns()
func (_Atlas *AtlasTransactor) Finalize(opts *bind.TransactOpts, party uint8, partyAddress common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "finalize", party, partyAddress)
}

// Finalize is a paid mutator transaction binding the contract method 0xcd960226.
//
// Solidity: function finalize(uint8 party, address partyAddress) returns()
func (_Atlas *AtlasSession) Finalize(party uint8, partyAddress common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Finalize(&_Atlas.TransactOpts, party, partyAddress)
}

// Finalize is a paid mutator transaction binding the contract method 0xcd960226.
//
// Solidity: function finalize(uint8 party, address partyAddress) returns()
func (_Atlas *AtlasTransactorSession) Finalize(party uint8, partyAddress common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Finalize(&_Atlas.TransactOpts, party, partyAddress)
}

// Metacall is a paid mutator transaction binding the contract method 0xceb3e79b.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactor) Metacall(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "metacall", userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0xceb3e79b.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0xceb3e79b.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactorSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Permit(&_Atlas.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Permit(&_Atlas.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address searcherFrom, uint256 maxApprovedGasSpend) payable returns(bool)
func (_Atlas *AtlasTransactor) Reconcile(opts *bind.TransactOpts, environment common.Address, searcherFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "reconcile", environment, searcherFrom, maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address searcherFrom, uint256 maxApprovedGasSpend) payable returns(bool)
func (_Atlas *AtlasSession) Reconcile(environment common.Address, searcherFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, environment, searcherFrom, maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address searcherFrom, uint256 maxApprovedGasSpend) payable returns(bool)
func (_Atlas *AtlasTransactorSession) Reconcile(environment common.Address, searcherFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, environment, searcherFrom, maxApprovedGasSpend)
}

// RequestFrom is a paid mutator transaction binding the contract method 0x0baf06e5.
//
// Solidity: function requestFrom(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasTransactor) RequestFrom(opts *bind.TransactOpts, donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "requestFrom", donor, recipient, amt)
}

// RequestFrom is a paid mutator transaction binding the contract method 0x0baf06e5.
//
// Solidity: function requestFrom(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasSession) RequestFrom(donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.RequestFrom(&_Atlas.TransactOpts, donor, recipient, amt)
}

// RequestFrom is a paid mutator transaction binding the contract method 0x0baf06e5.
//
// Solidity: function requestFrom(uint8 donor, uint8 recipient, uint256 amt) returns()
func (_Atlas *AtlasTransactorSession) RequestFrom(donor uint8, recipient uint8, amt *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.RequestFrom(&_Atlas.TransactOpts, donor, recipient, amt)
}

// RequestGasFrom is a paid mutator transaction binding the contract method 0xe293db5e.
//
// Solidity: function requestGasFrom(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) RequestGasFrom(opts *bind.TransactOpts, donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "requestGasFrom", donor, recipient, amt, lockState)
}

// RequestGasFrom is a paid mutator transaction binding the contract method 0xe293db5e.
//
// Solidity: function requestGasFrom(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasSession) RequestGasFrom(donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.RequestGasFrom(&_Atlas.TransactOpts, donor, recipient, amt, lockState)
}

// RequestGasFrom is a paid mutator transaction binding the contract method 0xe293db5e.
//
// Solidity: function requestGasFrom(uint8 donor, uint8 recipient, uint256 amt, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) RequestGasFrom(donor uint8, recipient uint8, amt *big.Int, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.RequestGasFrom(&_Atlas.TransactOpts, donor, recipient, amt, lockState)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Transfer(&_Atlas.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Transfer(&_Atlas.TransactOpts, to, amount)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) TransferDAppERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferDAppERC20", token, destination, amount, user, controller, callConfig, lockState)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.TransferFrom(&_Atlas.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.TransferFrom(&_Atlas.TransactOpts, from, to, amount)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) TransferUserERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferUserERC20", token, destination, amount, user, controller, callConfig, lockState)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// ValidateBalances is a paid mutator transaction binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() returns(bool valid)
func (_Atlas *AtlasTransactor) ValidateBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "validateBalances")
}

// ValidateBalances is a paid mutator transaction binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() returns(bool valid)
func (_Atlas *AtlasSession) ValidateBalances() (*types.Transaction, error) {
	return _Atlas.Contract.ValidateBalances(&_Atlas.TransactOpts)
}

// ValidateBalances is a paid mutator transaction binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() returns(bool valid)
func (_Atlas *AtlasTransactorSession) ValidateBalances() (*types.Transaction, error) {
	return _Atlas.Contract.ValidateBalances(&_Atlas.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Withdraw(&_Atlas.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Withdraw(&_Atlas.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Atlas.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Atlas.Contract.Fallback(&_Atlas.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Atlas.Contract.Fallback(&_Atlas.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasSession) Receive() (*types.Transaction, error) {
	return _Atlas.Contract.Receive(&_Atlas.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasTransactorSession) Receive() (*types.Transaction, error) {
	return _Atlas.Contract.Receive(&_Atlas.TransactOpts)
}

// AtlasApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Atlas contract.
type AtlasApprovalIterator struct {
	Event *AtlasApproval // Event containing the contract specifics and raw log

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
func (it *AtlasApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasApproval)
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
		it.Event = new(AtlasApproval)
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
func (it *AtlasApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasApproval represents a Approval event raised by the Atlas contract.
type AtlasApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Atlas *AtlasFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AtlasApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AtlasApprovalIterator{contract: _Atlas.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Atlas *AtlasFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AtlasApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasApproval)
				if err := _Atlas.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Atlas *AtlasFilterer) ParseApproval(log types.Log) (*AtlasApproval, error) {
	event := new(AtlasApproval)
	if err := _Atlas.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasMEVPaymentFailureIterator is returned from FilterMEVPaymentFailure and is used to iterate over the raw logs and unpacked data for MEVPaymentFailure events raised by the Atlas contract.
type AtlasMEVPaymentFailureIterator struct {
	Event *AtlasMEVPaymentFailure // Event containing the contract specifics and raw log

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
func (it *AtlasMEVPaymentFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasMEVPaymentFailure)
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
		it.Event = new(AtlasMEVPaymentFailure)
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
func (it *AtlasMEVPaymentFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasMEVPaymentFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasMEVPaymentFailure represents a MEVPaymentFailure event raised by the Atlas contract.
type AtlasMEVPaymentFailure struct {
	Controller common.Address
	CallConfig uint32
	BidToken   common.Address
	BidAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMEVPaymentFailure is a free log retrieval operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) FilterMEVPaymentFailure(opts *bind.FilterOpts, controller []common.Address) (*AtlasMEVPaymentFailureIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasMEVPaymentFailureIterator{contract: _Atlas.contract, event: "MEVPaymentFailure", logs: logs, sub: sub}, nil
}

// WatchMEVPaymentFailure is a free log subscription operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) WatchMEVPaymentFailure(opts *bind.WatchOpts, sink chan<- *AtlasMEVPaymentFailure, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasMEVPaymentFailure)
				if err := _Atlas.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
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

// ParseMEVPaymentFailure is a log parse operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) ParseMEVPaymentFailure(log types.Log) (*AtlasMEVPaymentFailure, error) {
	event := new(AtlasMEVPaymentFailure)
	if err := _Atlas.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasNewExecutionEnvironmentIterator is returned from FilterNewExecutionEnvironment and is used to iterate over the raw logs and unpacked data for NewExecutionEnvironment events raised by the Atlas contract.
type AtlasNewExecutionEnvironmentIterator struct {
	Event *AtlasNewExecutionEnvironment // Event containing the contract specifics and raw log

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
func (it *AtlasNewExecutionEnvironmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasNewExecutionEnvironment)
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
		it.Event = new(AtlasNewExecutionEnvironment)
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
func (it *AtlasNewExecutionEnvironmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasNewExecutionEnvironmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasNewExecutionEnvironment represents a NewExecutionEnvironment event raised by the Atlas contract.
type AtlasNewExecutionEnvironment struct {
	Environment common.Address
	User        common.Address
	Controller  common.Address
	CallConfig  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionEnvironment is a free log retrieval operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_Atlas *AtlasFilterer) FilterNewExecutionEnvironment(opts *bind.FilterOpts, environment []common.Address, user []common.Address, controller []common.Address) (*AtlasNewExecutionEnvironmentIterator, error) {

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

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasNewExecutionEnvironmentIterator{contract: _Atlas.contract, event: "NewExecutionEnvironment", logs: logs, sub: sub}, nil
}

// WatchNewExecutionEnvironment is a free log subscription operation binding the contract event 0x01c1f90ed27e3f065842b640b4919261ebb9a550e8d92430a90ef13045789812.
//
// Solidity: event NewExecutionEnvironment(address indexed environment, address indexed user, address indexed controller, uint32 callConfig)
func (_Atlas *AtlasFilterer) WatchNewExecutionEnvironment(opts *bind.WatchOpts, sink chan<- *AtlasNewExecutionEnvironment, environment []common.Address, user []common.Address, controller []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "NewExecutionEnvironment", environmentRule, userRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasNewExecutionEnvironment)
				if err := _Atlas.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseNewExecutionEnvironment(log types.Log) (*AtlasNewExecutionEnvironment, error) {
	event := new(AtlasNewExecutionEnvironment)
	if err := _Atlas.contract.UnpackLog(event, "NewExecutionEnvironment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasSolverTxResultIterator is returned from FilterSolverTxResult and is used to iterate over the raw logs and unpacked data for SolverTxResult events raised by the Atlas contract.
type AtlasSolverTxResultIterator struct {
	Event *AtlasSolverTxResult // Event containing the contract specifics and raw log

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
func (it *AtlasSolverTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSolverTxResult)
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
		it.Event = new(AtlasSolverTxResult)
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
func (it *AtlasSolverTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSolverTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSolverTxResult represents a SolverTxResult event raised by the Atlas contract.
type AtlasSolverTxResult struct {
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
func (_Atlas *AtlasFilterer) FilterSolverTxResult(opts *bind.FilterOpts, solverTo []common.Address, solverFrom []common.Address) (*AtlasSolverTxResultIterator, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return &AtlasSolverTxResultIterator{contract: _Atlas.contract, event: "SolverTxResult", logs: logs, sub: sub}, nil
}

// WatchSolverTxResult is a free log subscription operation binding the contract event 0xeea2753019c05c54d8dc4ff90699ee5034f9ac243b7c2f9e3a0db08f6597a838.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 nonce, uint256 result)
func (_Atlas *AtlasFilterer) WatchSolverTxResult(opts *bind.WatchOpts, sink chan<- *AtlasSolverTxResult, solverTo []common.Address, solverFrom []common.Address) (event.Subscription, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSolverTxResult)
				if err := _Atlas.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseSolverTxResult(log types.Log) (*AtlasSolverTxResult, error) {
	event := new(AtlasSolverTxResult)
	if err := _Atlas.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Atlas contract.
type AtlasTransferIterator struct {
	Event *AtlasTransfer // Event containing the contract specifics and raw log

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
func (it *AtlasTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasTransfer)
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
		it.Event = new(AtlasTransfer)
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
func (it *AtlasTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasTransfer represents a Transfer event raised by the Atlas contract.
type AtlasTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Atlas *AtlasFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AtlasTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AtlasTransferIterator{contract: _Atlas.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Atlas *AtlasFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AtlasTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasTransfer)
				if err := _Atlas.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Atlas *AtlasFilterer) ParseTransfer(log types.Log) (*AtlasTransfer, error) {
	event := new(AtlasTransfer)
	if err := _Atlas.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasUserTxResultIterator is returned from FilterUserTxResult and is used to iterate over the raw logs and unpacked data for UserTxResult events raised by the Atlas contract.
type AtlasUserTxResultIterator struct {
	Event *AtlasUserTxResult // Event containing the contract specifics and raw log

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
func (it *AtlasUserTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasUserTxResult)
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
		it.Event = new(AtlasUserTxResult)
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
func (it *AtlasUserTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasUserTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasUserTxResult represents a UserTxResult event raised by the Atlas contract.
type AtlasUserTxResult struct {
	User          common.Address
	ValueReturned *big.Int
	GasRefunded   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserTxResult is a free log retrieval operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Atlas *AtlasFilterer) FilterUserTxResult(opts *bind.FilterOpts, user []common.Address) (*AtlasUserTxResultIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return &AtlasUserTxResultIterator{contract: _Atlas.contract, event: "UserTxResult", logs: logs, sub: sub}, nil
}

// WatchUserTxResult is a free log subscription operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Atlas *AtlasFilterer) WatchUserTxResult(opts *bind.WatchOpts, sink chan<- *AtlasUserTxResult, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasUserTxResult)
				if err := _Atlas.contract.UnpackLog(event, "UserTxResult", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseUserTxResult(log types.Log) (*AtlasUserTxResult, error) {
	event := new(AtlasUserTxResult)
	if err := _Atlas.contract.UnpackLog(event, "UserTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
