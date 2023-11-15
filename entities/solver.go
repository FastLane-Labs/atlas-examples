package entities

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/AtlasVerification"
	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"
	"github.com/FastLane-Labs/atlas-examples/contracts/Permit2"
	"github.com/FastLane-Labs/atlas-examples/contracts/SimpleRFQSolver"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"
	"github.com/FastLane-Labs/atlas-examples/contracts/UniswapUniversalRouter"
	"github.com/FastLane-Labs/atlas-examples/contracts/WETH9"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	atlEthBalanceMin = big.NewInt(1e17)
)

type Solver struct {
	signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	ethClient *ethclient.Client

	atlas             *Atlas.Atlas
	atlasVerification *AtlasVerification.AtlasVerification
	dappController    *SwapIntentController.SwapIntentController
	txBuilder         *TxBuilder.TxBuilder

	weth *WETH9.WETH9
	uni  *ERC20.ERC20

	permit2                *Permit2.Permit2
	uniswapUniversalRouter *UniswapUniversalRouter.UniswapUniversalRouter

	addresses map[string]common.Address

	// The channel where solvers receive new intent to fulfill
	newSwapIntentOperationChan <-chan *SwapIntentOperation

	// The channel where solvers send their solutions to intents
	solverOperationSubmitChan chan<- *Atlas.SolverOperation

	shutdownChan <-chan struct{}

	log *log.Logger
}

func NewSolver(pk string, ethClient *ethclient.Client, chainId int64, atlas *Atlas.Atlas, atlasVerification *AtlasVerification.AtlasVerification, dappController *SwapIntentController.SwapIntentController,
	txBuilder *TxBuilder.TxBuilder, weth *WETH9.WETH9, uni *ERC20.ERC20, addresses map[string]common.Address, permit2 *Permit2.Permit2, uniswapUniversalRouter *UniswapUniversalRouter.UniswapUniversalRouter,
	newSwapIntentOperationChan chan *SwapIntentOperation, solverOperationSubmitChan chan *Atlas.SolverOperation, shutdownChan chan struct{}) *Solver {
	logger := log.New(os.Stdout, "[SOLVER]\t", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)

	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		logger.Fatalf("could not load solver's private key: %s", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		logger.Fatalf("could not initialize solver's signer: %s", err)
	}

	s := &Solver{
		signer:                     signer,
		privateKey:                 privateKey,
		ethClient:                  ethClient,
		atlas:                      atlas,
		atlasVerification:          atlasVerification,
		dappController:             dappController,
		txBuilder:                  txBuilder,
		weth:                       weth,
		uni:                        uni,
		permit2:                    permit2,
		uniswapUniversalRouter:     uniswapUniversalRouter,
		addresses:                  addresses,
		newSwapIntentOperationChan: newSwapIntentOperationChan,
		solverOperationSubmitChan:  solverOperationSubmitChan,
		shutdownChan:               shutdownChan,
		log:                        logger,
	}
	go s.run()
	return s
}

func (s *Solver) run() {
	for {
		select {
		case swapIntentOperation := <-s.newSwapIntentOperationChan:
			s.log.Println("Received a new swap intent")
			s.getUniForSolverContract(swapIntentOperation.SwapIntent.AmountUserBuys)

			dConfig, err := s.dappController.GetDAppConfig(nil, SwapIntentController.UserOperation(*swapIntentOperation.UserOperation))
			if err != nil {
				s.log.Fatalf("could not get dApp config: %s", err)
			}

			atlEthBalance, err := s.atlas.BalanceOf(nil, s.signer.From)
			if err != nil {
				s.log.Fatalf("could not get solver's atlEth balance: %s", err)
			}

			depositNeeded := new(big.Int).Sub(atlEthBalanceMin, atlEthBalance)
			if depositNeeded.Cmp(common.Big0) > 0 {
				s.log.Println("Getting AtlEth")
				s.signer.Value = depositNeeded
				tx, err := s.atlas.Deposit(s.signer)
				if err != nil {
					s.log.Fatalf("could not deposit ATL-ETH for solver: %s", err)
				}

				_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
				if err != nil {
					s.log.Fatalf("could not wait for deposit transaction to be mined: %s", err)
				}
				s.log.Printf("Got AtlEth: %s", tx.Hash().Hex())
			}

			solverOperation := s.buildSolverOperation(
				Atlas.DAppConfig(dConfig),
				*swapIntentOperation.SwapIntent,
				*swapIntentOperation.UserOperation,
				swapIntentOperation.ExecutionEnvironment,
				new(big.Int).Div(swapIntentOperation.SwapIntent.AmountUserSells, big.NewInt(100)), // Bid 1% of the amount the user sells
			)

			// Submit the solver operation to the backend
			s.log.Println("Submits operation to the backend")
			s.solverOperationSubmitChan <- &solverOperation

		case <-s.shutdownChan:
			return
		}
	}
}

