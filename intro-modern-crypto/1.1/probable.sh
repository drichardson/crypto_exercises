#uubin/bash

paste \
    <(cat english_text_stats.txt |sort -k 2 -n -r) \
    <(tr -d '\n'|tr -d ' '|go run ../../tools/frequency_distribution.go |sort -k 2 -n -r) 
