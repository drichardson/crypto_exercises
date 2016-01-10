#!/bin/bash

set -e

go build ../../tools/pattern_frequency.go

INPUT=$1
shift

if [ ! -f "$INPUT" ]; then
    echo "Missing input file"
    exit 1
fi

patternfreq() {
    ./pattern_frequency -length $1
}

# Find n most frequency patterns of a given length. Only counts
# patterns that appear more than once.
n_most_frequent_patterns() {
    local n=$1
    local length=$2
    patternfreq $length | grep -v "\t1$" | sort -k 2 -n | tail -n $n
}

for x in 2 3 4 5 6 7 8 9 10; do
    tr -d '\n' < "$INPUT" | n_most_frequent_patterns 10 $x
done
