include .bingo/Variables.mk

setup:
	@go install github.com/bwplotka/bingo@latest
	@bingo get -l github.com/bwplotka/bingo@v0.8.0
