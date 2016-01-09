package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var patternLength *int = flag.Int("length", 2, "Scan for patterns of this length")

func main() {
	flag.Parse()
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input.", err)
		os.Exit(1)
	}

	freq := make(map[string]int)

	for i, stop := 0, len(in)-*patternLength; i < stop; i++ {
		pattern := in[i : i+*patternLength]
		freq[string(pattern)] += 1
	}

	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
