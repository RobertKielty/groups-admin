# Go module settings
MOD_DIR := go
BIN := groups-admin
OUT := $(MOD_DIR)/$(BIN)
GO := go

# Use local caches to avoid polluting global dirs (absolute paths required)
export GOMODCACHE := $(CURDIR)/$(MOD_DIR)/.gocache/mod
export GOCACHE := $(CURDIR)/$(MOD_DIR)/.gocache/build

.PHONY: build run vet fmt tidy test clean help .cache-dirs

.cache-dirs:
	@mkdir -p $(GOMODCACHE) $(GOCACHE)

help:
	@echo "Common targets: build, run, vet, fmt, tidy, test, clean"
	@echo "Example run: make run ARGS=\"--baseUrl <url> --srcEmail <email> --srcPass <pass> --cmd srcUserSubs\""

build: .cache-dirs vet ## Build the CLI binary
	cd $(MOD_DIR) && $(GO) build -o $(BIN) .
	@ls -lh $(OUT)

run: build ## Run the CLI with ARGS="..."
	cd $(MOD_DIR) && ./$(BIN) $(ARGS)

vet: .cache-dirs ## Static analysis
	cd $(MOD_DIR) && $(GO) vet ./...

fmt: ## Format source
	cd $(MOD_DIR) && gofmt -s -w .

tidy: ## Tidy dependencies
	cd $(MOD_DIR) && $(GO) mod tidy

test: .cache-dirs ## Run unit tests
	cd $(MOD_DIR) && $(GO) test ./...

clean: ## Remove built binary
	rm -f $(OUT)
