default: help

.PHONY: all
all: bin

.PHONY: clean
clean:
	rm -fr $(CURDIR)/bin

.PHONY: server
server: ## Start server
	@sh ./script/server

.PHONY: test
test: ## Run tests
	@sh ./script/test

.PHONY: help
help:
	@echo "Valid targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
