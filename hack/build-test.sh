#!/bin/bash
set -e
OUTFILE=/tmp/pikolo-test
go build -o $OUTFILE *.go

chmod +x $OUTFILE