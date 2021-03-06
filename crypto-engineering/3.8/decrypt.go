package main

import (
	"crypto/aes"
	"fmt"
	"os"
)

var ciphertext []byte = []byte{
	0x53, 0x9b, 0x33, 0x3b, 0x39, 0x70, 0x6d, 0x14, 0x90, 0x28, 0xcf, 0xe1,
	0xd9, 0xd4, 0xa4, 0x07,
}

var key []byte = []byte{
	0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
}

func main() {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating AES cipher.", err)
		os.Exit(1)
	}

	plaintext := make([]byte, cipher.BlockSize())
	cipher.Decrypt(plaintext, ciphertext)
	os.Stdout.Write(plaintext)
}
