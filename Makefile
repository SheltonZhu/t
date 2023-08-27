.DEFAULT_GOAL := help
# ==============================================================================
# Variables
GIT_TAG = $(shell git rev-parse --short=8 HEAD)
VERSION ?= $(shell git describe --tags 2>/dev/null)

# ==============================================================================
# Usage
define USAGE_OPTIONS
Options:
endef
export USAGE_OPTIONS

# ==============================================================================
# Targets

## init: 下载依赖工具
.PHONY: init
init:
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1

## test: 测试
.PHONY: test
test:
	go test -coverpkg=./pkg/... -coverprofile=./coverage.out ./pkg/...
	go tool cover -html=./coverage.out -o coverage.html

## gen-mock: 使用 mockgen 生成接口 mock文件
.PHONY: gen-mock
gen-mock:
	go generate -run mockgen ./...

## help: 帮助信息
.PHONY: help
help: Makefile
	@echo "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
