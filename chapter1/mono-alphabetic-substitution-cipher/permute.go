package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: expected 1 string argument.")
		os.Exit(1)
	}

	s := os.Args[1]

	permute([]byte(s), []byte{}, func(permutation []byte) {
		fmt.Printf("%s\n", permutation)
	})
}

func makePrefix(prefix []byte, e byte) []byte {
	newPrefix := make([]byte, len(prefix))
	copy(newPrefix, prefix)
	return append(newPrefix, e)
}

func permute(rest []byte, prefix []byte, visitor func([]byte)) {
	if len(rest) == 0 {
		visitor(prefix)
		return
	}

	newPrefix := makePrefix(prefix, rest[0])
	permute(rest[1:], newPrefix, visitor)

	for i := 1; i < len(rest); i++ {
		newPrefix := makePrefix(prefix, rest[i])
		newRest := make([]byte, len(rest)-1)
		copy(newRest, rest[0:i])
		copy(newRest[i:], rest[i+1:])
		permute(newRest, newPrefix, visitor)
	}
}
