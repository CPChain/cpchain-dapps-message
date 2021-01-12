
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build

compile:
	@docker run --rm -v $(ROOT_DIR):/go/src/bitbucket.org/cpchain/chain/message -it cpchain2018/abigen abigen --sol ./message/message.sol --pkg message --out ./message/message.go

test:
	@go test

build:
	@mkdir -p build
	@go build -o build/main cmd/main.go
