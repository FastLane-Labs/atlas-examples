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

func (u *User) ApproveErc20Atlas(beneficiary common.Address, amount *big.Int) {
	erc20, err := ERC20.NewERC20(beneficiary, u.ethClient)
	if err != nil {
		log.Fatalf("could not load ERC20 contract: %s", err)
	}

	tx, err := erc20.Approve(u.Signer, u.addresses["atlas"], amount)
	if err != nil {
		log.Fatalf("could not approve ERC20: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for ERC20 approval transaction to be mined: %s", err)
	}
}

func (u *User) GetOrCreateExecutionEnvironment(userOperation Atlas.UserOperation, dConfig Atlas.DAppConfig) common.Address {
	execEnv, err := u.atlas.GetExecutionEnvironment(nil, userOperation, dConfig.To)
	if err != nil {
		log.Fatalf("could not get execution environment: %s", err)
	}

	bytecode, err := u.ethClient.CodeAt(context.Background(), execEnv, nil)
	if err != nil {
		log.Fatalf("could not get execution environment's bytecode: %s", err)
	}

	if len(bytecode) == 0 {
		tx, err := u.atlas.CreateExecutionEnvironment(u.Signer, dConfig)
		if err != nil {
			log.Fatalf("could not create execution environment: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
		if err != nil {
			log.Fatalf("could not wait for execution environment creation transaction to be mined: %s", err)
		}
	}

	return execEnv
}

func (u *User) BuildUserOperation(swapIntent SwapIntentController.SwapIntent) Atlas.UserOperation {
	abi, err := SwapIntentController.SwapIntentControllerMetaData.GetAbi()
	if err != nil {
		log.Fatalf("could not get SwapIntentController abi: %s", err)
	}

	userOpData, err := abi.Methods["swap"].Inputs.Pack(swapIntent)
	if err != nil {
		log.Fatalf("could not pack swap intent: %s", err)
	}
	userOpData = append(abi.Methods["swap"].ID, userOpData...)

	op, err := u.txBuilder.BuildUserOperation(
		nil,
		u.Signer.From,
		u.addresses["dappController"],
		big.NewInt(100000000000),
		common.Big0,
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

	userOp.Signature = signature

	return userOp
}

func (u *User) Metacall(dConfig Atlas.DAppConfig, userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation, verification Atlas.Verification) {
	signer := u.Signer
	signer.Value = common.Big0

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

	_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for metacall transaction to be mined: %s", err)
	}
}
