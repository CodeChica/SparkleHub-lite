PACKAGES = $(shell go list ./...)

default: help

.PHONY: all
all: bin

bin/server:
	CGO_ENABLED=0 go build -o "bin/server" cmd/server/main.go

bin: bin/server ## Build binary

.PHONY: clean
clean:
	rm -fr $(CURDIR)/bin $(CURDIR)/tmp/**/*.pem

.PHONY: server
server: bin/server ## Start server
	bin/server

.PHONY: test
test: ## Run tests
	go test -v -race $(PACKAGES) $(ARGS)

.PHONY: help
help:
	@echo "Valid targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
