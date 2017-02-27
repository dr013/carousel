#!/usr/bin/env bash

# Put this file to the /etc/profile.d/
export GOPATH=/srv/go
export PATH=$PATH:$(go env GOPATH)/bin
