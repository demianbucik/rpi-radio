FROM golang:1.14 AS builder

RUN dpkg --add-architecture arm64

RUN apt-get update && \
    apt-get install -y \
    gcc-aarch64-linux-gnu \
    libvlc-dev:arm64

    # vlc-plugin-base:arm64 \
    # vlc-plugin-video-output:arm64

WORKDIR /app/go
