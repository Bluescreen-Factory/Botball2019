#!/bin/bash

GOMAIN="$HOME/Documents/KISS/go/src/botball/speed_equalizer/main.go"

export GOPATH=$HOME/Documents/KISS/go
export PATH=$PATH:/usr/local/go/bin

echo "starting ..."

rm log.txt
go run $GOMAIN 2>&1 |tee log.txt
