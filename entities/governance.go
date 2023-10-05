package entities

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Governance struct {
	privateKey *ecdsa.PrivateKey

	// The channel where governance signature is requested/returned
	governanceSignatureChan chan []byte

	shutdownChan <-chan struct{}

	log *log.Logger
}

func NewGovernance(pk string, governanceSignatureChan chan []byte, shutdownChan chan struct{}) *Governance {
	logger := log.New(os.Stdout, "[GOVERNANCE]\t", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)

	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(pk))
	if err != nil {
		logger.Fatalf("could not load governance's private key: %s", err)
	}

	g := &Governance{
		privateKey:              privateKey,
		governanceSignatureChan: governanceSignatureChan,
		shutdownChan:            shutdownChan,
		log:                     logger,
	}
	go g.run()
	return g
}

func (g *Governance) GetGovernanceAddress() common.Address {
	return crypto.PubkeyToAddress(g.privateKey.PublicKey)
}

func (g *Governance) run() {
	for {
		select {
		case payloadToSign := <-g.governanceSignatureChan:
			g.log.Println("Received payload to sign")
			signature, err := crypto.Sign(payloadToSign, g.privateKey)
			if err != nil {
				g.log.Fatalf("could not sign payload: %s", err)
			}
			signature[64] += 27 // Transform V from 0/1 to 27/28

			// Return the signed payload
			g.log.Println("Sends signed payload to backend")
			g.governanceSignatureChan <- signature

		case <-g.shutdownChan:
			return
		}
	}
}
