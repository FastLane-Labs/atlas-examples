package account

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Governance struct {
	Signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	atlas          *Atlas.Atlas
	dappController *SwapIntentController.SwapIntentController
	txBuilder      *TxBuilder.TxBuilder

	ethClient *ethclient.Client
}

func NewGovernance(pk string, atlas *Atlas.Atlas, dappController *SwapIntentController.SwapIntentController, txBuilder *TxBuilder.TxBuilder, ethClient *ethclient.Client) *Governance {
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		log.Fatalf("could not load governance's private key: %s", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, common.Big1)
	if err != nil {
		log.Fatalf("could not initialize governance's signer: %s", err)
	}

	return &Governance{
		Signer:         signer,
		privateKey:     privateKey,
		atlas:          atlas,
		dappController: dappController,
		txBuilder:      txBuilder,
		ethClient:      ethClient,
	}
}

func (g *Governance) GetDAppConfig() Atlas.DAppConfig {
	dConfig, err := g.dappController.GetDAppConfig(nil)
	if err != nil {
		log.Fatalf("could not get DApp config: %s", err)
	}
	return Atlas.DAppConfig(dConfig)
}

func (g *Governance) InitializeGovernance(dappControllerAddress common.Address) {
	tx, err := g.atlas.InitializeGovernance(
		g.Signer,
		dappControllerAddress,
	)
	if err != nil {
		log.Fatalf("could not initialize governance: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), g.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for initialize governance transaction to be mined: %s", err)
	}
}

func (g *Governance) IntegrateDapp(dappControllerAddress common.Address) {
	tx, err := g.atlas.IntegrateDApp(
		g.Signer,
		dappControllerAddress,
		dappControllerAddress,
	)
	if err != nil {
		log.Fatalf("could not integrate dApp: %s", err)
	}

	_, err = bind.WaitMined(context.Background(), g.ethClient, tx)
	if err != nil {
		log.Fatalf("could not wait for integrate dApp transaction to be mined: %s", err)
	}
}

func (g *Governance) BuildVerification(dConfig Atlas.DAppConfig, userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation) Atlas.DAppOperation {
	solverOps := make([]TxBuilder.SolverOperation, len(solverOperations))
	for i, op := range solverOperations {
		bids := make([]TxBuilder.BidData, len(op.Bids))
		for j, bid := range op.Bids {
			bids[j] = TxBuilder.BidData(bid)
		}
		solverOps[i] = TxBuilder.SolverOperation{To: op.To, Call: TxBuilder.SolverCall(op.Call), Signature: op.Signature, Bids: bids}
	}

	v, err := g.txBuilder.BuildDAppOperation(
		nil,
		g.Signer.From,
		TxBuilder.DAppConfig(dConfig),
		TxBuilder.UserOperation{To: userOperation.To, Call: TxBuilder.UserCall(userOperation.Call), Signature: userOperation.Signature},
		solverOps,
	)
	if err != nil {
		log.Fatalf("could not build verification: %s", err)
	}

	verification := Atlas.DAppOperation{
		To:       v.To,
		Approval: Atlas.DAppApproval(v.Approval),
	}

	signaturePayload, err := g.atlas.GetDAppOperationPayload(nil, verification)
	if err != nil {
		log.Fatalf("could not get verification payload: %s", err)
	}

	signature, err := crypto.Sign(signaturePayload[:], g.privateKey)
	if err != nil {
		log.Fatalf("could not sign verification payload: %s", err)
	}

	signature[64] += 27 // Transform V from 0/1 to 27/28
	verification.Signature = signature

	return verification
}
