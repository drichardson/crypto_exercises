#!/bin/bash

set -e

freqdist=$(mktemp)
cat ciphertext |go run frequency_distribution.go|tail -n +3 > "$freqdist"
paste <(cat "$freqdist"|cut -d ' ' -f 1) <(cat "$freqdist"|cut -d ' ' -f 3|go run english_statistical_guess.go)
