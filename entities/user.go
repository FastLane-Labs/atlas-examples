package entities

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/AtlasFactory"
	"github.com/FastLane-Labs/atlas-examples/contracts/AtlasVerification"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"
	"github.com/FastLane-Labs/atlas-examples/contracts/WETH9"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	WETH_AMOUNT_TO_SELL  = big.NewInt(1e17)
	UNI_AMOUNT_TO_BUY, _ = new(big.Int).SetString("40000000000000000000", 10)
)

type User struct {
	signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	ethClient *ethclient.Client

	atlas             *Atlas.Atlas
	atlasFactory      *AtlasFactory.AtlasFactory
	atlasVerification *AtlasVerification.AtlasVerification
	dappController    *SwapIntentController.SwapIntentController
	txBuilder         *TxBuilder.TxBuilder

	weth *WETH9.WETH9

	addresses map[string]common.Address

	// The channel where the user will send their swap intents
	swapIntentOperationSubmitChan chan<- *SwapIntentOperation

	log *log.Logger
}

func NewUser(pk string, ethClient *ethclient.Client, chainId int64, atlas *Atlas.Atlas, atlasFactory *AtlasFactory.AtlasFactory, atlasVerification *AtlasVerification.AtlasVerification, dappController *SwapIntentController.SwapIntentController,
	txBuilder *TxBuilder.TxBuilder, weth *WETH9.WETH9, addresses map[string]common.Address, swapIntentOperationSubmitChan chan *SwapIntentOperation) *User {
	logger := log.New(os.Stdout, "[USER]\t", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)

	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		logger.Fatalf("could not load user's private key: %s", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		logger.Fatalf("could not initialize user's signer: %s", err)
	}

	return &User{
		signer:                        signer,
		privateKey:                    privateKey,
		ethClient:                     ethClient,
		atlas:                         atlas,
		atlasFactory:                  atlasFactory,
		atlasVerification:             atlasVerification,
		dappController:                dappController,
		txBuilder:                     txBuilder,
		weth:                          weth,
		addresses:                     addresses,
		swapIntentOperationSubmitChan: swapIntentOperationSubmitChan,
		log:                           logger,
	}
}

func (u *User) StartSwapIntent() {
	swapIntent := SwapIntentController.SwapIntent{
		TokenUserBuys:          UNI_ADDRESS,
		AmountUserBuys:         UNI_AMOUNT_TO_BUY,
		TokenUserSells:         WETH_ADDRESS,
		AmountUserSells:        WETH_AMOUNT_TO_SELL,
		AuctionBaseCurrency:    ETH_ADDRESS,
		SolverMustReimburseGas: false,
		Conditions:             []SwapIntentController.Condition{},
	}

	userOperation := u.buildUserOperation(swapIntent)
	executionEnvironment := u.getOrCreateExecutionEnvironment(userOperation.Dapp)

	u.getWethAndApproveAtlas(WETH_AMOUNT_TO_SELL)

	// Submit the intent to the backend
	u.log.Println("submits swap intent to the backend")
	u.swapIntentOperationSubmitChan <- &SwapIntentOperation{
		SwapIntent:           &swapIntent,
		UserOperation:        &userOperation,
		ExecutionEnvironment: executionEnvironment,
	}
}

func (u *User) getWethAndApproveAtlas(amount *big.Int) {
	// Get some WETH if needed
	wethBalance, err := u.weth.BalanceOf(nil, u.signer.From)
	if err != nil {
		u.log.Fatalf("could not get user's WETH balance: %s", err)
	}

	wethNeeded := new(big.Int).Sub(amount, wethBalance)
	if wethNeeded.Cmp(common.Big0) > 0 {
		u.signer.Value = wethNeeded
		tx, err := u.weth.Deposit(u.signer)
		if err != nil {
			u.log.Fatalf("could not get WETH for user: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
		if err != nil {
			u.log.Fatalf("could not wait for deposit transaction to be mined: %s", err)
		}
	}

	// Approve those WETH for Atlas
	approve(WETH_ADDRESS, u.addresses["atlas"], wethNeeded, u.ethClient, u.signer)
}

func (u *User) getOrCreateExecutionEnvironment(dAppControl common.Address) common.Address {
	execEnvData, err := u.atlasFactory.GetExecutionEnvironment(nil, u.signer.From, dAppControl)
	if err != nil {
		u.log.Fatalf("could not get execution environment data: %s", err)
	}

	if !execEnvData.Exists {
		u.signer.Value = new(big.Int).Set(common.Big0)
		tx, err := u.atlas.CreateExecutionEnvironment(u.signer, dAppControl)
		if err != nil {
			u.log.Fatalf("could not create execution environment: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), u.ethClient, tx)
		if err != nil {
			u.log.Fatalf("could not wait for execution environment creation transaction to be mined: %s", err)
		}
	}

	return execEnvData.ExecutionEnvironment
}

func (u *User) buildUserOperation(swapIntent SwapIntentController.SwapIntent) Atlas.UserOperation {
	abi, err := SwapIntentController.SwapIntentControllerMetaData.GetAbi()
	if err != nil {
		u.log.Fatalf("could not get SwapIntentController abi: %s", err)
	}

	userOpData, err := abi.Pack("swap", swapIntent)
	if err != nil {
		u.log.Fatalf("could not pack swap intent: %s", err)
	}

	currentBlock, err := u.ethClient.BlockNumber(context.Background())
	if err != nil {
		u.log.Fatalf("could not get current block number: %s", err)
	}

	op, err := u.txBuilder.BuildUserOperation(
		nil,
		u.signer.From,
		u.addresses["dAppController"],
		big.NewInt(10*1e9),
		common.Big0,
		big.NewInt(int64(currentBlock)+100),
		userOpData,
	)
	if err != nil {
		u.log.Fatalf("could not build user operation: %s", err)
	}

	userOp := Atlas.UserOperation(op)
	userOpPayload, err := u.atlasVerification.GetUserOperationPayload(nil, AtlasVerification.UserOperation(userOp))
	if err != nil {
		u.log.Fatalf("could not get user operation payload: %s", err)
	}

	signature, err := crypto.Sign(userOpPayload[:], u.privateKey)
	if err != nil {
		u.log.Fatalf("could not sign user operation payload: %s", err)
	}

	signature[64] += 27 // Transform V from 0/1 to 27/28
	userOp.Signature = signature

	return userOp
}
