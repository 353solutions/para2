package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
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

	/*
		data := []byte{'G', 0xF0, 0x9F, 0x98, 0x81}
		r := bytes.NewReader(data)
		for {
			r, _, err := r.ReadRune()
			if errors.Is(err, io.EOF) {
				break
			}

			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Printf("%c", r)
		}
	*/

}

func (wc WordCount) String() string {
	return fmt.Sprintf("%d %d %d", wc.lines, wc.words, wc.bytes)
}

// Write implements io.Writer
func (wc *WordCount) Write(data []byte) (int, error) {
	wc.bytes += len(data)
	for _, b := range data {
		if b == '\n' {
			wc.lines += 1
		}

		if wc.inWord && unicode.IsSpace(rune(b)) {
			wc.inWord = false
		}
		if !wc.inWord && !unicode.IsSpace(rune(b)) {
			wc.inWord = true
			wc.words += 1
		}
	}

	return len(data), nil
}

type WordCount struct {
	bytes int
	words int
	lines int

	inWord bool
}
