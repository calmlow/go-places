# Prerequisites for this Makefile is:
#
# - Go installed
#

BIN_NAME := go-places
MAIN_GO := ./main.go

# First taget = default target
.PHONY: build install test clean

build: test
	go build -o ./$(BIN_NAME) $(MAIN_GO)

install: build
	@echo "Make sure you have the alias setup or the bind stuff. See readme"

test:
	go test ./...

clean:
	@rm -f ./$(BIN_NAME)
