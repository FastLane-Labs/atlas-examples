package account

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type User struct {
	Signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	atlas     *Atlas.Atlas
	txBuilder *TxBuilder.TxBuilder

	addresses map[string]common.Address

	ethClient *ethclient.Client
}

func NewUser(pk string, atlas *Atlas.Atlas, txBuilder *TxBuilder.TxBuilder, addresses map[string]common.Address, ethClient *ethclient.Client) *User {
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		log.Fatalf("could not load user's private key: %s", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, common.Big1)
	if err != nil {
		log.Fatalf("could not initialize user's signer: %s", err)
	}

	return &User{
		Signer:     signer,
		privateKey: privateKey,
		atlas:      atlas,
		txBuilder:  txBuilder,
		addresses:  addresses,
		ethClient:  ethClient,
	}
}

func (u *User) ApproveErc20Atlas(tokenAddress common.Address, amount *big.Int) {
	erc20, err := ERC20.NewERC20(tokenAddress, u.ethClient)
	if err != nil {
		log.Fatalf("could not load ERC20 contract: %s", err)
	}

	u.Signer.Value = new(big.Int).Set(common.Big0)
	tx, err := erc20.Approve(u.Signer, u.addresses["atlas"], amount)
	if err != nil {
		log.Fatalf("could not approve ERC20 for user: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for ERC20 approval transaction to be mined: %s", err)
	}
}

func (u *User) GetOrCreateExecutionEnvironment(dConfig Atlas.DAppConfig) common.Address {
	execEnvData, err := u.atlas.GetExecutionEnvironment(nil, u.Signer.From, dConfig.To)
	if err != nil {
		log.Fatalf("could not get execution environment data: %s", err)
	}

	if !execEnvData.Exists {
		u.Signer.Value = new(big.Int).Set(common.Big0)
		tx, err := u.atlas.CreateExecutionEnvironment(u.Signer, dConfig)
		if err != nil {
			log.Fatalf("could not create execution environment: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
		if err != nil {
			log.Fatalf("could not wait for execution environment creation transaction to be mined: %s", err)
		}
	}

	return execEnvData.ExecutionEnvironment
}

func (u *User) BuildUserOperation(swapIntent SwapIntentController.SwapIntent) Atlas.UserOperation {
	abi, err := SwapIntentController.SwapIntentControllerMetaData.GetAbi()
	if err != nil {
		log.Fatalf("could not get SwapIntentController abi: %s", err)
	}

	userOpData, err := abi.Pack("swap", swapIntent)
	if err != nil {
		log.Fatalf("could not pack swap intent: %s", err)
	}

	currentBlock, err := u.ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("could not get current block number: %s", err)
	}

	op, err := u.txBuilder.BuildUserOperation(
		nil,
		u.Signer.From,
		u.addresses["dAppController"],
		big.NewInt(1000000000000),
		common.Big0,
		big.NewInt(int64(currentBlock)+100),
		userOpData,
	)
	if err != nil {
		log.Fatalf("could not build user operation: %s", err)
	}

	userOp := Atlas.UserOperation{
		To:   op.To,
		Call: Atlas.UserCall(op.Call),
	}

	userOpPayload, err := u.atlas.GetUserOperationPayload(nil, userOp)
	if err != nil {
		log.Fatalf("could not get user operation payload: %s", err)
	}

	signature, err := crypto.Sign(userOpPayload[:], u.privateKey)
	if err != nil {
		log.Fatalf("could not sign user operation payload: %s", err)
	}

	signature[64] += 27 // Transform V from 0/1 to 27/28
	userOp.Signature = signature

	return userOp
}

func (u *User) Metacall(dConfig Atlas.DAppConfig, userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation, verification Atlas.DAppOperation) {
	signer := u.Signer
	signer.Value = common.Big0
	signer.GasLimit = 5000000

	tx, err := u.atlas.Metacall(
		signer,
		dConfig,
		userOperation,
		solverOperations,
		verification,
	)
	if err != nil {
		log.Fatalf("could not call metacall: %s", err)
	}

	receipt, err := bind.WaitMined(context.Background(), u.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for metacall transaction to be mined: %s", err)
	}

	if receipt.Status == 1 {
		log.Println("<< Metacall successful >>")
	} else {
		log.Println("<< Metacall reverted >>")
	}
}
