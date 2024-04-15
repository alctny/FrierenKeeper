#! env /bin/bash

mkdir -p ~/.local/bin
go build -o frierencli -ldflags '-s -w' ./cmd/cli/
go build -o frierentui -ldflags '-s -w' ./cmd/tui/
mv frierencli frierentui ~/.local/bin