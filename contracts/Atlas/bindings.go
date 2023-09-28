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

// BidData is an auto generated low-level Go binding around an user-defined struct.
type BidData struct {
	Token     common.Address
	BidAmount *big.Int
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
}

// DAppProof is an auto generated low-level Go binding around an user-defined struct.
type DAppProof struct {
	From            common.Address
	To              common.Address
	Nonce           *big.Int
	Deadline        *big.Int
	UserOpHash      [32]byte
	CallChainHash   [32]byte
	ControlCodeHash [32]byte
}

// SolverCall is an auto generated low-level Go binding around an user-defined struct.
type SolverCall struct {
	From            common.Address
	To              common.Address
	Value           *big.Int
	Gas             *big.Int
	Nonce           *big.Int
	MaxFeePerGas    *big.Int
	UserOpHash      [32]byte
	ControlCodeHash [32]byte
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
	Deadline     *big.Int
	Gas          *big.Int
	Nonce        *big.Int
	MaxFeePerGas *big.Int
	Value        *big.Int
	Control      common.Address
	Data         []byte
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	To        common.Address
	Call      UserCall
	Signature []byte
}

// Verification is an auto generated low-level Go binding around an user-defined struct.
type Verification struct {
	To        common.Address
	Proof     DAppProof
	Signature []byte
}

// AtlasMetaData contains all meta data concerning the Atlas contract.
var AtlasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_escrowDuration\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlteredControlHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredUserHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashChainBroken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentUnfulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolverHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuctionWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverBidUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverEVMError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverFailedCallback\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMsgValueUnpaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverOperationReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationSucceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserUnexpectedSuccess\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structBidData[]\",\"name\":\"winningBids\",\"type\":\"tuple[]\"}],\"name\":\"MEVPaymentFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"SolverTxResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasRefunded\",\"type\":\"uint256\"}],\"name\":\"UserTxResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BUNDLER_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUNDLER_PREMIUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAPP_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"addSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"asyncNonceFills\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"}],\"name\":\"createExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeDonations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dapps\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverSigner\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dappControl\",\"type\":\"address\"}],\"name\":\"disableDApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"surplusRecipient\",\"type\":\"address\"}],\"name\":\"donateToBundler\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_environment\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrowDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"uCall\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"auctionWon\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"accruedGasRebate\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"execution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEscrowAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"getExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"sCall\",\"type\":\"tuple\"}],\"name\":\"getSolverPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOperationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structVerification\",\"name\":\"verification\",\"type\":\"tuple\"}],\"name\":\"getVerificationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastUpdate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"initializeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dappControl\",\"type\":\"address\"}],\"name\":\"integrateDApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structVerification\",\"name\":\"verification\",\"type\":\"tuple\"}],\"name\":\"metacall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"auctionWon\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structVerification\",\"name\":\"verification\",\"type\":\"tuple\"}],\"name\":\"metacallSimulation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governanceSignatory\",\"type\":\"address\"}],\"name\":\"nextGovernanceNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverSigner\",\"type\":\"address\"}],\"name\":\"nextSolverNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"nextUserNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"removeSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signatories\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverSigner\",\"type\":\"address\"}],\"name\":\"solverEscrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverSigner\",\"type\":\"address\"}],\"name\":\"solverLastActiveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"solverSafetyCallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSafe\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolverCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"internalType\":\"structBidData[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"controlCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structDAppProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structVerification\",\"name\":\"verification\",\"type\":\"tuple\"}],\"name\":\"testSolverCalls\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"uCall\",\"type\":\"tuple\"}],\"name\":\"testUserOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"call\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"testUserOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUserCall\",\"name\":\"uCall\",\"type\":\"tuple\"}],\"name\":\"testUserOperationWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"transferDAppERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"lockState\",\"type\":\"uint16\"}],\"name\":\"transferUserERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// BUNDLERBASE is a free data retrieval call binding the contract method 0x5152f50f.
//
// Solidity: function BUNDLER_BASE() view returns(uint256)
func (_Atlas *AtlasCaller) BUNDLERBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "BUNDLER_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BUNDLERBASE is a free data retrieval call binding the contract method 0x5152f50f.
//
// Solidity: function BUNDLER_BASE() view returns(uint256)
func (_Atlas *AtlasSession) BUNDLERBASE() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERBASE(&_Atlas.CallOpts)
}

// BUNDLERBASE is a free data retrieval call binding the contract method 0x5152f50f.
//
// Solidity: function BUNDLER_BASE() view returns(uint256)
func (_Atlas *AtlasCallerSession) BUNDLERBASE() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERBASE(&_Atlas.CallOpts)
}

// BUNDLERPREMIUM is a free data retrieval call binding the contract method 0x191338b4.
//
// Solidity: function BUNDLER_PREMIUM() view returns(uint256)
func (_Atlas *AtlasCaller) BUNDLERPREMIUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "BUNDLER_PREMIUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BUNDLERPREMIUM is a free data retrieval call binding the contract method 0x191338b4.
//
// Solidity: function BUNDLER_PREMIUM() view returns(uint256)
func (_Atlas *AtlasSession) BUNDLERPREMIUM() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERPREMIUM(&_Atlas.CallOpts)
}

// BUNDLERPREMIUM is a free data retrieval call binding the contract method 0x191338b4.
//
// Solidity: function BUNDLER_PREMIUM() view returns(uint256)
func (_Atlas *AtlasCallerSession) BUNDLERPREMIUM() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERPREMIUM(&_Atlas.CallOpts)
}

// DAPPTYPEHASH is a free data retrieval call binding the contract method 0xac85596b.
//
// Solidity: function DAPP_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasCaller) DAPPTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "DAPP_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAPPTYPEHASH is a free data retrieval call binding the contract method 0xac85596b.
//
// Solidity: function DAPP_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasSession) DAPPTYPEHASH() ([32]byte, error) {
	return _Atlas.Contract.DAPPTYPEHASH(&_Atlas.CallOpts)
}

// DAPPTYPEHASH is a free data retrieval call binding the contract method 0xac85596b.
//
// Solidity: function DAPP_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasCallerSession) DAPPTYPEHASH() ([32]byte, error) {
	return _Atlas.Contract.DAPPTYPEHASH(&_Atlas.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_Atlas *AtlasCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_Atlas *AtlasSession) ISTEST() (bool, error) {
	return _Atlas.Contract.ISTEST(&_Atlas.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_Atlas *AtlasCallerSession) ISTEST() (bool, error) {
	return _Atlas.Contract.ISTEST(&_Atlas.CallOpts)
}

// USERTYPEHASH is a free data retrieval call binding the contract method 0x99a903a6.
//
// Solidity: function USER_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasCaller) USERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "USER_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USERTYPEHASH is a free data retrieval call binding the contract method 0x99a903a6.
//
// Solidity: function USER_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasSession) USERTYPEHASH() ([32]byte, error) {
	return _Atlas.Contract.USERTYPEHASH(&_Atlas.CallOpts)
}

