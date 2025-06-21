# Sqlite Agent

[English](./README.md) | [中文](./README_ZH.md)

Add api calls to sqlite

![image](/doc/img/sqlite_agent.jpg)

## build

https://github.com/mattn/go-sqlite3?tab=readme-ov-file#cross-compile

### linux

`CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"`

### mac(ARM)

`CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=arm64 GOOS=darwin CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"`

### win

`CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOARCH=amd64 GOOS=windows CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"`

## run

`./sqlite-agent --config=config.yaml`

## help

``` bash
./sqlite-agent --help   
Usage of ./sqlite-agent:
  -config string
        Path to the configuration file
  -dev
        Run in development mode
```

## config

Please refer to example.yaml and check the comments to understand what configuration means. It is simple 😉

## QA

- You must modify the api_auth in the configuration file to prevent abuse of the interface!!!
- In release, I will provide the standard win/mac/linux version, but I cannot test the effectiveness of the win platform!!!
