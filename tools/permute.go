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

	permute([]byte(s), func(permutation []byte) {
		fmt.Printf("%s\n", permutation)
	})
}

func makePrefix(prefix []byte, e byte) []byte {
	newPrefix := make([]byte, len(prefix))
	copy(newPrefix, prefix)
	return append(newPrefix, e)
}

func permute(rest []byte, visitor func([]byte)) {
	switch len(rest) {
	case 0:
		break
	case 1:
		visitor(rest)
	case 2:
		visitor(rest)
		t := rest[0]
		rest[0] = rest[1]
		rest[1] = t
		visitor(rest)
	default:
		permuteThreeOrMore(rest, []byte{}, visitor)
	}
}

func permuteThreeOrMore(rest []byte, prefix []byte, visitor func([]byte)) {
	if len(rest) == 0 {
		visitor(prefix)
		return
	}

	newPrefix := makePrefix(prefix, rest[0])
	permuteThreeOrMore(rest[1:], newPrefix, visitor)

	for i := 1; i < len(rest); i++ {
		newPrefix := makePrefix(prefix, rest[i])
		newRest := make([]byte, len(rest)-1)
		copy(newRest, rest[0:i])
		copy(newRest[i:], rest[i+1:])
		permuteThreeOrMore(newRest, newPrefix, visitor)
	}
}
