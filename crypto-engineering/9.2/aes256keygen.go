package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	var key [256 / 8]byte // 256 bit key
	if _, err := rand.Read(key[:]); err != nil {
		log.Println("Error reading random data for key.", err)
	}
	fmt.Println(hex.EncodeToString(key[:]))
}
