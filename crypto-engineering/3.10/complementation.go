package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// Complementation property of DES means that ^E(k,P) == E(^k,^P).

	key := []byte{1, 2, 3, 4, 5, 6, 7, 0}       // 56-bits +8 parity bits
	plaintext := []byte{1, 2, 3, 4, 5, 6, 7, 8} // 64-bits

	keyComp := complementArray(key)
	plaintextComp := complementArray(plaintext)

	c1 := E(key, plaintext)
	c1Comp := complementArray(c1)
	c2 := E(keyComp, plaintextComp)

	fmt.Println(" E( k, P) = ", hex.EncodeToString(c1))
	fmt.Println("^E( k, P) = ", hex.EncodeToString(c1Comp))
	fmt.Println(" E(^k,^P) = ", hex.EncodeToString(c2))

	if bytes.Compare(c1Comp, c2) == 0 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}

func complementArray(a []byte) []byte {
	r := make([]byte, len(a))
	for i, v := range a {
		r[i] = ^v
	}
	return r
}

func E(key, plaintext []byte) []byte {
	cipher, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR: failed to create DES cipher.", err)
		os.Exit(1)
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.Encrypt(ciphertext, plaintext)
	return ciphertext
}
