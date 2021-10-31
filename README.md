# Success Rate Tool

A simple Go implementation of the [Storj success rate tool](https://github.com/ReneSmeekes/storj_success_rate) by ReneSmeekes.

On average, a speedup of ~7x was observed e.g. parsing a log file of 1.1 GiB went from 22 seconds down to 3.

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

Multiple paths may be specified to get combined results.
