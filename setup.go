package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/FastLane-Labs/atlas-examples/account"
	"github.com/FastLane-Labs/atlas-examples/contracts/Atlas"
	"github.com/FastLane-Labs/atlas-examples/contracts/ERC20"
	"github.com/FastLane-Labs/atlas-examples/contracts/SwapIntentController"
	"github.com/FastLane-Labs/atlas-examples/contracts/TxBuilder"
	"github.com/FastLane-Labs/atlas-examples/contracts/UniswapV3Router"
	"github.com/FastLane-Labs/atlas-examples/contracts/WETH9"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

type Config struct {
	AtlasAddress          common.Address `json:"atlas"`
	DAppControllerAddress common.Address `json:"dAppController"`
	SolverContractAddress common.Address `json:"solverContract"`
	TxBuilderAddress      common.Address `json:"txBuilder"`
}

type App struct {
	Governance *account.Governance
	User       *account.User
	Solver     *account.Solver

	Weth            *WETH9.WETH9
	Dai             *ERC20.ERC20
	UniswapV3Router *UniswapV3Router.UniswapV3Router

	Addresses map[string]common.Address
}

func setup() *App {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	app := &App{
		Addresses: make(map[string]common.Address),
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

	// Load tokens contracts
	app.Weth, err = WETH9.NewWETH9(WETH_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize WETH9 contract")
	}

	app.Dai, err = ERC20.NewERC20(DAI_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize DAI contract")
	}

	// Load Uniswap V3 router
	app.UniswapV3Router, err = UniswapV3Router.NewUniswapV3Router(UniswapV3Router_ADDRESS, ethClient)
	if err != nil {
		log.Fatal("could not initialize Uniswap V3 router contract")
	}

	// Load Atlas related contracts
	atlas, err := Atlas.NewAtlas(config.AtlasAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize Atlas contract")
	}

	dAppController, err := SwapIntentController.NewSwapIntentController(config.DAppControllerAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize DAppController contract")
	}

	txBuilder, err := TxBuilder.NewTxBuilder(config.TxBuilderAddress, ethClient)
	if err != nil {
		log.Fatal("could not initialize Tx builder contract")
	}

	// Load governance account
	app.Governance = account.NewGovernance(env["GOVERNANCE_PK"], atlas, dAppController, txBuilder)

	// Load user account
	app.User = account.NewUser(env["USER_PK"], atlas, txBuilder, app.Addresses, ethClient)

	// Load solver account
	app.Solver = account.NewSolver(env["SOLVER_PK"], atlas, txBuilder, app.Addresses, ethClient)

	return app
}
