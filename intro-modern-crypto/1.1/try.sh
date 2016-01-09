#!/bin/bash

go build permute.go

PERMUTE="inshr"
# etainshrdlucwmfygpobvkxjzq

./permute $PERMUTE | while read -r perm; do
    echo "perm: $perm"
    #cat ciphertext |tr FQWGLOVHBPJIZRMECYKASDX "${perm}shrdlucwmfygpobvkx"
    cat ciphertext |tr FQWGLOVHBPJIZRMECYKASDX "eta${perm}dlucwmfygpobvkxjzq"

done
