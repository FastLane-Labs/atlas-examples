ABIS := $(wildcard ./contracts/abi/*.json)

bindings:
	for abi in $(ABIS); do \
		mkdir -p ./contracts/$$(basename $$abi .json); \
		abigen --abi $$abi --pkg $$(basename $$abi .json) --type $$(basename $$abi .json) --out ./contracts/$$(basename $$abi .json)/bindings.go; \
	done
