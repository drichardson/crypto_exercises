package main

import (
	"crypto/sha512"
	"fmt"
	"os"
)

// xxd -p -r message.hex | xxd -i
var data []byte = []byte{
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x20, 0x20, 0x20,
}

func main() {
	h := sha512.New()
	n, err := h.Write(data)
	if err != nil {
		fmt.Println("Error writing data to hash.", err)
		os.Exit(1)
	}
	if n != len(data) {
		fmt.Println("Expected to write", len(data), "bytes but only wrote", n, "bytes.")
		os.Exit(1)
	}
	digest := h.Sum(nil)
	os.Stdout.Write(digest)
}
