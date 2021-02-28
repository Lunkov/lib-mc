#!/bin/bash

GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic
gocritic check -enableAll