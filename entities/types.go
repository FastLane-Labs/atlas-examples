package entities

import (
	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type SwapIntentOperation struct {
	SwapIntent           *SwapIntentController.SwapIntent
	UserOperation        *Atlas.UserOperation
	ExecutionEnvironment common.Address
}

var (
	userOpType, _ = abi.NewType("tuple", "struct UserCall", []abi.ArgumentMarshaling{
		{Name: "From", Type: "address"},
		{Name: "To", Type: "address"},
		{Name: "Value", Type: "uint256"},
		{Name: "Gas", Type: "uint256"},
		{Name: "MaxFeePerGas", Type: "uint256"},
		{Name: "Nonce", Type: "uint256"},
		{Name: "Deadline", Type: "uint256"},
		{Name: "Control", Type: "address"},
		{Name: "Dapp", Type: "address"},
		{Name: "Data", Type: "bytes"},
		{Name: "Signature", Type: "bytes"},
	})

	userOpArg = abi.Arguments{{Type: userOpType}}
)

var (
	addressType, _ = abi.NewType("address", "address", nil)
	uint256Type, _ = abi.NewType("uint256", "uint256", nil)
	bytesType, _   = abi.NewType("bytes", "bytes", nil)
	boolType, _    = abi.NewType("bool", "bool", nil)

	v3SwapExactOut = abi.Arguments{
		abi.Argument{Name: "recipient", Type: addressType},
		abi.Argument{Name: "amountOut", Type: uint256Type},
		abi.Argument{Name: "amountInMaximum", Type: uint256Type},
		abi.Argument{Name: "path", Type: bytesType},
		abi.Argument{Name: "tokenInFromSender", Type: boolType},
	}
)
