package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var digestHexStr = flag.String("digest", "A9", "Find a pre-image that hashes to this digest value. Must be between 1 and 6 bytes.")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalln("Error creating file for profile.", err)
		}
		pprof.StartCPUProfile(f)
		// defer adds to list of things to execute when surrounding
		// function (not lexical block) returns.
		defer pprof.StopCPUProfile()
	}

	digest, err := hex.DecodeString(*digestHexStr)
	if err != nil {
		log.Fatalln("Error parsing digest.", err)
	}

	if n := len(digest) * 8; n < 8 || n > 48 || n%8 != 0 {
		log.Fatalln("Invalid value for n.")
	}

	pre := findPreImage(digest)
	if pre != nil {
		fmt.Println(hex.EncodeToString(pre))
	} else {
		log.Println("No preimage found for ", *digestHexStr)
	}
}

func findPreImage(digest []byte) []byte {
	//data := make([]byte, 8)
	digestLen := len(digest)
	var dataBuf [8]byte
	data := dataBuf[:]

	for i := uint64(0); ; i++ {
		uint64ToBytes(i, &dataBuf)
		dataDigestArray := sha512.Sum512(data)
		dataDigest := dataDigestArray[:digestLen]

		if bytes.Equal(dataDigest, digest) {
			return data[:]
		}

		if i == math.MaxUint64 {
			// In reality, won't ever get here, but shows the intention of the loop.
			break
		}
	}

	return nil
}

func uint64ToBytes(u uint64, out *[8]byte) {
	(*out)[0] = byte(u >> 56 & 255)
	(*out)[1] = byte(u >> 48 & 255)
	(*out)[2] = byte(u >> 40 & 255)
	(*out)[3] = byte(u >> 32 & 255)
	(*out)[4] = byte(u >> 24 & 255)
	(*out)[5] = byte(u >> 16 & 255)
	(*out)[6] = byte(u >> 8 & 255)
	(*out)[7] = byte(u & 255)
}
