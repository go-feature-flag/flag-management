GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor



all: help
## Build:
build: build-api ## Build all the binaries and put the output in out/bin/

create-out-dir:
	mkdir -p out/bin

build-server: create-out-dir ## Build the migration cli in out/bin/
	CGO_ENABLED=0 GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/flag-management-api .

clean: ## Remove build related file
	-rm -fr ./bin ./out ./release
	-rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

## Dev:
swagger: ## Build swagger documentation
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseInternal --markdownFiles server/docs --output server/docs

setup-env:
	docker stop goff || true
	docker rm goff || true
	docker run --name goff -e POSTGRES_DB=gofeatureflag -e POSTGRES_PASSWORD=my-secret-pw -p 5432:5432 -e POSTGRES_USER=goff-user -d postgres
	sleep 2
	migrate -source "file://database_migration" -database "postgres://goff-user:my-secret-pw@localhost:5432/gofeatureflag?sslmode=disable" up

## Test:
test: test-server
test-server: ## Run the tests of the project
	$(GOTEST) -v -race ./...

## Coverage:
coverage: coverage-server ## Run all the coverage on your project
coverage-server: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -tags=docker -coverprofile=coverage.cov ./...

## Lint:
lint: lint-server ## Run all the linters on your project
lint-server: ## Use golintci-lint on your project
	mkdir -p ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest # Install linters
	./bin/golangci-lint run --timeout=5m --timeout=5m ./... --enable-only=gci --fix # Run linters

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
