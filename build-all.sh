#!/bin/sh
GOOS=windows GOARCH=386 go build -o bitmicro-miner-monitor-windows-i386.exe
GOOS=windows GOARCH=amd64 go build -o bitmicro-miner-monitor-windows-amd64.exe
GOOS=linux GOARCH=386 go build -o bitmicro-miner-monitor-linux-i386
GOOS=linux GOARCH=amd64 go build -o bitmicro-miner-monitor-linux-amd64
GOOS=darwin GOARCH=386 go build -o bitmicro-miner-monitor-macos-i386
GOOS=darwin GOARCH=amd64 go build -o bitmicro-miner-monitor-macos-amd64
