// Given list of frequencies one per line, print a list of letters
// based on English text letter frequency.
// For example:
// INPUT
// 8.2
// OUTPUT
// a i ...
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type letterFrequency struct {
	letter    string
	frequency float64
}

// Average Letter Frequency of english language text. Figure 1.3 from Introduction to Modern
// Cryptogoraphy, 2nd Edition.
var averageLetterFreq = []letterFrequency{
	{"a", 8.2},
	{"b", 1.5},
	{"c", 2.8},
	{"d", 4.3},
	{"e", 12.7},
	{"f", 2.2},
	{"g", 2.0},
	{"h", 6.1},
	{"i", 7.0},
	{"j", 0.2},
	{"k", 0.8},
	{"l", 4.0},
	{"m", 2.4},
	{"n", 6.7},
	{"o", 1.5},
	{"p", 1.9},
	{"q", 0.1},
	{"r", 6.0},
	{"s", 6.3},
	{"t", 9.1},
	{"u", 2.8},
	{"v", 1.0},
	{"w", 2.4},
	{"x", 0.2},
	{"y", 2.0},
	{"z", 0.1},
}

type frequencySorter struct {
	frequencies []letterFrequency
	by          func(a, b *letterFrequency) bool
}

func (fs *frequencySorter) Len() int {
	return len(fs.frequencies)
}

func (fs *frequencySorter) Less(i, j int) bool {
	return fs.by(&fs.frequencies[i], &fs.frequencies[j])
}

func (fs *frequencySorter) Swap(i, j int) {
	tmp := fs.frequencies[i]
	fs.frequencies[i] = fs.frequencies[j]
	fs.frequencies[j] = tmp
}

func main() {
	/*
		byFrequency := &frequencySorter{
			frequencies: averageLetterFreq,
			by: func(a, b *letterFrequency) bool {
				return a.frequency > b.frequency
			},
		}

		sort.Sort(byFrequency)

		fmt.Println("By Frequency")
		for _, v := range averageLetterFreq {
			fmt.Printf("%v: %.2f\n", v.letter, v.frequency)
		}
	*/

	for {
		var f float64
		_, err := fmt.Scanln(&f)
		if err != nil {
			if err != io.EOF {
				fmt.Println("ERROR: couldn't parse frequency.", err)
				os.Exit(1)
			}
			break
		}
		byDistanceFromFrequency := &frequencySorter{
			frequencies: averageLetterFreq,
			by: func(a, b *letterFrequency) bool {
				return math.Abs(a.frequency-f) < math.Abs(b.frequency-f)
			},
		}

		sort.Sort(byDistanceFromFrequency)

		for _, v := range averageLetterFreq {
			fmt.Printf("%s ", v.letter)
		}
		fmt.Println()
	}
}
