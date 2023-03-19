include .bingo/Variables.mk
.DEFAULT_GOAL := test

setup:
	@go install github.com/bwplotka/bingo@latest
	@bingo get
	@$(BINGO) get -l bingo
	@$(BINGO) get -l golangci-lint

test:
	@$(GOLANGCI_LINT) run ./...
	go test ./...

fix:
	@$(GOLANGCI_LINT) run --fix ./...
