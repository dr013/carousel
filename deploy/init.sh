#!/usr/bin/env bash

go get -u github.com/golang/dep/...
dep init
go get github.com/revel/revel
go get github.com/revel/cmd/revel
go get github.com/cbonello/revel-csrf
dep ensure -update
