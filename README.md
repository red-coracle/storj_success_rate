# Success Rate Tool

A simple Go implementation of the [Storj success rate tool](https://github.com/ReneSmeekes/storj_success_rate) by ReneSmeekes.

## Install

Either build from source or [download from releases](https://github.com/red-coracle/storj_success_rate/releases)

### Build

```
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/successrate-amd64-linux .
# Windows
GOOS=windows GOARCH=amd64 go build -o bin/successrate-amd64-windows.exe .
# OSX
GOOS=darwin GOARCH=amd64 go build -o bin/successrate-amd64-darwin .
```

## Usage

`./successrate /path/to/log/file.log`
