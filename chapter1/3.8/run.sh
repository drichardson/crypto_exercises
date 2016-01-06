#!/bin/bash

make

echo Go decrypt
go run ./decrypt.go | xxd -p

echo libtom decrypt
./decrypt-libtom | xxd -p
