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
	if len(*cipherKey) != 26 {
		fmt.Println("ERROR: cipher key must contain 26 characters but contains", len(*cipherKey))
		os.Exit(1)
	}

	key := []byte(*cipherKey)

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
