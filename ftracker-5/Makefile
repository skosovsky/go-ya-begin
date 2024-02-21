.SILENT:
APP=go-first-floor-sprint-five

.PHONY: help
help: Makefile ## Show this help
	@echo
	@echo "Choose a command run in "$(APP)":"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: build
build: ## Build an application
	@echo "Building ${APP} ..."
	mkdir -p build
	go build -o build/${APP} main.go

run: ## Run an application
	@echo "Starting ${APP} ..."
	go run main.go

test: ## Run an application
	@echo "Testing ${APP} ..."
	go test

bench: ## Run an application
	@echo "Benchmarking ${APP} ..."
	go test -bench=. .

clean: ## Clean a garbage
	@echo "Cleaning"
	go clean
	rm -rf build

lint: ## Check a code by golangci-lint
	@echo "Linter checking..."
	golangci-lint run -c golangci.yml
