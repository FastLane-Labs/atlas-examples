package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"
	"github.com/FastLane-Labs/atlas-examples/contracts/UniswapV3Router"
	"github.com/FastLane-Labs/atlas-examples/contracts/WETH9"
	"github.com/FastLane-Labs/atlas-examples/entities"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

type Config struct {
	AtlasAddress          common.Address `json:"atlas"`
	SimulatorAddress      common.Address `json:"simulator"`
	DAppControllerAddress common.Address `json:"dAppController"`
	SolverContractAddress common.Address `json:"solverContract"`
	TxBuilderAddress      common.Address `json:"txBuilder"`
}

type App struct {
	Atlas          *Atlas.Atlas
	DAppController *SwapIntentController.SwapIntentController
	TxBuilder      *TxBuilder.TxBuilder

	Weth            *WETH9.WETH9
	Dai             *ERC20.ERC20
	UniswapV3Router *UniswapV3Router.UniswapV3Router

	EthClient *ethclient.Client

	PrivateKeys map[string]string
	Addresses   map[string]common.Address
}

func setup() *App {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	app := &App{
		PrivateKeys: make(map[string]string),
		Addresses:   make(map[string]common.Address),
	}

	// Load config
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("could not load config file: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("could not decode config file: %s", err)
	}

	app.Addresses["atlas"] = config.AtlasAddress
	app.Addresses["simulator"] = config.SimulatorAddress
	app.Addresses["dAppController"] = config.DAppControllerAddress
	app.Addresses["solverContract"] = config.SolverContractAddress

	// Load .env file
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("could not load .env file: %s", err)
	}

	// Disable certs verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Init eth client
	ethClient, err := ethclient.Dial(env["TESTNET_URL"])
	if err != nil {
		log.Fatalf("could not initialize eth client: %s", err)
	}

	// Dummy request to generate basic auth header
	dr, _ := http.NewRequest("", "", nil)
	dr.SetBasicAuth(env["TESTNET_USERNAME"], env["TESTNET_PASSWORD"])

	// Set basic auth header for all calls
	ethClient.Client().SetHeader("Authorization", dr.Header.Get("Authorization"))

	app.EthClient = ethClient

	// Load Atlas related contracts
	app.Atlas, err = Atlas.NewAtlas(config.AtlasAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize Atlas contract")
	}

	app.DAppController, err = SwapIntentController.NewSwapIntentController(config.DAppControllerAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize DAppController contract")
	}

	app.TxBuilder, err = TxBuilder.NewTxBuilder(config.TxBuilderAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize Tx builder contract")
	}

	// Load tokens contracts
	app.Weth, err = WETH9.NewWETH9(entities.WETH_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize WETH9 contract")
	}

	app.Dai, err = ERC20.NewERC20(entities.DAI_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize DAI contract")
	}

	// Load Uniswap V3 router
	app.UniswapV3Router, err = UniswapV3Router.NewUniswapV3Router(entities.UniswapV3Router_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize Uniswap V3 router contract")
	}

	// Private keys
	app.PrivateKeys["bundler"] = env["BUNDLER_PK"]
	app.PrivateKeys["governance"] = env["GOVERNANCE_PK"]
	app.PrivateKeys["user"] = env["USER_PK"]
	app.PrivateKeys["solver"] = env["SOLVER_PK"]

	return app
}
