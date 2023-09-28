# Atlas examples
A simple example for the swap intent case using Atlas.

## Setup

**Environment**

Create `.env` file, patterned on `.env.example`. Credentials and private keys for the testnet are privately communicated.

**Config file**

`config.json` needs to be up to date for the script to run properly.

**Go bindings**

To generate or regenerate the needed contract bindings, `abigen` needs to be installed: https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings.

Place new or updated ABIs in `contracts/abi` as `json` files, then run:

```
make bindings
```

This will generate each go bindings in their own package, so they can be imported without conflicting with each other.
Example:
- We have `contracts/abi/ERC20.json`.
- We run `make bindings`.
- `contracts/ERC20/bindings.go` has been generated.
- Import it in your go file with `import "github.com/FastLane-Labs/atlas-examples/contracts/ERC20"`.

**Run the script**

```
go run .
```
