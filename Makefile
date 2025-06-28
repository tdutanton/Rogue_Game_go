GOFUMPT := $(shell which gofumpt)
GOLINT := $(shell which golint)
SAVE_FILES_DIR := internal/adapters/secondary/storage
LEADERBOARD_FILE_DIR := internal/adapters/secondary/storage

SAVE_FILE_NAME := dungeon_save.json
LEADERBOARD_FILE_NAME := leaderboard.json

all: run
.PHONY: all

run:
	@go run cmd/main.go
.PHONY: run

run_debug:
	@CONFIG_PATH="./configs/debug.yaml" go run cmd/main.go
.PHONY: run_debug

run_prod:
	@CONFIG_PATH="./configs/prod.yaml" go run cmd/main.go
.PHONY: run_prod

clean:
	@rm -rf $(SAVE_FILES_DIR)/$(SAVE_FILE_NAME)
	@rm -rf $(LEADERBOARD_FILE_DIR)/$(LEADERBOARD_FILE_NAME)
.PHONY: clean

fmt:
	@go fmt ./...
.PHONY: fmt

lint: ensure-golint
	@golint ./...
.PHONY: lint

vet:
	@go vet ./...
.PHONY: vet

std_linters: fmt lint vet
.PHONY: std_linters

fumpt: ensure-fumpt std_linters
	@gofumpt -l -w .
.PHONY: fumpt

dvi:
	@echo "To see the documentation open http://localhost:8080 in your browser"
	@pkgsite
.PHONY: dvi

ensure-fumpt:
ifeq (, $(GOFUMPT))
	@echo "Gofumpt is not installed, installing..."
	@go install mvdan.cc/gofumpt@latest
	@echo "Gofumpt installed"
endif
.PHONY: ensure-fumpt

ensure-golint:
ifeq (, $(GOLINT))
	@echo "Golint is not installed, installing..."
	@go install golang.org/x/lint/golint@latest
	@echo "Golint installed"
endif
.PHONY: ensure-golint

ensure-pkgsite:
ifeq (, $(PKGSITE))
	@echo "pkgsite is not installed, installing..."
	@go install golang.org/x/pkgsite/cmd/pkgsite@latest
	@echo "pkgsite installed"
endif
.PHONY: ensure-pkgsite