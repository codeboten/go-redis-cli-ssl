all: build

.PHONY: build
build:
	go build

.PHONY: deps
deps:
	go mod init