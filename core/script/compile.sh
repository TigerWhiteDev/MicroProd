#!/bin/bash
export GOPATH=$(pwd)/core/
cd core/src/
go build .
mv src ../../rinyx