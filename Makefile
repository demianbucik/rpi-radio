WORKDIR := $(shell pwd)

builder:
	docker build -f Dockerfile.builder -t radio-builder .

build:
	[ -d artifacts ] || mkdir artifacts
	docker run \
		-v $(WORKDIR)/radio:/app/go \
		-v $(WORKDIR)/artifacts:/artifacts \
		radio-builder bash -c \
		'CGO_ENABLED=1 CGO_CFLAGS=-w CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /artifacts'

help:
	@echo "help"

.DEFAULT_GOAL := help
