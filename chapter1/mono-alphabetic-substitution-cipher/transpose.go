// transpose
// a1 a2 a3
// b1 b2 b3
// into
// a1 b1
// a2 b2
// a3 b3

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Row struct {
	columns []string
}

func main() {
	flag.Parse()

	var rows []Row
	var maxCol int

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("ERROR: failed to read input line.", err)
				os.Exit(1)
			}
			break
		}
		columns := strings.Fields(line)
		if l := len(columns); l > maxCol {
			maxCol = l
		}
		rows = append(rows, Row{columns})
	}

	fmt.Println("max col", maxCol)
	for i := 0; i < maxCol; i++ {
		for j, row := range rows {
			if j > 0 {
				fmt.Print(" ")
			}
			if i >= len(row.columns) {
				fmt.Print("?")
			} else {
				fmt.Print(row.columns[i])
			}
		}
		fmt.Println()
	}
}
