package tokenizer_test

import (
	"fmt"
	"para2/tokenizer"
)

func ExampleTokenize() {
	text := "Who's on first?"
	for _, tok := range tokenizer.Tokenize(text) {
		fmt.Println(tok)
	}

	// Output:
	// first
}