// USERTYPEHASH is a free data retrieval call binding the contract method 0x99a903a6.
//
// Solidity: function USER_TYPE_HASH() view returns(bytes32)
func (_Atlas *AtlasCallerSession) USERTYPEHASH() ([32]byte, error) {
	return _Atlas.Contract.USERTYPEHASH(&_Atlas.CallOpts)
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

// AsyncNonceFills is a free data retrieval call binding the contract method 0x11a31507.
//
// Solidity: function asyncNonceFills(bytes32 ) view returns(address)
func (_Atlas *AtlasCaller) AsyncNonceFills(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "asyncNonceFills", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AsyncNonceFills is a free data retrieval call binding the contract method 0x11a31507.
//
// Solidity: function asyncNonceFills(bytes32 ) view returns(address)
func (_Atlas *AtlasSession) AsyncNonceFills(arg0 [32]byte) (common.Address, error) {
	return _Atlas.Contract.AsyncNonceFills(&_Atlas.CallOpts, arg0)
}

// AsyncNonceFills is a free data retrieval call binding the contract method 0x11a31507.
//
// Solidity: function asyncNonceFills(bytes32 ) view returns(address)
func (_Atlas *AtlasCallerSession) AsyncNonceFills(arg0 [32]byte) (common.Address, error) {
	return _Atlas.Contract.AsyncNonceFills(&_Atlas.CallOpts, arg0)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Atlas *AtlasCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Atlas *AtlasSession) Atlas() (common.Address, error) {
	return _Atlas.Contract.Atlas(&_Atlas.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Atlas *AtlasCallerSession) Atlas() (common.Address, error) {
	return _Atlas.Contract.Atlas(&_Atlas.CallOpts)
}

// CumulativeDonations is a free data retrieval call binding the contract method 0x5e437f58.
//
// Solidity: function cumulativeDonations() view returns(uint256)
func (_Atlas *AtlasCaller) CumulativeDonations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "cumulativeDonations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeDonations is a free data retrieval call binding the contract method 0x5e437f58.
//
// Solidity: function cumulativeDonations() view returns(uint256)
func (_Atlas *AtlasSession) CumulativeDonations() (*big.Int, error) {
	return _Atlas.Contract.CumulativeDonations(&_Atlas.CallOpts)
}

// CumulativeDonations is a free data retrieval call binding the contract method 0x5e437f58.
//
// Solidity: function cumulativeDonations() view returns(uint256)
func (_Atlas *AtlasCallerSession) CumulativeDonations() (*big.Int, error) {
	return _Atlas.Contract.CumulativeDonations(&_Atlas.CallOpts)
}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_Atlas *AtlasCaller) Dapps(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "dapps", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_Atlas *AtlasSession) Dapps(arg0 [32]byte) ([32]byte, error) {
	return _Atlas.Contract.Dapps(&_Atlas.CallOpts, arg0)
}

// Dapps is a free data retrieval call binding the contract method 0xb6453318.
//
// Solidity: function dapps(bytes32 ) view returns(bytes32)
func (_Atlas *AtlasCallerSession) Dapps(arg0 [32]byte) ([32]byte, error) {
	return _Atlas.Contract.Dapps(&_Atlas.CallOpts, arg0)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Atlas *AtlasCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "eip712Domain")

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
func (_Atlas *AtlasSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Atlas.Contract.Eip712Domain(&_Atlas.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Atlas *AtlasCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Atlas.Contract.Eip712Domain(&_Atlas.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address _environment)
func (_Atlas *AtlasCaller) Environment(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "environment")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address _environment)
func (_Atlas *AtlasSession) Environment() (common.Address, error) {
	return _Atlas.Contract.Environment(&_Atlas.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address _environment)
func (_Atlas *AtlasCallerSession) Environment() (common.Address, error) {
	return _Atlas.Contract.Environment(&_Atlas.CallOpts)
}

// EscrowDuration is a free data retrieval call binding the contract method 0x57c2c2ba.
//
// Solidity: function escrowDuration() view returns(uint32)
func (_Atlas *AtlasCaller) EscrowDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "escrowDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// EscrowDuration is a free data retrieval call binding the contract method 0x57c2c2ba.
//
// Solidity: function escrowDuration() view returns(uint32)
func (_Atlas *AtlasSession) EscrowDuration() (uint32, error) {
	return _Atlas.Contract.EscrowDuration(&_Atlas.CallOpts)
}

// EscrowDuration is a free data retrieval call binding the contract method 0x57c2c2ba.
//
// Solidity: function escrowDuration() view returns(uint32)
func (_Atlas *AtlasCallerSession) EscrowDuration() (uint32, error) {
	return _Atlas.Contract.EscrowDuration(&_Atlas.CallOpts)
}

// Execution is a free data retrieval call binding the contract method 0x4918f569.
//
// Solidity: function execution() view returns(address)
func (_Atlas *AtlasCaller) Execution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "execution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Execution is a free data retrieval call binding the contract method 0x4918f569.
//
// Solidity: function execution() view returns(address)
func (_Atlas *AtlasSession) Execution() (common.Address, error) {
	return _Atlas.Contract.Execution(&_Atlas.CallOpts)
}

// Execution is a free data retrieval call binding the contract method 0x4918f569.
//
// Solidity: function execution() view returns(address)
func (_Atlas *AtlasCallerSession) Execution() (common.Address, error) {
	return _Atlas.Contract.Execution(&_Atlas.CallOpts)
}

// GetEscrowAddress is a free data retrieval call binding the contract method 0x122fd48a.
//
// Solidity: function getEscrowAddress() view returns(address escrowAddress)
func (_Atlas *AtlasCaller) GetEscrowAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getEscrowAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEscrowAddress is a free data retrieval call binding the contract method 0x122fd48a.
//
// Solidity: function getEscrowAddress() view returns(address escrowAddress)
func (_Atlas *AtlasSession) GetEscrowAddress() (common.Address, error) {
	return _Atlas.Contract.GetEscrowAddress(&_Atlas.CallOpts)
}

// GetEscrowAddress is a free data retrieval call binding the contract method 0x122fd48a.
//
// Solidity: function getEscrowAddress() view returns(address escrowAddress)
func (_Atlas *AtlasCallerSession) GetEscrowAddress() (common.Address, error) {
	return _Atlas.Contract.GetEscrowAddress(&_Atlas.CallOpts)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x71ed2766.
//
// Solidity: function getExecutionEnvironment((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, address controller) view returns(address executionEnvironment)
func (_Atlas *AtlasCaller) GetExecutionEnvironment(opts *bind.CallOpts, userOp UserOperation, controller common.Address) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getExecutionEnvironment", userOp, controller)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x71ed2766.
//
// Solidity: function getExecutionEnvironment((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, address controller) view returns(address executionEnvironment)
func (_Atlas *AtlasSession) GetExecutionEnvironment(userOp UserOperation, controller common.Address) (common.Address, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, userOp, controller)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x71ed2766.
//
// Solidity: function getExecutionEnvironment((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, address controller) view returns(address executionEnvironment)
func (_Atlas *AtlasCallerSession) GetExecutionEnvironment(userOp UserOperation, controller common.Address) (common.Address, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, userOp, controller)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0x69a466fb.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes) sCall) view returns(bytes32 payload)
func (_Atlas *AtlasCaller) GetSolverPayload(opts *bind.CallOpts, sCall SolverCall) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getSolverPayload", sCall)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSolverPayload is a free data retrieval call binding the contract method 0x69a466fb.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes) sCall) view returns(bytes32 payload)
func (_Atlas *AtlasSession) GetSolverPayload(sCall SolverCall) ([32]byte, error) {
	return _Atlas.Contract.GetSolverPayload(&_Atlas.CallOpts, sCall)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0x69a466fb.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes) sCall) view returns(bytes32 payload)
func (_Atlas *AtlasCallerSession) GetSolverPayload(sCall SolverCall) ([32]byte, error) {
	return _Atlas.Contract.GetSolverPayload(&_Atlas.CallOpts, sCall)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x58e65507.
//
// Solidity: function getUserOperationPayload((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) view returns(bytes32 payload)
func (_Atlas *AtlasCaller) GetUserOperationPayload(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getUserOperationPayload", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x58e65507.
//
// Solidity: function getUserOperationPayload((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) view returns(bytes32 payload)
func (_Atlas *AtlasSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _Atlas.Contract.GetUserOperationPayload(&_Atlas.CallOpts, userOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x58e65507.
//
// Solidity: function getUserOperationPayload((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) view returns(bytes32 payload)
func (_Atlas *AtlasCallerSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _Atlas.Contract.GetUserOperationPayload(&_Atlas.CallOpts, userOp)
}

// GetVerificationPayload is a free data retrieval call binding the contract method 0x2cd89465.
//
// Solidity: function getVerificationPayload((address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) view returns(bytes32 payload)
func (_Atlas *AtlasCaller) GetVerificationPayload(opts *bind.CallOpts, verification Verification) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getVerificationPayload", verification)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerificationPayload is a free data retrieval call binding the contract method 0x2cd89465.
//
// Solidity: function getVerificationPayload((address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) view returns(bytes32 payload)
func (_Atlas *AtlasSession) GetVerificationPayload(verification Verification) ([32]byte, error) {
	return _Atlas.Contract.GetVerificationPayload(&_Atlas.CallOpts, verification)
}

// GetVerificationPayload is a free data retrieval call binding the contract method 0x2cd89465.
//
// Solidity: function getVerificationPayload((address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) view returns(bytes32 payload)
func (_Atlas *AtlasCallerSession) GetVerificationPayload(verification Verification) ([32]byte, error) {
	return _Atlas.Contract.GetVerificationPayload(&_Atlas.CallOpts, verification)
}

// Governance is a free data retrieval call binding the contract method 0x8426e6c1.
//
// Solidity: function governance(address ) view returns(address governance, uint32 callConfig, uint64 lastUpdate)
func (_Atlas *AtlasCaller) Governance(opts *bind.CallOpts, arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "governance", arg0)

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
func (_Atlas *AtlasSession) Governance(arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	return _Atlas.Contract.Governance(&_Atlas.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x8426e6c1.
//
// Solidity: function governance(address ) view returns(address governance, uint32 callConfig, uint64 lastUpdate)
func (_Atlas *AtlasCallerSession) Governance(arg0 common.Address) (struct {
	Governance common.Address
	CallConfig uint32
	LastUpdate uint64
}, error) {
	return _Atlas.Contract.Governance(&_Atlas.CallOpts, arg0)
}

// NextGovernanceNonce is a free data retrieval call binding the contract method 0x79a4fcd9.
//
// Solidity: function nextGovernanceNonce(address governanceSignatory) view returns(uint256 nextNonce)
func (_Atlas *AtlasCaller) NextGovernanceNonce(opts *bind.CallOpts, governanceSignatory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nextGovernanceNonce", governanceSignatory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextGovernanceNonce is a free data retrieval call binding the contract method 0x79a4fcd9.
//
// Solidity: function nextGovernanceNonce(address governanceSignatory) view returns(uint256 nextNonce)
func (_Atlas *AtlasSession) NextGovernanceNonce(governanceSignatory common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextGovernanceNonce(&_Atlas.CallOpts, governanceSignatory)
}

// NextGovernanceNonce is a free data retrieval call binding the contract method 0x79a4fcd9.
//
// Solidity: function nextGovernanceNonce(address governanceSignatory) view returns(uint256 nextNonce)
func (_Atlas *AtlasCallerSession) NextGovernanceNonce(governanceSignatory common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextGovernanceNonce(&_Atlas.CallOpts, governanceSignatory)
}

// NextSolverNonce is a free data retrieval call binding the contract method 0xcebf5a5c.
//
// Solidity: function nextSolverNonce(address solverSigner) view returns(uint256 nextNonce)
func (_Atlas *AtlasCaller) NextSolverNonce(opts *bind.CallOpts, solverSigner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nextSolverNonce", solverSigner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSolverNonce is a free data retrieval call binding the contract method 0xcebf5a5c.
//
// Solidity: function nextSolverNonce(address solverSigner) view returns(uint256 nextNonce)
func (_Atlas *AtlasSession) NextSolverNonce(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextSolverNonce(&_Atlas.CallOpts, solverSigner)
}

// NextSolverNonce is a free data retrieval call binding the contract method 0xcebf5a5c.
//
// Solidity: function nextSolverNonce(address solverSigner) view returns(uint256 nextNonce)
func (_Atlas *AtlasCallerSession) NextSolverNonce(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextSolverNonce(&_Atlas.CallOpts, solverSigner)
}

// NextUserNonce is a free data retrieval call binding the contract method 0xdb593a04.
//
// Solidity: function nextUserNonce(address user) view returns(uint256 nextNonce)
func (_Atlas *AtlasCaller) NextUserNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nextUserNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextUserNonce is a free data retrieval call binding the contract method 0xdb593a04.
//
// Solidity: function nextUserNonce(address user) view returns(uint256 nextNonce)
func (_Atlas *AtlasSession) NextUserNonce(user common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextUserNonce(&_Atlas.CallOpts, user)
}

// NextUserNonce is a free data retrieval call binding the contract method 0xdb593a04.
//
// Solidity: function nextUserNonce(address user) view returns(uint256 nextNonce)
func (_Atlas *AtlasCallerSession) NextUserNonce(user common.Address) (*big.Int, error) {
	return _Atlas.Contract.NextUserNonce(&_Atlas.CallOpts, user)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasSession) Salt() ([32]byte, error) {
	return _Atlas.Contract.Salt(&_Atlas.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasCallerSession) Salt() ([32]byte, error) {
	return _Atlas.Contract.Salt(&_Atlas.CallOpts)
}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(address governance, bool enabled, uint64 nonce)
func (_Atlas *AtlasCaller) Signatories(opts *bind.CallOpts, arg0 common.Address) (struct {
	Governance common.Address
	Enabled    bool
	Nonce      uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "signatories", arg0)

	outstruct := new(struct {
		Governance common.Address
		Enabled    bool
		Nonce      uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Governance = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Nonce = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(address governance, bool enabled, uint64 nonce)
func (_Atlas *AtlasSession) Signatories(arg0 common.Address) (struct {
	Governance common.Address
	Enabled    bool
	Nonce      uint64
}, error) {
	return _Atlas.Contract.Signatories(&_Atlas.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(address governance, bool enabled, uint64 nonce)
func (_Atlas *AtlasCallerSession) Signatories(arg0 common.Address) (struct {
	Governance common.Address
	Enabled    bool
	Nonce      uint64
}, error) {
	return _Atlas.Contract.Signatories(&_Atlas.CallOpts, arg0)
}

// SolverEscrowBalance is a free data retrieval call binding the contract method 0xe055cf3d.
//
// Solidity: function solverEscrowBalance(address solverSigner) view returns(uint256 balance)
func (_Atlas *AtlasCaller) SolverEscrowBalance(opts *bind.CallOpts, solverSigner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solverEscrowBalance", solverSigner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SolverEscrowBalance is a free data retrieval call binding the contract method 0xe055cf3d.
//
// Solidity: function solverEscrowBalance(address solverSigner) view returns(uint256 balance)
func (_Atlas *AtlasSession) SolverEscrowBalance(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.SolverEscrowBalance(&_Atlas.CallOpts, solverSigner)
}

// SolverEscrowBalance is a free data retrieval call binding the contract method 0xe055cf3d.
//
// Solidity: function solverEscrowBalance(address solverSigner) view returns(uint256 balance)
func (_Atlas *AtlasCallerSession) SolverEscrowBalance(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.SolverEscrowBalance(&_Atlas.CallOpts, solverSigner)
}

// SolverLastActiveBlock is a free data retrieval call binding the contract method 0x8deeab9d.
//
// Solidity: function solverLastActiveBlock(address solverSigner) view returns(uint256 lastBlock)
func (_Atlas *AtlasCaller) SolverLastActiveBlock(opts *bind.CallOpts, solverSigner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solverLastActiveBlock", solverSigner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SolverLastActiveBlock is a free data retrieval call binding the contract method 0x8deeab9d.
//
// Solidity: function solverLastActiveBlock(address solverSigner) view returns(uint256 lastBlock)
func (_Atlas *AtlasSession) SolverLastActiveBlock(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.SolverLastActiveBlock(&_Atlas.CallOpts, solverSigner)
}

// SolverLastActiveBlock is a free data retrieval call binding the contract method 0x8deeab9d.
//
// Solidity: function solverLastActiveBlock(address solverSigner) view returns(uint256 lastBlock)
func (_Atlas *AtlasCallerSession) SolverLastActiveBlock(solverSigner common.Address) (*big.Int, error) {
	return _Atlas.Contract.SolverLastActiveBlock(&_Atlas.CallOpts, solverSigner)
}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint256)
func (_Atlas *AtlasCaller) UserNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "userNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint256)
func (_Atlas *AtlasSession) UserNonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.UserNonces(&_Atlas.CallOpts, arg0)
}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint256)
func (_Atlas *AtlasCallerSession) UserNonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.UserNonces(&_Atlas.CallOpts, arg0)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasTransactor) AddSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "addSignatory", controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.AddSignatory(&_Atlas.TransactOpts, controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasTransactorSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.AddSignatory(&_Atlas.TransactOpts, controller, signatory)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x0e1bc033.
//
// Solidity: function createExecutionEnvironment((address,uint32) dConfig) returns(address executionEnvironment)
func (_Atlas *AtlasTransactor) CreateExecutionEnvironment(opts *bind.TransactOpts, dConfig DAppConfig) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "createExecutionEnvironment", dConfig)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x0e1bc033.
//
// Solidity: function createExecutionEnvironment((address,uint32) dConfig) returns(address executionEnvironment)
func (_Atlas *AtlasSession) CreateExecutionEnvironment(dConfig DAppConfig) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dConfig)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x0e1bc033.
//
// Solidity: function createExecutionEnvironment((address,uint32) dConfig) returns(address executionEnvironment)
func (_Atlas *AtlasTransactorSession) CreateExecutionEnvironment(dConfig DAppConfig) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dConfig)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address solverSigner) payable returns(uint256 newBalance)
func (_Atlas *AtlasTransactor) Deposit(opts *bind.TransactOpts, solverSigner common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "deposit", solverSigner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address solverSigner) payable returns(uint256 newBalance)
func (_Atlas *AtlasSession) Deposit(solverSigner common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts, solverSigner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address solverSigner) payable returns(uint256 newBalance)
func (_Atlas *AtlasTransactorSession) Deposit(solverSigner common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts, solverSigner)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x0041bddf.
//
// Solidity: function disableDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasTransactor) DisableDApp(opts *bind.TransactOpts, controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "disableDApp", controller, dappControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x0041bddf.
//
// Solidity: function disableDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasSession) DisableDApp(controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.DisableDApp(&_Atlas.TransactOpts, controller, dappControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x0041bddf.
//
// Solidity: function disableDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasTransactorSession) DisableDApp(controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.DisableDApp(&_Atlas.TransactOpts, controller, dappControl)
}

// DonateToBundler is a paid mutator transaction binding the contract method 0x8b2b6f6e.
//
// Solidity: function donateToBundler(address surplusRecipient) payable returns()
func (_Atlas *AtlasTransactor) DonateToBundler(opts *bind.TransactOpts, surplusRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "donateToBundler", surplusRecipient)
}

// DonateToBundler is a paid mutator transaction binding the contract method 0x8b2b6f6e.
//
// Solidity: function donateToBundler(address surplusRecipient) payable returns()
func (_Atlas *AtlasSession) DonateToBundler(surplusRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.DonateToBundler(&_Atlas.TransactOpts, surplusRecipient)
}

// DonateToBundler is a paid mutator transaction binding the contract method 0x8b2b6f6e.
//
// Solidity: function donateToBundler(address surplusRecipient) payable returns()
func (_Atlas *AtlasTransactorSession) DonateToBundler(surplusRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.DonateToBundler(&_Atlas.TransactOpts, surplusRecipient)
}

// Execute is a paid mutator transaction binding the contract method 0xdd8ec6f0.
//
// Solidity: function execute((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, address executionEnvironment, bytes32 callChainHash) payable returns(bool auctionWon, uint256 accruedGasRebate)
func (_Atlas *AtlasTransactor) Execute(opts *bind.TransactOpts, dConfig DAppConfig, uCall UserCall, solverOps []SolverOperation, executionEnvironment common.Address, callChainHash [32]byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "execute", dConfig, uCall, solverOps, executionEnvironment, callChainHash)
}

// Execute is a paid mutator transaction binding the contract method 0xdd8ec6f0.
//
// Solidity: function execute((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, address executionEnvironment, bytes32 callChainHash) payable returns(bool auctionWon, uint256 accruedGasRebate)
func (_Atlas *AtlasSession) Execute(dConfig DAppConfig, uCall UserCall, solverOps []SolverOperation, executionEnvironment common.Address, callChainHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, uCall, solverOps, executionEnvironment, callChainHash)
}

// Execute is a paid mutator transaction binding the contract method 0xdd8ec6f0.
//
// Solidity: function execute((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, address executionEnvironment, bytes32 callChainHash) payable returns(bool auctionWon, uint256 accruedGasRebate)
func (_Atlas *AtlasTransactorSession) Execute(dConfig DAppConfig, uCall UserCall, solverOps []SolverOperation, executionEnvironment common.Address, callChainHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, uCall, solverOps, executionEnvironment, callChainHash)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_Atlas *AtlasTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_Atlas *AtlasSession) Failed() (*types.Transaction, error) {
	return _Atlas.Contract.Failed(&_Atlas.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_Atlas *AtlasTransactorSession) Failed() (*types.Transaction, error) {
	return _Atlas.Contract.Failed(&_Atlas.TransactOpts)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_Atlas *AtlasTransactor) InitializeGovernance(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "initializeGovernance", controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_Atlas *AtlasSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.InitializeGovernance(&_Atlas.TransactOpts, controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_Atlas *AtlasTransactorSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.InitializeGovernance(&_Atlas.TransactOpts, controller)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x9033d096.
//
// Solidity: function integrateDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasTransactor) IntegrateDApp(opts *bind.TransactOpts, controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "integrateDApp", controller, dappControl)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x9033d096.
//
// Solidity: function integrateDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasSession) IntegrateDApp(controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.IntegrateDApp(&_Atlas.TransactOpts, controller, dappControl)
}

// IntegrateDApp is a paid mutator transaction binding the contract method 0x9033d096.
//
// Solidity: function integrateDApp(address controller, address dappControl) returns()
func (_Atlas *AtlasTransactorSession) IntegrateDApp(controller common.Address, dappControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.IntegrateDApp(&_Atlas.TransactOpts, controller, dappControl)
}

// Metacall is a paid mutator transaction binding the contract method 0xc42242b2.
//
// Solidity: function metacall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactor) Metacall(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "metacall", dConfig, userOp, solverOps, verification)
}

// Metacall is a paid mutator transaction binding the contract method 0xc42242b2.
//
// Solidity: function metacall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool auctionWon)
func (_Atlas *AtlasSession) Metacall(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// Metacall is a paid mutator transaction binding the contract method 0xc42242b2.
//
// Solidity: function metacall((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactorSession) Metacall(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x93e13d71.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns()
func (_Atlas *AtlasTransactor) MetacallSimulation(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "metacallSimulation", dConfig, userOp, solverOps, verification)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x93e13d71.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns()
func (_Atlas *AtlasSession) MetacallSimulation(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.MetacallSimulation(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x93e13d71.
//
// Solidity: function metacallSimulation((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns()
func (_Atlas *AtlasTransactorSession) MetacallSimulation(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.MetacallSimulation(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasTransactor) RemoveSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "removeSignatory", controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.RemoveSignatory(&_Atlas.TransactOpts, controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_Atlas *AtlasTransactorSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.RemoveSignatory(&_Atlas.TransactOpts, controller, signatory)
}

// SolverSafetyCallback is a paid mutator transaction binding the contract method 0x0340c6c5.
//
// Solidity: function solverSafetyCallback(address msgSender) payable returns(bool isSafe)
func (_Atlas *AtlasTransactor) SolverSafetyCallback(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "solverSafetyCallback", msgSender)
}

// SolverSafetyCallback is a paid mutator transaction binding the contract method 0x0340c6c5.
//
// Solidity: function solverSafetyCallback(address msgSender) payable returns(bool isSafe)
func (_Atlas *AtlasSession) SolverSafetyCallback(msgSender common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.SolverSafetyCallback(&_Atlas.TransactOpts, msgSender)
}

// SolverSafetyCallback is a paid mutator transaction binding the contract method 0x0340c6c5.
//
// Solidity: function solverSafetyCallback(address msgSender) payable returns(bool isSafe)
func (_Atlas *AtlasTransactorSession) SolverSafetyCallback(msgSender common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.SolverSafetyCallback(&_Atlas.TransactOpts, msgSender)
}

// TestSolverCalls is a paid mutator transaction binding the contract method 0x4697c29b.
//
// Solidity: function testSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool success)
func (_Atlas *AtlasTransactor) TestSolverCalls(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "testSolverCalls", dConfig, userOp, solverOps, verification)
}

// TestSolverCalls is a paid mutator transaction binding the contract method 0x4697c29b.
//
// Solidity: function testSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool success)
func (_Atlas *AtlasSession) TestSolverCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.TestSolverCalls(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// TestSolverCalls is a paid mutator transaction binding the contract method 0x4697c29b.
//
// Solidity: function testSolverCalls((address,uint32) dConfig, (address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp, (address,(address,address,uint256,uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes),bytes,(address,uint256)[])[] solverOps, (address,(address,address,uint256,uint256,bytes32,bytes32,bytes32),bytes) verification) payable returns(bool success)
func (_Atlas *AtlasTransactorSession) TestSolverCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, verification Verification) (*types.Transaction, error) {
	return _Atlas.Contract.TestSolverCalls(&_Atlas.TransactOpts, dConfig, userOp, solverOps, verification)
}

// TestUserOperation is a paid mutator transaction binding the contract method 0x16c2bac3.
//
// Solidity: function testUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns(bool)
func (_Atlas *AtlasTransactor) TestUserOperation(opts *bind.TransactOpts, uCall UserCall) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "testUserOperation", uCall)
}

// TestUserOperation is a paid mutator transaction binding the contract method 0x16c2bac3.
//
// Solidity: function testUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns(bool)
func (_Atlas *AtlasSession) TestUserOperation(uCall UserCall) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperation(&_Atlas.TransactOpts, uCall)
}

// TestUserOperation is a paid mutator transaction binding the contract method 0x16c2bac3.
//
// Solidity: function testUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns(bool)
func (_Atlas *AtlasTransactorSession) TestUserOperation(uCall UserCall) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperation(&_Atlas.TransactOpts, uCall)
}

// TestUserOperation0 is a paid mutator transaction binding the contract method 0xd2a8c540.
//
// Solidity: function testUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) returns(bool)
func (_Atlas *AtlasTransactor) TestUserOperation0(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "testUserOperation0", userOp)
}

// TestUserOperation0 is a paid mutator transaction binding the contract method 0xd2a8c540.
//
// Solidity: function testUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) returns(bool)
func (_Atlas *AtlasSession) TestUserOperation0(userOp UserOperation) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperation0(&_Atlas.TransactOpts, userOp)
}

// TestUserOperation0 is a paid mutator transaction binding the contract method 0xd2a8c540.
//
// Solidity: function testUserOperation((address,(address,address,uint256,uint256,uint256,uint256,uint256,address,bytes),bytes) userOp) returns(bool)
func (_Atlas *AtlasTransactorSession) TestUserOperation0(userOp UserOperation) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperation0(&_Atlas.TransactOpts, userOp)
}

// TestUserOperationWrapper is a paid mutator transaction binding the contract method 0xefa1f647.
//
// Solidity: function testUserOperationWrapper((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns()
func (_Atlas *AtlasTransactor) TestUserOperationWrapper(opts *bind.TransactOpts, dConfig DAppConfig, uCall UserCall) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "testUserOperationWrapper", dConfig, uCall)
}

