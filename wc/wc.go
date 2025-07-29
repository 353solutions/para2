package main

import (
	"fmt"
	"io"
	"os"
)

// $ go run wc.go < road.txt
//	8  51 253

// 8 lines
// 51 words (one or more non-whitespace)
// 253 bytes
func main() {
	var wc WordCount

	if _, err := io.Copy(&wc, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(wc)

}

type WordCount struct{}
