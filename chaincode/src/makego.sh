#!/bin/bash
# Script to compile go modules of the custody asset use case
# Exit on first error, print all commands.
set -ev

rm main
go build main.go data.go invokeCustodian.go invokeBank.go