func (s *Solver) getUniForSolverContract(amount *big.Int) {
	solverContractUniBalance, err := s.uni.BalanceOf(nil, s.addresses["solverContract"])
	if err != nil {
		s.log.Fatalf("could not get solver's UNI balance: %s", err)
	}

	uniNeeded := new(big.Int).Sub(amount, solverContractUniBalance)
	if uniNeeded.Cmp(common.Big0) > 0 {
		// Solver wraps some ETH if needed
		solverWethBalance, err := s.weth.BalanceOf(nil, s.signer.From)
		if err != nil {
			s.log.Fatalf("could not get solver's WETH balance: %s", err)
		}

		wethNeeded := new(big.Int).Sub(big.NewInt(2*1e17), solverWethBalance)
		if wethNeeded.Cmp(common.Big0) > 0 {
			s.log.Println("Wrapping ETH")
			s.signer.Value = wethNeeded
			tx, err := s.weth.Deposit(s.signer)
			if err != nil {
				s.log.Fatalf("could not get WETH for solver: %s", err)
			}

			_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
			if err != nil {
				s.log.Fatalf("could not wait for deposit transaction to be mined: %s", err)
			}
			s.log.Printf("Wrapped ETH: %s", tx.Hash().Hex())
		}

		allowance, err := s.weth.Allowance(nil, s.signer.From, Permit2_ADDRESS)
		if err != nil {
			s.log.Fatalf("could not get solver's WETH allowance: %s", err)
		}

		if allowance.Cmp(common.Big0) == 0 {
			s.log.Println("Approving WETH for Permit2")
			tx, err := s.weth.Approve(
				s.signer,
				Permit2_ADDRESS,
				MaxUint256,
			)
			if err != nil {
				s.log.Fatalf("could not approve WETH for solver: %s", err)
			}

			_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
			if err != nil {
				s.log.Fatalf("could not wait for WETH approval transaction to be mined: %s", err)
			}
			s.log.Printf("Approved WETH: %s", tx.Hash().Hex())
		}

		permit2Allowance, err := s.permit2.Allowance(nil, s.signer.From, WETH_ADDRESS, UniswapUniversalRouter_ADDRESS)
		if err != nil {
			s.log.Fatalf("could not get solver's Permit2 allowance: %s", err)
		}

		if permit2Allowance.Amount.Cmp(common.Big0) == 0 {
			s.log.Println("Approving Permit2")
			tx, err := s.permit2.Approve(s.signer, WETH_ADDRESS, UniswapUniversalRouter_ADDRESS, MaxUint160, MaxUint48)
			if err != nil {
				s.log.Fatalf("could not approve Permit2 for solver: %s", err)
			}

			_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
			if err != nil {
				s.log.Fatalf("could not wait for Permit2 approval transaction to be mined: %s", err)
			}
			s.log.Printf("Approved Permit2: %s", tx.Hash().Hex())
		}

		path := common.LeftPadBytes(UNI_ADDRESS.Bytes(), 20)
		path = append(path, common.LeftPadBytes(big.NewInt(3000).Bytes(), 3)...)
		path = append(path, common.LeftPadBytes(WETH_ADDRESS.Bytes(), 20)...)

		commandInputs, err := v3SwapExactOut.Pack(
			s.addresses["solverContract"],
			uniNeeded,
			big.NewInt(2*1e17),
			path,
			true,
		)
		if err != nil {
			s.log.Fatalf("could not pack command inputs: %s", err)
		}

		s.log.Println("Swapping ETH for UNI")
		s.signer.Value = common.Big0
		tx, err := s.uniswapUniversalRouter.Execute(
			s.signer,
			common.Hex2Bytes("01"), // V3_SWAP_EXACT_OUT
			[][]byte{commandInputs},
		)
		if err != nil {
			s.log.Fatalf("could not get UNI for solver: %s", err)
		}

		_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
		if err != nil {
			s.log.Fatalf("could not wait for execute transaction to be mined: %s", err)
		}
		s.log.Printf("Swapped ETH for UNI: %s", tx.Hash().Hex())
	}
}

func (s *Solver) buildSolverOperation(dConfig Atlas.DAppConfig, swapIntent SwapIntentController.SwapIntent, userOperation Atlas.UserOperation, executionEnvironment common.Address, bidAmount *big.Int) Atlas.SolverOperation {
	abi, err := SimpleRFQSolver.SimpleRFQSolverMetaData.GetAbi()
	if err != nil {
		s.log.Fatalf("could not get solver's ABI: %s", err)
	}

	solverOpData, err := abi.Pack("fulfillRFQ", swapIntent, executionEnvironment)
	if err != nil {
		s.log.Fatalf("could not pack solver operation data: %s", err)
	}

	op, err := s.txBuilder.BuildSolverOperation(
		nil,
		TxBuilder.UserOperation(userOperation),
		solverOpData,
		s.signer.From,
		s.addresses["solverContract"],
		bidAmount,
	)
	if err != nil {
		s.log.Fatalf("could not build solver operation: %s", err)
	}

	solverOp := Atlas.SolverOperation(op)
	solverOpPayload, err := s.atlasVerification.GetSolverPayload(nil, AtlasVerification.SolverOperation(solverOp))
	if err != nil {
		s.log.Fatalf("could not get solver operation payload: %s", err)
	}

	signature, err := crypto.Sign(solverOpPayload[:], s.privateKey)
	if err != nil {
		s.log.Fatalf("could not sign solver operation payload: %s", err)
	}

	signature[64] += 27 // Transform V from 0/1 to 27/28
	solverOp.Signature = signature

	return solverOp
}
