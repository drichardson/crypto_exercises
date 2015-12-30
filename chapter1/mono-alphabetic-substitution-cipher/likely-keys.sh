#!/bin/bash

set -e

./likely-character-substitutions.sh |go run transpose.go |tr -d ' '
