#!/bin/sh

apt-get update && apt-get install -y musl-tools

ln -s /usr/bin/musl-gcc /usr/local/bin/musl-gcc

apt-get clean && rm -rf /var/lib/apt/lists/*

go mod tidy

CC=/usr/local/bin/musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' -tags musl -o /app-run /app/cmd/api/main.go
