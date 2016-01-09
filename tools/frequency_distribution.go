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
		freq[string(c)] = freq[string(c)] + 1
	}

	var total int
	for _, v := range freq {
		total += v
	}

	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
