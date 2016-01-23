// A a version of birthday_attack.go that I optimized using runtime/pprof.
package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"os"
	"runtime/pprof"
)

var nBits = flag.Uint("n", 8, "Number of bits. Must be multiple of 8 between 8 and 48 inclusive.")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println("Error creating file for profile.", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		// defer adds to list of things to execute when surrounding
		// function (not lexical block) returns.
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			fmt.Println("Error creating file for memprofile.", err)
			os.Exit(1)
		}
		defer func() {
			pprof.WriteHeapProfile(f)
			f.Close()
		}()
	}

	// Hash input
	h, err := NewSha512n(int(*nBits))
	if err != nil {
		fmt.Println("Error creating sha512n.", err)
		os.Exit(1)
	}
	item1, item2, digest := birthdayAttack(h, *nBits)
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

func uint64ToBytes(u uint64, out []byte) {
	out[0] = byte(u >> 56 & 255)
	out[1] = byte(u >> 48 & 255)
	out[2] = byte(u >> 40 & 255)
	out[3] = byte(u >> 32 & 255)
	out[4] = byte(u >> 24 & 255)
	out[5] = byte(u >> 16 & 255)
	out[6] = byte(u >> 8 & 255)
	out[7] = byte(u & 255)
}

func digestToIndex(digest []byte) uint64 {
	var index uint64
	for _, v := range digest {
		index = (index << 8) | uint64(v)
	}
	return index
}

// Perform birthday attack on 8, 16, 24, 32, 40, or 48 bit SHA-512-n digest.
func birthdayAttack(h hash.Hash, nBits uint) (item1, item2, digestOut []byte) {
	// i is 64 bits. Since 2^64 > 2^48 (the largest SHA-512-n digest)
	// at least one collision will be found due to the pigeon hole principle.
	// There's a ~50% chance of finding such a collision in the first 2^(n/2)
	// tries.

	type try struct {
		data   [8]byte
		digest [6]byte
		set    bool
	}
	//tries := make(map[string][]byte)
	// 2^24 = sqrt(2^48), so having that many entries lets us get to
	// ~50% chance of successful attack. Once we've reached that point,
	// the table will start to degrade like a chained hash table.
	// sqrt(2^n) == 2^(n/2) == 1 << (n/2) == 1 << (n >> 1).
	tableSize := uint64(1 << (nBits >> 1))
	tableSize <<= 2 // make the table times bigger to reduce false index collisions from ~60% to ~10%
	tableIndexMask := tableSize - 1
	tries := make([]try, tableSize)
	fmt.Printf("Table len is %d. tableIndexMask is %x, nBits: %d, tableSize: %d\n", len(tries), tableIndexMask, nBits, tableSize)

	// Loop below doesn't try i == 0. That's okay since we really only need to hit 2^48 different
	// values, and the index is a uint64.

	//var digestBuf [8]byte
	var dataBuf [8]byte
	var indexCollisionCounter int
	digestLen := nBits / 8
	var digest512 [sha512.Size]byte

	stopValue := uint64(1) << nBits

	for i := uint64(1); i < stopValue; i++ {
		data := dataBuf[:]
		uint64ToBytes(i, data)
		digest512 = sha512.Sum512(data)
		digest := digest512[:digestLen]

		tableIndex := digestToIndex(digest) & tableIndexMask
		//fmt.Println("index", i, "tableIndex", tableIndex)
		tableDigest := tries[tableIndex].digest[:digestLen]
		if bytes.Equal(tableDigest, digest) {
			item1 = data
			item2 = tries[tableIndex].data[:]
			digestOut = digest
			// If indexCollisionCounter is significant relative to the size of the search space, you need a bigger table.
			fmt.Println("indexCollisionCounter", indexCollisionCounter, "iterations", i, "percent collisions", float64(indexCollisionCounter)/float64(i))
			return
		}
		if tries[tableIndex].set {
			indexCollisionCounter++
		}
		//fmt.Println("Replacing ", hex.EncodeToString(tries[tableIndex].digest[:]), " with ", hex.EncodeToString(digest))
		copy(tries[tableIndex].digest[:], digest)
		copy(tries[tableIndex].data[:], data)
		tries[tableIndex].set = true
	}

	return nil, nil, nil
}
