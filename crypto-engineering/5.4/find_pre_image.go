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
	"runtime/trace"
)

var digestHexStr = flag.String("digest", "A9", "Find a pre-image that hashes to this digest value. Must be between 1 and 6 bytes.")
var cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
var traceFilename = flag.String("trace", "", "Write trace to file.")
var div = flag.Uint("div", 0, "Divide 64-bit collision search space into 2^div buckets each with 2^(64-div) items. Must be between 1 and 63")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalln("Error creating file for profile.", err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatalln("Error starting CPU profile.", err)
		}
		// defer adds to list of things to execute when surrounding
		// function (not lexical block) returns.
		defer pprof.StopCPUProfile()
	}

	if *traceFilename != "" {
		f, err := os.Create(*traceFilename)
		if err != nil {
			log.Fatalln("Error creating file for trace.", err)
		}
		err = trace.Start(f)
		if err != nil {
			log.Fatalln("Error starting trace.", err)
		}
		defer trace.Stop()
	}

	digest, err := hex.DecodeString(*digestHexStr)
	if err != nil {
		log.Fatalln("Error parsing digest.", err)
	}

	if n := len(digest) * 8; n < 8 || n > 48 || n%8 != 0 {
		log.Fatalln("Invalid value for n.")
	}

	if d := *div; d > 64 {
		log.Fatalln("Invalid value for div.")
	}

	//pre := findPreImage(digest)
	pre := findPreImageConcurrent(digest, *div)
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

func findPreImageConcurrent(digest []byte, div uint) []byte {
	// Divide the search 64 bit search space findPreImageInRange can use
	// into 2^div buckets, each with 2^(64-div) items.
	// 2^div * 2^(64-div) = 2^(div+64-div) = 2^64.
	buckets := uint64(1) << div
	bucketSize := uint64(1) << (64 - div)
	result := make(chan []byte)

	for i := uint64(0); i < buckets; i++ {
		i := i
		go func() {
			start := i * bucketSize
			end := (i+1)*bucketSize - 1
			findPreImageInRange(digest, start, end, result)
		}()
	}

	for i := uint64(0); i < buckets; i++ {
		preimage := <-result
		if preimage != nil {
			return preimage
		}
	}

	return nil
}

func findPreImageInRange(digest []byte, start, end uint64, out chan []byte) {
	digestLen := len(digest)
	var dataBuf [8]byte
	data := dataBuf[:]

	for i := start; ; i++ {
		uint64ToBytes(i+start, &dataBuf)
		dataDigestArray := sha512.Sum512(data)
		dataDigest := dataDigestArray[:digestLen]

		if bytes.Equal(dataDigest, digest) {
			out <- data[:]
			return
		}

		if i == end {
			break
		}
	}

	out <- nil
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
