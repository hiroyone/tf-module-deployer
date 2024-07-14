# Define the name of the binary
BINARY_NAME=bin/tf-module-deployer

# List of packages to test
PKGS=$(shell go list ./... | grep -v /vendor/)

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Format the code
fmt:
	go fmt ./...

# Run go vet for static code analysis
vet:
	go vet ./...

# Run lint checks using golangci-lint
lint:
	golangci-lint run

# Run all tests
test:
	go test -v $(PKGS)

# Clean up the binary and other generated files
clean:
	rm -f $(BINARY_NAME)

# Install necessary tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run the precommit checks (fmt, vet, lint, test)
precommit: fmt vet lint test

# Run the application
run:
	./$(BINARY_NAME) --help

.PHONY: all build fmt vet lint test clean install-tools precommit run