#!/bin/bash

set -e

go build frequency_distribution.go
go build transpose.go

#paste -d " " <(cat english_text_stats.txt|sort -n -k 2 -r) <(cat ciphertext | ./frequency_distribution|tail -n +3|sort -k 2 -n -r) |./transpose|./transpose

paste \
    <(cat english_text_stats.txt |sort -k 2 -n -r|cut -f 1) \
    <(cat ciphertext |./frequency_distribution|sort -k 2 -n -r|cut -f 1)  \
    | ./transpose | tr -d ' '
