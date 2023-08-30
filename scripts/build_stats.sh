#!/bin/bash

set -e

src=$(pwd)
bin="$src/bin"

if [ -d $bin ]; then rm -rf $bin && mkdir -p $bin; fi

GOBIN=$bin GO111MODULE=on go build -o $bin -v ./cmd/stats

echo "build done."
