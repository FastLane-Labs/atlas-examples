package entities

import (
	"context"
	"log"
	"math/big"

	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ETH_ADDRESS             = common.HexToAddress("0x0")
	WETH_ADDRESS            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI_ADDRESS             = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	UniswapV3Router_ADDRESS = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
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
