package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	var total int
	for _, v := range freq {
		total += v
	}

	for k, v := range freq {
		percentage := 100.0 * float64(v) / float64(total)
		fmt.Printf("%s\t%d\t%.1f\n", k, v, percentage)
	}
}
