# Make does not offer a recursive wildcard function, so here's one:
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

SHELL := /bin/bash
NAME := terraform-drift-check
BUILD_TARGET = build
MAIN_SRC_FILE=main.go
GO := GO111MODULE=on go
REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
ORG := chrismellard
ORG_REPO := $(ORG)/$(NAME)
RELEASE_ORG_REPO := $(ORG_REPO)
ROOT_PACKAGE := github.com/$(ORG_REPO)
GO_VERSION := 1.15
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
CGO_ENABLED = 0

# set dev version unless VERSION is explicitly set via environment
VERSION ?= $(shell echo "$$(git for-each-ref refs/tags/ --count=1 --sort=-version:refname --format='%(refname:short)' 2>/dev/null)-dev+$(REV)" | sed 's/^v//')

# Full build flags used when building binaries. Not used for test compilation/execution.
BUILDFLAGS :=  -ldflags \
  " -X $(ROOT_PACKAGE)/pkg/version.Version=$(VERSION)\
		-X $(ROOT_PACKAGE)/pkg/version.Revision='$(REV)'\
		-X $(ROOT_PACKAGE)/pkg/version.Branch='$(BRANCH)'\
		-X $(ROOT_PACKAGE)/pkg/version.BuildDate='$(BUILD_DATE)'\
		-X $(ROOT_PACKAGE)/pkg/version.GoVersion='$(GO_VERSION)'\
		$(BUILD_TIME_CONFIG_FLAGS)"

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

build: $(GO_DEPENDENCIES) clean
	go mod download
	CGO_ENABLED=$(CGO_ENABLED) $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o build/$(NAME) $(MAIN_SRC_FILE)

install: $(GO_DEPENDENCIES) ## Install the binary
	GOBIN=${GOPATH}/bin $(GO) install $(BUILDFLAGS) $(MAIN_SRC_FILE)

linux: ## Build for Linux
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o build/linux/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/linux/$(NAME)

arm: ## Build for ARM
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=arm $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o build/arm/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/arm/$(NAME)

win: ## Build for Windows
	CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o build/win/$(NAME)-windows-amd64.exe $(MAIN_SRC_FILE)

darwin: ## Build for OSX
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=amd64 $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o build/darwin/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/darwin/$(NAME)

.PHONY: release
release: clean linux

release-all: release linux win darwin

.PHONY: clean
clean: ## Clean the generated artifacts
	rm -rf build release dist
