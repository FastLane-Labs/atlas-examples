package account

import (
	"crypto/ecdsa"
	"log"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Governance struct {
	Signer     *bind.TransactOpts
	privateKey *ecdsa.PrivateKey

	atlas          *Atlas.Atlas
	dappController *SwapIntentController.SwapIntentController
	txBuilder      *TxBuilder.TxBuilder
}

func NewGovernance(pk string, atlas *Atlas.Atlas, dappController *SwapIntentController.SwapIntentController, txBuilder *TxBuilder.TxBuilder) *Governance {
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
	}
}

func (g *Governance) GetDAppConfig() Atlas.DAppConfig {
	dConfig, err := g.dappController.GetDAppConfig(nil)
	if err != nil {
		log.Fatalf("could not get DApp config: %s", err)
	}
	return Atlas.DAppConfig(dConfig)
}

func (g *Governance) BuildVerification(dConfig Atlas.DAppConfig, userOperation Atlas.UserOperation, solverOperations []Atlas.SolverOperation) Atlas.Verification {
	solverOps := make([]TxBuilder.SolverOperation, len(solverOperations))
	for i, op := range solverOperations {
		bids := make([]TxBuilder.BidData, len(op.Bids))
		for j, bid := range op.Bids {
			bids[j] = TxBuilder.BidData(bid)
		}
		solverOps[i] = TxBuilder.SolverOperation{To: op.To, Call: TxBuilder.SolverCall(op.Call), Signature: op.Signature, Bids: bids}
	}

	v, err := g.txBuilder.BuildVerification(
		nil,
		g.Signer.From,
		TxBuilder.DAppConfig(dConfig),
		TxBuilder.UserOperation{To: userOperation.To, Call: TxBuilder.UserCall(userOperation.Call), Signature: userOperation.Signature},
		solverOps,
	)
	if err != nil {
		log.Fatalf("could not build verification: %s", err)
	}

	verification := Atlas.Verification{
		To:    v.To,
		Proof: Atlas.DAppProof(v.Proof),
	}

	signaturePayload, err := g.atlas.GetVerificationPayload(nil, verification)
	if err != nil {
		log.Fatalf("could not get verification payload: %s", err)
	}

	signature, err := crypto.Sign(signaturePayload[:], g.privateKey)
	if err != nil {
		log.Fatalf("could not sign verification payload: %s", err)
	}

	verification.Signature = signature

	return verification
}
