package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
)

var count = flag.Uint("count", 10, "Number of numbers to generate.")

// Generate random numbers in the set {0, 1, 2, ..., 191}. Use
// a naive appraoch, generating a random 8-bit value and reducing
// it modulo 192.
func main() {
	flag.Parse()

	for i := uint(0); i < *count; i++ {
		var b [1]byte
		if _, err := rand.Read(b[:]); err != nil {
			log.Fatal(err)
		}
		v := int(b[0] % 192)
		fmt.Println(v)
	}
}
