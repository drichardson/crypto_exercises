#!/bin/bash

go build find_pre_image.go

fiveTimes() {
    local digest=$1
    shift
    echo "DIGEST: $digest"
    for i in $(seq 1 5); do
        echo Run $i
        time ./find_pre_image -div 3 -digest "$digest"
    done
}

fiveTimes A9
fiveTimes 3D4B
fiveTimes 3A7F27
fiveTimes C3C0357C

