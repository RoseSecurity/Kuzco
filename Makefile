BINARY_NAME=kuzco
VERSION=v1
GO=go

default: help

help: ## List Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: build

fmt: ## Format Go files
	gofumpt -w .

build: ## Build Kuzco
	env $(if $(GOOS),GOOS=$(GOOS)) $(if $(GOARCH),GOARCH=$(GOARCH)) $(GO) build -o build/$(BINARY_NAME) -ldflags "-X 'github.com/RoseSecurity/kuzco/cmd.Version=local'" main.go

install: ## Install dependencies
	$(GO) install ./...@latest

clean: ## Clean up build artifacts
	$(GO) clean
	rm ./build/$(BINARY_NAME)

run: build ## Run Kuzco
	./build/$(BINARY_NAME)

docs: build ## Generate documentation
	./build/$(BINARY_NAME) docs

.PHONY: all build install clean run fmt help
