package main

import (
	"os"
	"os/signal"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/entities"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	app := setup()

	shutdownChan := make(chan struct{})
	swapIntentOperationSubmitChan := make(chan *entities.SwapIntentOperation)
	swapIntentOperationBroadcastChan := make(chan *entities.SwapIntentOperation)
	solverOperationSubmitChan := make(chan *Atlas.SolverOperation)
	governanceSignatureChan := make(chan []byte)

	// Launch the governance goroutine
	governance := entities.NewGovernance(app.PrivateKeys["governance"], governanceSignatureChan, shutdownChan)

	// Launch the backend goroutine
	entities.NewBackend(app.PrivateKeys["bundler"], app.EthClient, app.ChainId, app.Atlas, app.AtlasVerification, app.DAppController, app.TxBuilder, governance.GetGovernanceAddress(),
		swapIntentOperationSubmitChan, swapIntentOperationBroadcastChan, solverOperationSubmitChan, governanceSignatureChan, shutdownChan)

	// Launch the solver goroutine
	entities.NewSolver(app.PrivateKeys["solver"], app.EthClient, app.ChainId, app.Atlas, app.AtlasVerification, app.DAppController, app.TxBuilder, app.Weth, app.Uni, app.Addresses,
		app.Permit2, app.UniswapUniversalRouter, swapIntentOperationBroadcastChan, solverOperationSubmitChan, shutdownChan)

	user := entities.NewUser(app.PrivateKeys["user"], app.EthClient, app.ChainId, app.Atlas, app.AtlasFactory, app.AtlasVerification, app.DAppController, app.TxBuilder, app.Weth, app.Addresses,
		swapIntentOperationSubmitChan)

	// User initiate a swap intent
	user.StartSwapIntent()

	// CTRL + C to stop the program
	<-interrupt

	close(shutdownChan)
	close(swapIntentOperationSubmitChan)
	close(swapIntentOperationBroadcastChan)
	close(solverOperationSubmitChan)
	close(governanceSignatureChan)
}
