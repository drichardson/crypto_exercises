package main

import (
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var keyHex *string = flag.String("key", "", "AES key as a hex string.")
var ciphertextHex *string = flag.String("ciphertext", "", "Ciphertext to decrypt.")

func main() {
	flag.Parse()

	if *keyHex == "" {
		fmt.Println("Missing key")
		os.Exit(1)
	}

	if *ciphertextHex == "" {
		fmt.Println("Missing ciphertext")
		os.Exit(1)
	}

	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Println("Error decoding hex key.", err)
		os.Exit(1)
	}

	ciphertext, err := hex.DecodeString(*ciphertextHex)
	if err != nil {
		fmt.Println("Error decoding hex ciphertext.", err)
		os.Exit(1)
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating AES cipher.", err)
		os.Exit(1)
	}

	plaintext := make([]byte, cipher.BlockSize())
	cipher.Decrypt(plaintext, ciphertext)
	fmt.Println("Decrypted Hex: ", hex.EncodeToString(plaintext))
}
