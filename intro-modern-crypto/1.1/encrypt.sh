#!/bin/bash

KEY=$1

if [ "${#KEY}" != 26 ]; then
    echo "key must be 26 characters"
    exit 1
fi

DUPLICATES=$(echo -n "$KEY"|fold -w1|sort|uniq -d)

if [ "${#DUPLICATES}" != 0 ]; then
    echo "key contains duplicate characters:"
    echo "$DUPLICATES"
    exit 1
fi

tr abcdefghijklmnopqrstuvwxyz "$KEY"
