#!/usr/bin/env bash

go get ./...
pushd cmd
go build -o ../build/otfdocgen
popd
