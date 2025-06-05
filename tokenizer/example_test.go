package tokenizer

import "fmt"

func ExampleTokenize() {
	text := "Who's on first?"
	for _, tok := range Tokenize(text) {
		fmt.Println(tok)
	}

	// Output:
	// first
}
