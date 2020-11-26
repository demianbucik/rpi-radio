WORKDIR := $(shell pwd)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_NAME := rpi-radio-$(COMMIT)
REMOTE_ADDR := 192.168.1.139
USER := ubuntu

builder:
	docker build -f Dockerfile.builder -t radio-builder .

build_server:
	mkdir -p artifacts
	docker run --rm \
		-v $(WORKDIR)/radio:/app/go \
		-v $(WORKDIR)/artifacts:/artifacts \
		radio-builder bash -c " \
			CGO_ENABLED=1 CGO_CFLAGS=-w CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /artifacts \
		"

build_client:
	npm --prefix client run build

build: build_server build_client

deploy:
	mkdir -p build/$(BUILD_NAME)/config
	cp scripts/run.sh build/$(BUILD_NAME)
	cp artifacts/radio build/$(BUILD_NAME)
	cp radio/config/prod.env.toml build/$(BUILD_NAME)/config/env.toml
	cp -r client/dist build/$(BUILD_NAME)
	tar --remove-files -czf build/$(BUILD_NAME).tar.gz -C build $(BUILD_NAME)
	scp build/$(BUILD_NAME).tar.gz $(USER)@$(REMOTE_ADDR):/app/radio
	ssh $(USER)@$(REMOTE_ADDR) " \
		cd /app/radio && \
		tar -xzf $(BUILD_NAME).tar.gz && \
		rm $(BUILD_NAME).tar.gz && \
		sudo systemctl stop radio && \
		cp -rT $(BUILD_NAME) rpi-radio-current && \
		ln -sf rpi-radio-current/* . && \
		sudo systemctl start radio && \
		rm -r $(BUILD_NAME) \
	"
	rm build/$(BUILD_NAME).tar.gz

help:
	@echo "help"

.DEFAULT_GOAL := help
