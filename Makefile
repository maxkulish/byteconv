
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BIN_DIR = $(ROOT_DIR)/bin
PROJ_NAME = byteconv

help: _help_

_help_:
	@echo make coverage - run test coverage and open html file with results
	@echo make benchmark - run benchmark tests with memory alocation

coverage:
	cd $(ROOT_DIR)
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

benchmark:
	cd $(ROOT_DIR)
	go test -bench . -benchmem

test:
	cd $(ROOT_DIR)
	go test -v .