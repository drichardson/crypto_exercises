package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var cipherKey *string = flag.String("key", "abcdefghijklmnopqrstuvwxyz", "The key maps position 1 to A, 2 to B, etc. Thus abc... means a=A, b=B, c=C, etc.")

func main() {
	flag.Parse()

	key := []byte(*cipherKey)

	// Make sure key is in key space (one of each in the set A-Z)
	if len(key) != 26 {
		fmt.Println("ERROR: cipher key must contain 26 characters but contains", len(key))
		os.Exit(1)
	}
	keyCharsSeen := make(map[byte]bool)
	for _, c := range key {
		if c < 'a' || c > 'z' {
			fmt.Println("ERROR: Invalid character", string(c), "in key.")
			os.Exit(1)
		}
		_, seen := keyCharsSeen[c]
		if !seen {
			keyCharsSeen[c] = true
		} else {
			fmt.Println("ERROR: Invalid key. Duplicate character", string(c))
			os.Exit(1)
		}
	}

	// Decode
	reader := bufio.NewReader(os.Stdin)
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
			// cipher text can only be A-Z. Pass through everything else.
			fmt.Print(string(c))
			continue
		}

		pos := c - 'A'
		fmt.Print(string(key[pos]))
	}
}
