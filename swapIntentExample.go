package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/UniswapV3Router"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	USER            = "user"
	SOLVER_CONTRACT = "solver contract"
)

var (
	WETH_AMOUNT_TO_SELL, _ = new(big.Int).SetString("10000000000000000000", 10)
	DAI_AMOUNT_TO_BUY, _   = new(big.Int).SetString("20000000000000000000", 10)
)

var (
	ETH_ADDRESS             = common.Address{0}
	WETH_ADDRESS            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI_ADDRESS             = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	UniswapV3Router_ADDRESS = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
)

func main() {
	// In this example, the user intents to sell 10 WETH for 20 DAI
	app := setup()

	// To make it work, we first need to fund the user and the solver's contract with some ETH and DAI
	userWethBalance, err := app.Weth.BalanceOf(nil, app.User.Signer.From)
	if err != nil {
		log.Fatalf("could not get user's WETH balance: %s", err)
	}

	wethNeeded := new(big.Int).Sub(WETH_AMOUNT_TO_SELL, userWethBalance)
	if wethNeeded.Cmp(common.Big0) > 0 {
		opts := app.User.Signer
		opts.Value = wethNeeded
		tx, err := app.Weth.Deposit(app.User.Signer)
		if err != nil {
			log.Fatalf("could not get WETH for user: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), app.EthClient, tx)
		if err != nil {
			log.Fatalf("could not wait for deposit transaction to be mined: %s", err)
		}
	}

	solverContractDaiBalance, err := app.Dai.BalanceOf(nil, app.Addresses["solverContract"])
	if err != nil {
		log.Fatalf("could not get solver's DAI balance: %s", err)
	}

	daiNeeded := new(big.Int).Sub(DAI_AMOUNT_TO_BUY, solverContractDaiBalance)
	if daiNeeded.Cmp(common.Big0) > 0 {
		// Solver wraps some ETH if needed
		solverWethBalance, err := app.Weth.BalanceOf(nil, app.Solver.Signer.From)
		if err != nil {
			log.Fatalf("could not get solver's WETH balance: %s", err)
		}

		wethNeeded = new(big.Int).Sub(big.NewInt(1e18), solverWethBalance)
		if wethNeeded.Cmp(common.Big0) > 0 {
			opts := app.Solver.Signer
			opts.Value = wethNeeded
			tx, err := app.Weth.Deposit(app.Solver.Signer)
			if err != nil {
				log.Fatalf("could not get WETH for solver: %s", err)
			}

			_, err = bind.WaitMined(context.Background(), app.EthClient, tx)
			if err != nil {
				log.Fatalf("could not wait for deposit transaction to be mined: %s", err)
			}
		}

		app.Solver.ApproveErc20(WETH_ADDRESS, UniswapV3Router_ADDRESS, big.NewInt(1e18))

		opts := app.Solver.Signer
		opts.Value = common.Big0
		tx, err := app.UniswapV3Router.ExactOutputSingle(
			opts,
			UniswapV3Router.ISwapRouterExactOutputSingleParams{
				TokenIn:           WETH_ADDRESS,
				TokenOut:          DAI_ADDRESS,
				Fee:               big.NewInt(3000),
				Recipient:         app.Addresses["solverContract"],
				Deadline:          big.NewInt(time.Now().Unix() + 10),
				AmountOut:         daiNeeded,
				AmountInMaximum:   big.NewInt(1e18),
				SqrtPriceLimitX96: common.Big0,
			},
		)
		if err != nil {
			log.Fatalf("could not get DAI for solver: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), app.EthClient, tx)
		if err != nil {
			log.Fatalf("could not wait for exactOutputSingle transaction to be mined: %s", err)
		}
	}

	// dApp configuration
	dConfig := app.Governance.GetDAppConfig()

	// The user's swap intent
	swapIntent := SwapIntentController.SwapIntent{
		TokenUserBuys:          DAI_ADDRESS,
		AmountUserBuys:         DAI_AMOUNT_TO_BUY,
		TokenUserSells:         WETH_ADDRESS,
		AmountUserSells:        WETH_AMOUNT_TO_SELL,
		AuctionBaseCurrency:    ETH_ADDRESS,
		SolverMustReimburseGas: false,
		Conditions:             []SwapIntentController.Condition{},
	}

	// Build the user operation from the swap intent
	userOperation := app.User.BuildUserOperation(swapIntent)
	b, err := rlp.EncodeToBytes(userOperation)
	if err != nil {
		log.Fatalf("could not encode user operation: %s", err)
	}
	log.Println("User operation data:", common.Bytes2Hex(b))

	// Create the user's execution environment if it does not exist yet
	executionEnvironment := app.User.GetOrCreateExecutionEnvironment(dConfig)

	// User must approve Atlas
	app.User.ApproveErc20Atlas(WETH_ADDRESS, WETH_AMOUNT_TO_SELL)

	// Solver deposit to escrow contract
	app.Solver.DepositToEscrow(big.NewInt(1e18))

	// Build the solvers operations (only 1 in this example)
	solverOperations := make([]Atlas.SolverOperation, 1)
	solverOperations[0] = app.Solver.BuildSolverOperation(dConfig, swapIntent, userOperation, executionEnvironment, big.NewInt(1e18))

	// Build dApp's verification
	verification := app.Governance.BuildVerification(dConfig, userOperation, solverOperations)

	log.Println("Before metacall")
	logBalances(app, app.User.Signer.From, USER)
	logBalances(app, app.Addresses["solverContract"], SOLVER_CONTRACT)

	// User calls Atlas's metacall
	app.User.Metacall(dConfig, userOperation, solverOperations, verification)

	log.Println("After metacall")
	logBalances(app, app.User.Signer.From, USER)
	logBalances(app, app.Addresses["solverContract"], SOLVER_CONTRACT)
}

// Retrieve onchain and log WETH and DAI balances for account
func logBalances(app *App, account common.Address, name string) {
	wethBalance, err := app.Weth.BalanceOf(nil, account)
	if err != nil {
		log.Fatalf("could not get %s's WETH balance: %s", name, err)
	}
	daiBalance, err := app.Dai.BalanceOf(nil, account)
	if err != nil {
		log.Fatalf("could not get %s's DAI balance: %s", name, err)
	}
	log.Printf("%s's balances: %s WETH - %s DAI", name, wethBalance.String(), daiBalance.String())
}
