#!/usr/bin/env bash

mkdir results &> /dev/null
ANALYSIS="TRUE" go run cmd/battlesnake/main.go
