package main

import (
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"os"
)

var nBits *int = flag.Int("n", 8, "Number of bits. Must be multiple of 8 between 8 and 48 inclusive.")

func main() {
	flag.Parse()

	// Hash input
	h, err := NewSha512n(*nBits)
	if err != nil {
		fmt.Println("Error creating sha512n.", err)
		os.Exit(1)
	}
	item1, item2, digest := birthdayAttack(h)
	if item1 != nil {
		fmt.Println("Found collision:")
		fmt.Println("item1: ", hex.EncodeToString(item1))
		fmt.Println("item2: ", hex.EncodeToString(item2))
		fmt.Println("digest: ", hex.EncodeToString(digest))
	} else {
		fmt.Println("No collision found.")
		os.Exit(1)
	}
}

func uint64ToBytes(u uint64) []byte {
	return []byte{
		byte(u >> 56 & 255),
		byte(u >> 48 & 255),
		byte(u >> 40 & 255),
		byte(u >> 32 & 255),
		byte(u >> 24 & 255),
		byte(u >> 16 & 255),
		byte(u >> 8 & 255),
		byte(u & 255),
	}
}

func byte8ToUInt64(b [8]byte) uint64 {
	b0 := uint64(b[0]) << 56
	b1 := uint64(b[1]) << 48
	b2 := uint64(b[2]) << 40
	b3 := uint64(b[3]) << 32
	b4 := uint64(b[4]) << 24
	b5 := uint64(b[5]) << 16
	b6 := uint64(b[6]) << 8
	b7 := uint64(b[7]) << 0

	return b0 | b1 | b2 | b3 | b4 | b5 | b6 | b7
}

// Perform birthday attack on 8, 16, 24, 32, 40, or 48 bit SHA-512-n digest.
func birthdayAttack(h hash.Hash) (item1, item2, digestOut []byte) {
	// i is 64 bits. Since 2^64 > 2^48 (the largest SHA-512-n digest)
	// at least one collision will be found due to the pigeon hole principle.
	// There's a ~50% chance of finding such a collision in the first 2^(n/2)
	// tries.

	tries := make(map[string][]byte)

	// Do 0 out of the loop, since 0 is used as our stop point.
	i := uint64(0)
	data := uint64ToBytes(i)
	h.Reset()
	h.Write(data)
	digest := h.Sum(nil)
	tries[string(digest)] = data

	for i++; i != 0; i++ {
		data := uint64ToBytes(i)
		h.Reset()
		h.Write(data)
		digest := h.Sum(nil)
		digestStr := string(digest)
		if prev, ok := tries[digestStr]; ok {
			// found collision
			item1 = data
			item2 = prev
			digestOut = digest
			return
		}
		tries[digestStr] = data
	}

	return nil, nil, nil
}

type sha512n struct {
	n      int
	sha512 hash.Hash
}

// SHA-512 truncated to first n bits, where 8 <= n <= 48 and n is a
// multiple of 8.
func NewSha512n(n int) (hash.Hash, error) {
	if n%8 != 0 || n < 8 || n > 48 {
		return nil, fmt.Errorf("n must be multiple of 8 between 8 and 48")
	}
	return &sha512n{n, sha512.New()}, nil
}

func (h *sha512n) BlockSize() int {
	return h.sha512.Size()
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h *sha512n) Sum(b []byte) []byte {
	sha512Sum := h.sha512.Sum(nil)
	sha512NSum := sha512Sum[:h.Size()]
	return append(b, sha512NSum...)
}

// Reset resets the Hash to its initial state.
func (h *sha512n) Reset() {
	h.sha512.Reset()
}

// Size returns the number of bytes Sum will return.
func (h *sha512n) Size() int {
	return h.n / 8
}

func (h *sha512n) Write(p []byte) (n int, err error) {
	return h.sha512.Write(p)
}
