.PHONY: clean help test
.DEFAULT: help

gen: ## go generate ./...
	go generate ./...

test: ## Run go vet, and test the whole repo
	go vet ./...
	go test ./...

clean: gen ## tidy modules, delete the bin folder, go generate
	go mod tidy

help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@printf "\033[36m%-30s\033[0m %s\n" "(target)" "Build a target binary: $(targets)"
