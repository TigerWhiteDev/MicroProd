#!/bin/bash
if [[ $GOPATH == "" ]]
then
    export GOPATH=$(pwd)/core/
fi
cd core/src/
go build .
mv src ../../rinyx