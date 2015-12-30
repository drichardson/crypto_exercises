package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	freq := make(map[string]int)
	for {
		c, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file.", err)
				os.Exit(1)
			}
			break
		}
		if c < 'A' || c > 'Z' {
			// cipher text can only be A-Z. Ignore everything else.
			continue
		}
		freq[string(c)] = freq[string(c)] + 1
	}

	var keys []string
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var total int
	for _, v := range freq {
		total += v
	}

	fmt.Println("Total Characters: ", total)
	fmt.Println("Character, Count, %")
	for _, key := range keys {
		percentage := 100.0 * float64(freq[key]) / float64(total)
		percentageDisplay := strconv.FormatFloat(percentage, 'f', 1, 64)
		fmt.Println(key, freq[key], percentageDisplay)
	}
}
