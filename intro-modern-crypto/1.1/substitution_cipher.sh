#!/bin/bash
# A mono-alphabetic substitution cipher.
# Usage: substitution_cipher.sh <KEY> <e|d>

KEY=$1
shift
OP=$1

usage() {
    echo "Usage: cipher.sh <KEY> <e|d>"
}

if [ "${#KEY}" != 26 ]; then
    echo "key must be 26 characters"
    usage
    exit 1
fi

DUPLICATES=$(echo -n "$KEY"|fold -w1|sort|uniq -d)

if [ "${#DUPLICATES}" != 0 ]; then
    echo "key contains duplicate characters:"
    echo "$DUPLICATES"
    usage
    exit 1
fi

alphabet=abcdefghijklmnopqrstuvwxyz

case $OP in
    e)
        tr $alphabet "$KEY"
        ;;
    d)
        tr "$KEY" $alphabet
        ;;
    *)
        echo "Invalid operation. Expected e or d."
        usage
        exit 1
        ;;
esac
