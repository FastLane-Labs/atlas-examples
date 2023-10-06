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

	// it is better that each entity will do its own setup. Less parameters when you call entities.NewXXXXX
	// it will also be clearer which values are needed for each entity.
	// will be great if you have a different setup for each entity (.env and config.json can be common).
	app := setup()

	shutdownChan := make(chan struct{})
	// swapIntentOperationSubmitChan used by user to submit swapIntentOperations to bundler
	swapIntentOperationSubmitChan := make(chan *entities.SwapIntentOperation)
	// swapIntentOperationBroadcastChan used by bundler to announce swapIntentOperations to solver
	swapIntentOperationBroadcastChan := make(chan *entities.SwapIntentOperation)
	// solverOperationSubmitChan used by solver to submit solverOperation back to bundler
	solverOperationSubmitChan := make(chan *Atlas.SolverOperation)
	// governanceSignatureChan - ??????
	governanceSignatureChan := make(chan []byte)

	// Q: what is the governance entity? It is not clear why governance is needed and when?
	// Q: is this an online component (service) that the dapp has to provide?
	// Launch the governance goroutine
	governance := entities.NewGovernance(app.PrivateKeys["governance"], governanceSignatureChan, shutdownChan)

	// Launch the backend (bundler) goroutine
	// Q: the backend is sending to governance and governance is sending back to backend - how can they do it on the same channel? looks wrong
	entities.NewBackend(app.PrivateKeys["bundler"], app.EthClient, app.Atlas, app.DAppController, app.TxBuilder, governance.GetGovernanceAddress(),
		swapIntentOperationSubmitChan, swapIntentOperationBroadcastChan, solverOperationSubmitChan, governanceSignatureChan, shutdownChan)

	// Launch the solver goroutine
	entities.NewSolver(app.PrivateKeys["solver"], app.EthClient, app.Atlas, app.DAppController, app.TxBuilder, app.Weth, app.Dai, app.Addresses,
		app.UniswapV3Router, swapIntentOperationBroadcastChan, solverOperationSubmitChan, shutdownChan)

	user := entities.NewUser(app.PrivateKeys["user"], app.EthClient, app.Atlas, app.DAppController, app.TxBuilder, app.Weth, app.Addresses,
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

	// Questions:
	// entities:
	//    backend - is this the bundler that gets the intent, distribute to solvers, collect solutio and create the metaTransactions? you can it backend and not bundler as there can be other configurations?
	//    governance - what is this entity? is it part of the dApp or part of Atlas?
	//    user - represents a user application or a real user + dApp ui?
	//
	// smart contracts - can you make the 9 smart contract to their "owner" - dApp, Atlas, system?
	// looks like simulator is not in use
	// what is txBuilder? provided by Atlas or dApp? if by atlas, why a separate contract?
	// why there is a need for SimpleRFQSolver?
}
