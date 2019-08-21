#!/bin/bash
set -e
OUTFILE=/usr/local/bin/pikolo
go build -o $OUTFILE *.go

chmod +x $OUTFILE