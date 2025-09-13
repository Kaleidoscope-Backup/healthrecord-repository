PROJECT_NAME := "healthrecord-repository"
PKG := "github.com/karte/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

lint: ## (WIP) Lint the files
	@golint -set_exit_status ${PKG_LIST}

utest: ## Run unit tests
	@go test -short ./resolver/resolver_unit_test/

itest: ## Run integration tests
	@go test -v ./service/service_test/

ftest: ## Run functional tests
	@go test -v ./resolver/resolver_functional_test/

test: ## Run all tests
	@go test -short ${PKG_LIST}

race: dep ## (WIP) Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## (WIP) Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## (WIP) Generate global code coverage report
	./scripts/coverage.sh;

coverhtml: ## (WIP) Generate global code coverage report in HTML
	./scripts/coverage.sh html;

dep: ## Get the dependencies
	@dep ensure
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

run-local: dep ## Run the Go Code Locally Via CLI
	@go run ./cmd/health_record_repository/main.go

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
