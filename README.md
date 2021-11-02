# Success Rate Tool

A simple Go implementation of the [Storj success rate tool](https://github.com/ReneSmeekes/storj_success_rate) by ReneSmeekes.

On average, a speedup of ~7x was observed, for example parsing a log file of 1.1 GiB went from 22 seconds down to 3.

## Install

Either build from source or [download from releases](https://github.com/red-coracle/storj_success_rate/releases)

### Build

```bash
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

You can also read input from stdin, e.g. preprocessing with grep
```bash
grep '2021-10' file.log | go run .
```
or exporting docker logs
```bash
docker logs storagenode | ./successrate
```
