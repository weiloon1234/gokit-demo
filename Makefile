.PHONY: default build-web build-cli build-all build-web-linux build-cli-windows build-all-cross copy-base-entity ent run-server run-server-build run-cli hot gokit-update go-clean-cache clean-build help

# Variables
BIN_DIR := bin
WEB_CMD := ./cmd/web
CLI_CMD := ./cmd/cli
VERBOSE ?= 0
QUIET := $(if $(filter 1, $(VERBOSE)),,@)

# Default target
default: build-all

# Build targets
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

build-web: $(BIN_DIR)
	@$(QUIET) echo "Building web server..."
	@go build -o $(BIN_DIR)/web $(WEB_CMD)

build-cli: $(BIN_DIR)
	@$(QUIET) echo "Building CLI tool..."
	@go build -o $(BIN_DIR)/cli $(CLI_CMD)

build-all: build-web build-cli

build-web-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/web-linux $(WEB_CMD)

build-cli-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/cli-windows.exe $(CLI_CMD)

build-all-cross: build-web-linux build-cli-windows

copy-ent: 
	go run $(CLI_CMD)/main.go copy-ent
	
copy-ent-hook: 
	go run $(CLI_CMD)/main.go copy-ent-hook

ent:
	go generate ./ent

# Run targets
run-server:
	@$(BIN_DIR)/web

run-server-build: build-web
	@$(BIN_DIR)/web

run-cli: build-cli
	@$(BIN_DIR)/cli $(filter-out $@,$(MAKECMDGOALS))

hot:
	air -c .air.toml

go-clean-cache:
	go clean -modcache

gokit-update:
	GOPROXY=direct go get -u github.com/weiloon1234/gokit@latest
	GOPROXY=direct go get -u github.com/weiloon1234/gokit-base-entity@latest

# Clean target
clean-build:
	@rm -rf $(BIN_DIR)

# Help target
help:
	@echo "Available targets:"
	@echo "  build-web         Build the web server executable"
	@echo "  build-cli         Build the CLI tool executable"
	@echo "  build-all         Build both the web server and CLI tool"
	@echo "  build-web-linux   Build the web server for Linux"
	@echo "  build-cli-windows Build the CLI tool for Windows"
	@echo "  build-all-cross   Build for multiple platforms"
	@echo "  copy-ent          Copy base entity to the project"
	@echo "  copy-ent-hook     Copy entity hook to the project"
	@echo "  ent               Generate ent schema"
	@echo "  run-server        Build and run the web server"
	@echo "  run-cli           Build and run the CLI tool
	@echo "  hot-web           Hot reload web server with air"
	@echo "  hot-cli           Hot reload CLI tool with air"
	@echo "  clean-build       Remove all built files"
