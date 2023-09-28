package account

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/SimpleRFQSolver"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Solver struct {
	Signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	atlas     *Atlas.Atlas
	txBuilder *TxBuilder.TxBuilder

	addresses map[string]common.Address

	ethClient *ethclient.Client
}

func NewSolver(pk string, atlas *Atlas.Atlas, txBuilder *TxBuilder.TxBuilder, addresses map[string]common.Address, ethClient *ethclient.Client) *Solver {
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		log.Fatalf("could not load solver's private key: %s", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, common.Big1)
	if err != nil {
		log.Fatalf("could not initialize solver's signer: %s", err)
	}

	return &Solver{
		Signer:     signer,
		privateKey: privateKey,
		atlas:      atlas,
		txBuilder:  txBuilder,
		addresses:  addresses,
		ethClient:  ethClient,
	}
}

func (s *Solver) DepositToEscrow(amount *big.Int) {
	opts := s.Signer
	opts.Value = amount

	tx, err := s.atlas.Deposit(opts, opts.From)
	if err != nil {
		log.Fatalf("could not deposit to escrow: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), s.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for execution environment creation transaction to be mined: %s", err)
	}
}

func (s *Solver) BuildSolverOperation(dConfig Atlas.DAppConfig, swapIntent SwapIntentController.SwapIntent, userOperation Atlas.UserOperation, executionEnvironment common.Address, bidAmount *big.Int) Atlas.SolverOperation {
	abi, err := SimpleRFQSolver.SimpleRFQSolverMetaData.GetAbi()
	if err != nil {
		log.Fatalf("could not get solver's ABI: %s", err)
	}

	solverOpData, err := abi.Pack("fulfillRFQ", swapIntent, executionEnvironment)
	if err != nil {
		log.Fatalf("could not pack solver operation data: %s", err)
	}
	solverOpData = append(abi.Methods["fulfillRFQ"].ID, solverOpData...)

	op, err := s.txBuilder.BuildSolverOperation(
		nil,
		TxBuilder.UserOperation{To: userOperation.To, Call: TxBuilder.UserCall(userOperation.Call), Signature: userOperation.Signature},
		TxBuilder.DAppConfig(dConfig),
		solverOpData,
		s.Signer.From,
		s.addresses["solverContract"],
		bidAmount,
	)
	if err != nil {
		log.Fatalf("could not build solver operation: %s", err)
	}

	bids := make([]Atlas.BidData, len(op.Bids))
	for i, bid := range op.Bids {
		bids[i] = Atlas.BidData(bid)
	}

	solverOp := Atlas.SolverOperation{
		To:   op.To,
		Call: Atlas.SolverCall(op.Call),
		Bids: bids,
	}

	solverOpPayload, err := s.atlas.GetSolverPayload(nil, solverOp.Call)
	if err != nil {
		log.Fatalf("could not get solver operation payload: %s", err)
	}

	signature, err := crypto.Sign(solverOpPayload[:], s.privateKey)
	if err != nil {
		log.Fatalf("could not sign solver operation payload: %s", err)
	}

	solverOp.Signature = signature

	return solverOp
}
