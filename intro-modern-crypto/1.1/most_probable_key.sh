#!/bin/bash

./probable.sh |cut -f1,3| go run ../../tools/transpose.go | go run ../../tools/transpose.go |sort | go run ../../tools/transpose.go |tr -d ' '
