package entities

import (
	"context"
	"log"
	"math/big"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ETH_ADDRESS                    = common.HexToAddress("0x0")
	WETH_ADDRESS                   = common.HexToAddress("0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14")
	UNI_ADDRESS                    = common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	UniswapUniversalRouter_ADDRESS = common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD")
)

func approve(tokenAddress common.Address, beneficiary common.Address, amount *big.Int, ethClient *ethclient.Client, signer *bind.TransactOpts) {
	erc20, err := ERC20.NewERC20(tokenAddress, ethClient)
	if err != nil {
		log.Fatalf("could not load ERC20 contract: %s", err)
	}

	signer.Value = new(big.Int).Set(common.Big0)
	tx, err := erc20.Approve(signer, beneficiary, amount)
	if err != nil {
		log.Fatalf("could not approve ERC20 for user: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for ERC20 approval transaction to be mined: %s", err)
	}
}

func printBalances(swapIntentOp *SwapIntentOperation, solverOp *Atlas.SolverOperation, ethClient *ethclient.Client) {
	// Contracts
	tokenUserBuys, err := ERC20.NewERC20(swapIntentOp.SwapIntent.TokenUserBuys, ethClient)
	if err != nil {
		log.Fatalf("could not load TokenUserBuys contract: %s", err)
	}

	tokenUserSells, err := ERC20.NewERC20(swapIntentOp.SwapIntent.TokenUserSells, ethClient)
	if err != nil {
		log.Fatalf("could not load TokenUserSells contract: %s", err)
	}

	// User
	tokenUserBuysBalance, err := tokenUserBuys.BalanceOf(nil, swapIntentOp.UserOperation.From)
	if err != nil {
		log.Fatalf("could not get user's TokenUserBuys balance: %s", err)
	}

	tokenUserSellsBalance, err := tokenUserSells.BalanceOf(nil, swapIntentOp.UserOperation.From)
	if err != nil {
		log.Fatalf("could not get user's TokenUserSells balance: %s", err)
	}

	log.Printf("\t\tUser balances: %s=%s, %s=%s", swapIntentOp.SwapIntent.TokenUserBuys.Hex(), tokenUserBuysBalance.String(), swapIntentOp.SwapIntent.TokenUserSells.Hex(), tokenUserSellsBalance.String())

	// Solver
	solverUserBuysBalance, err := tokenUserBuys.BalanceOf(nil, solverOp.Solver)
	if err != nil {
		log.Fatalf("could not get solvers's TokenUserBuys balance: %s", err)
	}

	solverUserSellsBalance, err := tokenUserSells.BalanceOf(nil, solverOp.Solver)
	if err != nil {
		log.Fatalf("could not get solver's TokenUserSells balance: %s", err)
	}

	log.Printf("\t\tSolver balances: %s=%s, %s=%s", swapIntentOp.SwapIntent.TokenUserBuys.Hex(), solverUserBuysBalance.String(), swapIntentOp.SwapIntent.TokenUserSells.Hex(), solverUserSellsBalance.String())
}
