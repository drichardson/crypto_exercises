#!/bin/bash

set -e

freqdist=$(mktemp)
cat ciphertext |go run frequency_distribution.go > "$freqdist"
paste <(cat "$freqdist"|cut -f 1) <(cat "$freqdist"|cut -f 3|go run english_statistical_guess.go)
