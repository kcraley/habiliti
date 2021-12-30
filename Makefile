# Variables used to compile the binary
BINARY:=habiliti
MODULE:=github.com/kcraley/habiliti
BUILDDATE=$(shell date "+%s")
GITVERSION=$(shell git describe --tags --always)

LDFLAGS=\
	-X $(module)/internal/version.buildDate=$(BUILDDATE) \
	-X $(module)/internal/version.gitVersion=$(GITVERSION)

.PHONY: deps
deps:
	go get

.PHONY: build
build: deps vet
	go build -ldflags="-w -s $(LDFLAGS)" -o $(BINARY)

.PHONY: vet
vet:
	go vet ./...
