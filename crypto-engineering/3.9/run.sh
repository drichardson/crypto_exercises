#!/bin/bash

make

echo Go encrypt
go run ./encrypt.go | xxd -p

echo libtom encrypt
./encrypt-libtom | xxd -p

echo openssl encrypt
./encrypt-openssl | xxd -p
