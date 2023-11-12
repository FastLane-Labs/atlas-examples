package entities

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/AtlasVerification"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Backend struct {
	// The backend is the bundler in this example
	bundlerSigner     *bind.TransactOpts
	bundlerPrivateKey *ecdsa.PrivateKey

	ethClient *ethclient.Client

	atlas             *Atlas.Atlas
	atlasVerification *AtlasVerification.AtlasVerification
	dappController    *SwapIntentController.SwapIntentController
	txBuilder         *TxBuilder.TxBuilder

	governanceAddress common.Address

	// Indexed by user call hash
	processingSwapIntents map[string]*SwapIntentOperation

	// The channel where users send their swap intent operations
	swapIntentOperationReceiveChan <-chan *SwapIntentOperation

	// The channel where swap intent operations are broadcasted to solvers
	swapIntentOperationBroadcastChan chan<- *SwapIntentOperation

	// The channel where solvers send their solutions to intents
	solverOperationReceiveChan <-chan *Atlas.SolverOperation

	// The channel where governance signature is requested/received
	governanceSignatureChan chan []byte

	shutdownChan <-chan struct{}

	log *log.Logger
}

func NewBackend(pk string, ethClient *ethclient.Client, chainId int64, atlas *Atlas.Atlas, atlasVerification *AtlasVerification.AtlasVerification, dappController *SwapIntentController.SwapIntentController, txBuilder *TxBuilder.TxBuilder,
	governanceAddress common.Address, swapIntentOperationReceiveChan, swapIntentOperationBroadcastChan chan *SwapIntentOperation,
	solverOperationReceiveChan chan *Atlas.SolverOperation, governanceSignatureChan chan []byte, shutdownChan chan struct{}) *Backend {
	logger := log.New(os.Stdout, "[BACKEND]\t", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)

	bundlerPrivateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		logger.Fatalf("could not load user's private key: %s", err)
	}

	bundlerSigner, err := bind.NewKeyedTransactorWithChainID(bundlerPrivateKey, big.NewInt(int64(chainId)))
	if err != nil {
		logger.Fatalf("could not initialize user's signer: %s", err)
	}

	b := &Backend{
		bundlerSigner:                    bundlerSigner,
		bundlerPrivateKey:                bundlerPrivateKey,
		ethClient:                        ethClient,
		atlas:                            atlas,
		atlasVerification:                atlasVerification,
		dappController:                   dappController,
		txBuilder:                        txBuilder,
		governanceAddress:                governanceAddress,
		processingSwapIntents:            make(map[string]*SwapIntentOperation),
		swapIntentOperationReceiveChan:   swapIntentOperationReceiveChan,
		swapIntentOperationBroadcastChan: swapIntentOperationBroadcastChan,
		solverOperationReceiveChan:       solverOperationReceiveChan,
		governanceSignatureChan:          governanceSignatureChan,
		shutdownChan:                     shutdownChan,
		log:                              logger,
	}
	go b.run()
	return b
}

func (b *Backend) run() {
	for {
		select {
		case swapIntentOperation := <-b.swapIntentOperationReceiveChan:
			b.log.Println("Received a new swap intent")

			// Swap intent received from user, save it as processing
			p, err := userOpArg.Pack(&swapIntentOperation.UserOperation)
			if err != nil {
				b.log.Fatalf("could not pack user call: %s", err)
			}
			userOpHash := common.Bytes2Hex(crypto.Keccak256(p))
			b.processingSwapIntents[userOpHash] = swapIntentOperation

			// Broadcast it to solvers
			b.log.Println("Broadcasts swap intent to solvers")
			b.swapIntentOperationBroadcastChan <- swapIntentOperation

		case solverOperation := <-b.solverOperationReceiveChan:
			b.log.Println("Received a solver operation")

			// Solver operation received from solver
			swapIntentOperation, pending := b.processingSwapIntents[common.Bytes2Hex(solverOperation.UserOpHash[:])]
			if !pending {
				// No pending intent for this solver operation, ignore it
				b.log.Println("Solver operation received for non-pending intent")
				continue
			}

			// In production, we would collect more solver operations, until the auction deadline
			// is reached. In this example, we'll just go forward with the first solver operation we get
			b.log.Println("Ends the auction for this intent")
			solverOperations := make([]Atlas.SolverOperation, 1)
			solverOperations[0] = *solverOperation

			// Build dApp operation
			dConfig, err := b.dappController.GetDAppConfig(nil, SwapIntentController.UserOperation(*swapIntentOperation.UserOperation))
			if err != nil {
				b.log.Fatalf("could not get dApp config: %s", err)
			}
			dAppOperation := b.BuildDAppOperation(Atlas.DAppConfig(dConfig), *swapIntentOperation.UserOperation, solverOperations)

			// Generate the payload to be signed by governance
			payloadToSign, err := b.atlasVerification.GetDAppOperationPayload(nil, AtlasVerification.DAppOperation(dAppOperation))
			if err != nil {
				b.log.Fatalf("could not get dApp operation payload: %s", err)
			}

			// Send the payload to governance
			b.log.Println("Sends payload to governance for signature")
			b.governanceSignatureChan <- payloadToSign[:]

			// Wait and get the signature from governance
			dAppOperation.Signature = <-b.governanceSignatureChan
			b.log.Println("Received signed payload from governance")

			// Before
			printBalances(swapIntentOperation, solverOperation, b.ethClient)

			// In this example, the backend is the bundler
			b.log.Println("Calls metacall")
			b.metacall(*swapIntentOperation.UserOperation, solverOperations, dAppOperation)

			// After
			printBalances(swapIntentOperation, solverOperation, b.ethClient)

			// Intent filled, remove it from processing
			delete(b.processingSwapIntents, common.Bytes2Hex(solverOperation.UserOpHash[:]))

		case <-b.shutdownChan:
			return
		}
	}
}

func (b *Backend) metacall(userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation, verification Atlas.DAppOperation) {
	signer := b.bundlerSigner
	signer.Value = common.Big0
	signer.GasLimit = 1000000

	tx, err := b.atlas.Metacall(
		signer,
		userOperation,
		solverOperations,
		verification,
	)
	if err != nil {
		b.log.Fatalf("could not call metacall: %s", err)
	}

	receipt, err := bind.WaitMined(context.Background(), b.ethClient, tx)
	if err != nil {
		b.log.Fatalf("could not wait for metacall transaction to be mined: %s", err)
	}

	var status string
	if receipt.Status == 1 {
		status = "Metacall successful"
	} else {
		status = "Metacall reverted"
	}
	b.log.Printf("%s: %s", status, tx.Hash().Hex())
}

// Builds the dAppOperation object without signature
func (b *Backend) BuildDAppOperation(dConfig Atlas.DAppConfig, userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation) Atlas.DAppOperation {
	solverOps := make([]TxBuilder.SolverOperation, len(solverOperations))
	for i, op := range solverOperations {
		solverOps[i] = TxBuilder.SolverOperation(op)
	}

	v, err := b.txBuilder.BuildDAppOperation(
		nil,
		b.governanceAddress,
		TxBuilder.UserOperation(userOperation),
		solverOps,
	)
	if err != nil {
		b.log.Fatalf("could not build verification: %s", err)
	}

	return Atlas.DAppOperation(v)
}
