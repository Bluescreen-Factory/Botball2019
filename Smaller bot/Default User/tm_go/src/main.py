#!/bin/bash

GOMAIN="test2.go"

export GOPATH=$HOME/Documents/KISS/go
export PATH=$PATH:/usr/local/go/bin

echo "starting ..."

rm log.txt
go run $GOMAIN 2>&1 |tee log.txt

# echo "done ******" >> log.txt
# echo "my line in logger (journalctl) ******" | logger
# ls "$(pwd)"
# cat "$(pwd)"/log.txt