// TestUserOperationWrapper is a paid mutator transaction binding the contract method 0xefa1f647.
//
// Solidity: function testUserOperationWrapper((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns()
func (_Atlas *AtlasSession) TestUserOperationWrapper(dConfig DAppConfig, uCall UserCall) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperationWrapper(&_Atlas.TransactOpts, dConfig, uCall)
}

// TestUserOperationWrapper is a paid mutator transaction binding the contract method 0xefa1f647.
//
// Solidity: function testUserOperationWrapper((address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes) uCall) returns()
func (_Atlas *AtlasTransactorSession) TestUserOperationWrapper(dConfig DAppConfig, uCall UserCall) (*types.Transaction, error) {
	return _Atlas.Contract.TestUserOperationWrapper(&_Atlas.TransactOpts, dConfig, uCall)
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

// AtlasEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Atlas contract.
type AtlasEIP712DomainChangedIterator struct {
	Event *AtlasEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AtlasEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasEIP712DomainChanged)
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
		it.Event = new(AtlasEIP712DomainChanged)
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
func (it *AtlasEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasEIP712DomainChanged represents a EIP712DomainChanged event raised by the Atlas contract.
type AtlasEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Atlas *AtlasFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AtlasEIP712DomainChangedIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AtlasEIP712DomainChangedIterator{contract: _Atlas.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Atlas *AtlasFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AtlasEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasEIP712DomainChanged)
				if err := _Atlas.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseEIP712DomainChanged(log types.Log) (*AtlasEIP712DomainChanged, error) {
	event := new(AtlasEIP712DomainChanged)
	if err := _Atlas.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
	Controller  common.Address
	CallConfig  uint32
	WinningBids []BidData
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMEVPaymentFailure is a free log retrieval operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
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

// WatchMEVPaymentFailure is a free log subscription operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
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

// ParseMEVPaymentFailure is a log parse operation binding the contract event 0x9c57c9b57eeb94cc2ff30fa4d78c17dd15eeb124a334d726ca964d075257684b.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, (address,uint256)[] winningBids)
func (_Atlas *AtlasFilterer) ParseMEVPaymentFailure(log types.Log) (*AtlasMEVPaymentFailure, error) {
	event := new(AtlasMEVPaymentFailure)
	if err := _Atlas.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
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

// AtlasLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Atlas contract.
type AtlasLogIterator struct {
	Event *AtlasLog // Event containing the contract specifics and raw log

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
func (it *AtlasLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLog)
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
		it.Event = new(AtlasLog)
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
func (it *AtlasLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLog represents a Log event raised by the Atlas contract.
type AtlasLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_Atlas *AtlasFilterer) FilterLog(opts *bind.FilterOpts) (*AtlasLogIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &AtlasLogIterator{contract: _Atlas.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_Atlas *AtlasFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *AtlasLog) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLog)
				if err := _Atlas.contract.UnpackLog(event, "log", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLog(log types.Log) (*AtlasLog, error) {
	event := new(AtlasLog)
	if err := _Atlas.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the Atlas contract.
type AtlasLogAddressIterator struct {
	Event *AtlasLogAddress // Event containing the contract specifics and raw log

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
func (it *AtlasLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogAddress)
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
		it.Event = new(AtlasLogAddress)
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
func (it *AtlasLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogAddress represents a LogAddress event raised by the Atlas contract.
type AtlasLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_Atlas *AtlasFilterer) FilterLogAddress(opts *bind.FilterOpts) (*AtlasLogAddressIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &AtlasLogAddressIterator{contract: _Atlas.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_Atlas *AtlasFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *AtlasLogAddress) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogAddress)
				if err := _Atlas.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogAddress(log types.Log) (*AtlasLogAddress, error) {
	event := new(AtlasLogAddress)
	if err := _Atlas.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the Atlas contract.
type AtlasLogArrayIterator struct {
	Event *AtlasLogArray // Event containing the contract specifics and raw log

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
func (it *AtlasLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogArray)
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
		it.Event = new(AtlasLogArray)
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
func (it *AtlasLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogArray represents a LogArray event raised by the Atlas contract.
type AtlasLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_Atlas *AtlasFilterer) FilterLogArray(opts *bind.FilterOpts) (*AtlasLogArrayIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &AtlasLogArrayIterator{contract: _Atlas.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_Atlas *AtlasFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *AtlasLogArray) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogArray)
				if err := _Atlas.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogArray(log types.Log) (*AtlasLogArray, error) {
	event := new(AtlasLogArray)
	if err := _Atlas.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the Atlas contract.
type AtlasLogArray0Iterator struct {
	Event *AtlasLogArray0 // Event containing the contract specifics and raw log

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
func (it *AtlasLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogArray0)
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
		it.Event = new(AtlasLogArray0)
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
func (it *AtlasLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogArray0 represents a LogArray0 event raised by the Atlas contract.
type AtlasLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_Atlas *AtlasFilterer) FilterLogArray0(opts *bind.FilterOpts) (*AtlasLogArray0Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &AtlasLogArray0Iterator{contract: _Atlas.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_Atlas *AtlasFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *AtlasLogArray0) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogArray0)
				if err := _Atlas.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogArray0(log types.Log) (*AtlasLogArray0, error) {
	event := new(AtlasLogArray0)
	if err := _Atlas.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the Atlas contract.
type AtlasLogArray1Iterator struct {
	Event *AtlasLogArray1 // Event containing the contract specifics and raw log

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
func (it *AtlasLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogArray1)
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
		it.Event = new(AtlasLogArray1)
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
func (it *AtlasLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogArray1 represents a LogArray1 event raised by the Atlas contract.
type AtlasLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_Atlas *AtlasFilterer) FilterLogArray1(opts *bind.FilterOpts) (*AtlasLogArray1Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &AtlasLogArray1Iterator{contract: _Atlas.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_Atlas *AtlasFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *AtlasLogArray1) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogArray1)
				if err := _Atlas.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogArray1(log types.Log) (*AtlasLogArray1, error) {
	event := new(AtlasLogArray1)
	if err := _Atlas.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the Atlas contract.
type AtlasLogBytesIterator struct {
	Event *AtlasLogBytes // Event containing the contract specifics and raw log

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
func (it *AtlasLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogBytes)
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
		it.Event = new(AtlasLogBytes)
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
func (it *AtlasLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogBytes represents a LogBytes event raised by the Atlas contract.
type AtlasLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_Atlas *AtlasFilterer) FilterLogBytes(opts *bind.FilterOpts) (*AtlasLogBytesIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &AtlasLogBytesIterator{contract: _Atlas.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_Atlas *AtlasFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *AtlasLogBytes) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogBytes)
				if err := _Atlas.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogBytes(log types.Log) (*AtlasLogBytes, error) {
	event := new(AtlasLogBytes)
	if err := _Atlas.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the Atlas contract.
type AtlasLogBytes32Iterator struct {
	Event *AtlasLogBytes32 // Event containing the contract specifics and raw log

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
func (it *AtlasLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogBytes32)
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
		it.Event = new(AtlasLogBytes32)
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
func (it *AtlasLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogBytes32 represents a LogBytes32 event raised by the Atlas contract.
type AtlasLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_Atlas *AtlasFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*AtlasLogBytes32Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &AtlasLogBytes32Iterator{contract: _Atlas.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_Atlas *AtlasFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *AtlasLogBytes32) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogBytes32)
				if err := _Atlas.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogBytes32(log types.Log) (*AtlasLogBytes32, error) {
	event := new(AtlasLogBytes32)
	if err := _Atlas.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the Atlas contract.
type AtlasLogIntIterator struct {
	Event *AtlasLogInt // Event containing the contract specifics and raw log

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
func (it *AtlasLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogInt)
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
		it.Event = new(AtlasLogInt)
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
func (it *AtlasLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogInt represents a LogInt event raised by the Atlas contract.
type AtlasLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_Atlas *AtlasFilterer) FilterLogInt(opts *bind.FilterOpts) (*AtlasLogIntIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &AtlasLogIntIterator{contract: _Atlas.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_Atlas *AtlasFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *AtlasLogInt) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogInt)
				if err := _Atlas.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogInt(log types.Log) (*AtlasLogInt, error) {
	event := new(AtlasLogInt)
	if err := _Atlas.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the Atlas contract.
type AtlasLogNamedAddressIterator struct {
	Event *AtlasLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedAddress)
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
		it.Event = new(AtlasLogNamedAddress)
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
func (it *AtlasLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedAddress represents a LogNamedAddress event raised by the Atlas contract.
type AtlasLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_Atlas *AtlasFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*AtlasLogNamedAddressIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedAddressIterator{contract: _Atlas.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_Atlas *AtlasFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedAddress)
				if err := _Atlas.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedAddress(log types.Log) (*AtlasLogNamedAddress, error) {
	event := new(AtlasLogNamedAddress)
	if err := _Atlas.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the Atlas contract.
type AtlasLogNamedArrayIterator struct {
	Event *AtlasLogNamedArray // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedArray)
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
		it.Event = new(AtlasLogNamedArray)
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
func (it *AtlasLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedArray represents a LogNamedArray event raised by the Atlas contract.
type AtlasLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_Atlas *AtlasFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*AtlasLogNamedArrayIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedArrayIterator{contract: _Atlas.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_Atlas *AtlasFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedArray)
				if err := _Atlas.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedArray(log types.Log) (*AtlasLogNamedArray, error) {
	event := new(AtlasLogNamedArray)
	if err := _Atlas.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the Atlas contract.
type AtlasLogNamedArray0Iterator struct {
	Event *AtlasLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedArray0)
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
		it.Event = new(AtlasLogNamedArray0)
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
func (it *AtlasLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedArray0 represents a LogNamedArray0 event raised by the Atlas contract.
type AtlasLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_Atlas *AtlasFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*AtlasLogNamedArray0Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedArray0Iterator{contract: _Atlas.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_Atlas *AtlasFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedArray0)
				if err := _Atlas.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedArray0(log types.Log) (*AtlasLogNamedArray0, error) {
	event := new(AtlasLogNamedArray0)
	if err := _Atlas.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the Atlas contract.
type AtlasLogNamedArray1Iterator struct {
	Event *AtlasLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedArray1)
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
		it.Event = new(AtlasLogNamedArray1)
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
func (it *AtlasLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedArray1 represents a LogNamedArray1 event raised by the Atlas contract.
type AtlasLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_Atlas *AtlasFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*AtlasLogNamedArray1Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedArray1Iterator{contract: _Atlas.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_Atlas *AtlasFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedArray1)
				if err := _Atlas.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedArray1(log types.Log) (*AtlasLogNamedArray1, error) {
	event := new(AtlasLogNamedArray1)
	if err := _Atlas.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the Atlas contract.
type AtlasLogNamedBytesIterator struct {
	Event *AtlasLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedBytes)
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
		it.Event = new(AtlasLogNamedBytes)
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
func (it *AtlasLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedBytes represents a LogNamedBytes event raised by the Atlas contract.
type AtlasLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_Atlas *AtlasFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*AtlasLogNamedBytesIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedBytesIterator{contract: _Atlas.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_Atlas *AtlasFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedBytes)
				if err := _Atlas.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedBytes(log types.Log) (*AtlasLogNamedBytes, error) {
	event := new(AtlasLogNamedBytes)
	if err := _Atlas.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the Atlas contract.
type AtlasLogNamedBytes32Iterator struct {
	Event *AtlasLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedBytes32)
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
		it.Event = new(AtlasLogNamedBytes32)
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
func (it *AtlasLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedBytes32 represents a LogNamedBytes32 event raised by the Atlas contract.
type AtlasLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_Atlas *AtlasFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*AtlasLogNamedBytes32Iterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedBytes32Iterator{contract: _Atlas.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_Atlas *AtlasFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedBytes32)
				if err := _Atlas.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedBytes32(log types.Log) (*AtlasLogNamedBytes32, error) {
	event := new(AtlasLogNamedBytes32)
	if err := _Atlas.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the Atlas contract.
type AtlasLogNamedDecimalIntIterator struct {
	Event *AtlasLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedDecimalInt)
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
		it.Event = new(AtlasLogNamedDecimalInt)
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
func (it *AtlasLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the Atlas contract.
type AtlasLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_Atlas *AtlasFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*AtlasLogNamedDecimalIntIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedDecimalIntIterator{contract: _Atlas.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_Atlas *AtlasFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedDecimalInt)
				if err := _Atlas.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedDecimalInt(log types.Log) (*AtlasLogNamedDecimalInt, error) {
	event := new(AtlasLogNamedDecimalInt)
	if err := _Atlas.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the Atlas contract.
type AtlasLogNamedDecimalUintIterator struct {
	Event *AtlasLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedDecimalUint)
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
		it.Event = new(AtlasLogNamedDecimalUint)
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
func (it *AtlasLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the Atlas contract.
type AtlasLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_Atlas *AtlasFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*AtlasLogNamedDecimalUintIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedDecimalUintIterator{contract: _Atlas.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_Atlas *AtlasFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedDecimalUint)
				if err := _Atlas.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedDecimalUint(log types.Log) (*AtlasLogNamedDecimalUint, error) {
	event := new(AtlasLogNamedDecimalUint)
	if err := _Atlas.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the Atlas contract.
type AtlasLogNamedIntIterator struct {
	Event *AtlasLogNamedInt // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedInt)
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
		it.Event = new(AtlasLogNamedInt)
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
func (it *AtlasLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedInt represents a LogNamedInt event raised by the Atlas contract.
type AtlasLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_Atlas *AtlasFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*AtlasLogNamedIntIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedIntIterator{contract: _Atlas.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_Atlas *AtlasFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedInt)
				if err := _Atlas.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedInt(log types.Log) (*AtlasLogNamedInt, error) {
	event := new(AtlasLogNamedInt)
	if err := _Atlas.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the Atlas contract.
type AtlasLogNamedStringIterator struct {
	Event *AtlasLogNamedString // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedString)
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
		it.Event = new(AtlasLogNamedString)
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
func (it *AtlasLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedString represents a LogNamedString event raised by the Atlas contract.
type AtlasLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_Atlas *AtlasFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*AtlasLogNamedStringIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedStringIterator{contract: _Atlas.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_Atlas *AtlasFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedString) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedString)
				if err := _Atlas.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedString(log types.Log) (*AtlasLogNamedString, error) {
	event := new(AtlasLogNamedString)
	if err := _Atlas.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the Atlas contract.
type AtlasLogNamedUintIterator struct {
	Event *AtlasLogNamedUint // Event containing the contract specifics and raw log

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
func (it *AtlasLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogNamedUint)
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
		it.Event = new(AtlasLogNamedUint)
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
func (it *AtlasLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogNamedUint represents a LogNamedUint event raised by the Atlas contract.
type AtlasLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_Atlas *AtlasFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*AtlasLogNamedUintIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &AtlasLogNamedUintIterator{contract: _Atlas.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_Atlas *AtlasFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *AtlasLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogNamedUint)
				if err := _Atlas.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogNamedUint(log types.Log) (*AtlasLogNamedUint, error) {
	event := new(AtlasLogNamedUint)
	if err := _Atlas.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the Atlas contract.
type AtlasLogStringIterator struct {
	Event *AtlasLogString // Event containing the contract specifics and raw log

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
func (it *AtlasLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogString)
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
		it.Event = new(AtlasLogString)
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
func (it *AtlasLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogString represents a LogString event raised by the Atlas contract.
type AtlasLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_Atlas *AtlasFilterer) FilterLogString(opts *bind.FilterOpts) (*AtlasLogStringIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &AtlasLogStringIterator{contract: _Atlas.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_Atlas *AtlasFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *AtlasLogString) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogString)
				if err := _Atlas.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogString(log types.Log) (*AtlasLogString, error) {
	event := new(AtlasLogString)
	if err := _Atlas.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the Atlas contract.
type AtlasLogUintIterator struct {
	Event *AtlasLogUint // Event containing the contract specifics and raw log

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
func (it *AtlasLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogUint)
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
		it.Event = new(AtlasLogUint)
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
func (it *AtlasLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogUint represents a LogUint event raised by the Atlas contract.
type AtlasLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_Atlas *AtlasFilterer) FilterLogUint(opts *bind.FilterOpts) (*AtlasLogUintIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &AtlasLogUintIterator{contract: _Atlas.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_Atlas *AtlasFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *AtlasLogUint) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogUint)
				if err := _Atlas.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogUint(log types.Log) (*AtlasLogUint, error) {
	event := new(AtlasLogUint)
	if err := _Atlas.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the Atlas contract.
type AtlasLogsIterator struct {
	Event *AtlasLogs // Event containing the contract specifics and raw log

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
func (it *AtlasLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasLogs)
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
		it.Event = new(AtlasLogs)
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
func (it *AtlasLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasLogs represents a Logs event raised by the Atlas contract.
type AtlasLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_Atlas *AtlasFilterer) FilterLogs(opts *bind.FilterOpts) (*AtlasLogsIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &AtlasLogsIterator{contract: _Atlas.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_Atlas *AtlasFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *AtlasLogs) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasLogs)
				if err := _Atlas.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseLogs(log types.Log) (*AtlasLogs, error) {
	event := new(AtlasLogs)
	if err := _Atlas.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
