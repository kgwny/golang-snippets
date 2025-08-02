# Makefile
.PHONY: tidy download build run clean test lint fmt help

APPNAME := myapp
BUILD_DIR := bin

## go mod tidy
tidy:
	go mod tidy

## download go modules
download:
	go mod download

## build the app
build:
	go build -o $(BUILD_DIR)/$(APPNAME) ./...

## run the app
run:
	go run main.go

## remove built files
clean:
	rm -rf $(BUILD_DIR)

## run tests
test:
	go test -v ./...

## run linters (go vet and staticcheck)
lint:
	go vet ./...
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "warning: staticcheck not found"; \
	fi

## format code
fmt:
	go fmt ./...

## show help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'
