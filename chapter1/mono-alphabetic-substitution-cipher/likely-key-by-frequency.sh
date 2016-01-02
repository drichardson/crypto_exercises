#!/bin/bash

set -e

go build frequency_distribution.go
go build transpose.go

paste \
    <(cat english_text_stats.txt |sort -k 2 -n -r|cut -f 1 -d ' ') \
    <(cat ciphertext |./frequency_distribution|tail -n +3|sort -k 2 -n -r|cut -f 1 -d ' ') \
    | \
    ./transpose |./transpose|sort|cut -f 2 -d ' '|./transpose|tr -d ' '|tr [A-Z] [a-z]
