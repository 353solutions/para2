package tokenizer

import (
	"regexp"
	"strings"
)

/*
func init() {
	var err error
	wordRe, err = regexp.Compile(`[a-zA-Z]+`)
	if err != nil {
		panic(err)
	}
}
*/

var suffixes = []string{"ed", "ing", "s"}

// working, works, worked -> work
func Stem(word string) string {
	for _, s := range suffixes {
		if strings.HasSuffix(word, s) {
			return word[:len(word)-len(s)]
		}
	}

	return word
}

// "Who's on first?" -> [who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func initialSplit(text string) []string {
	return wordRe.FindAllString(text, -1)
}

func Tokenize(text string) []string {
	words := initialSplit(text)

	var tokens []string
	for _, tok := range words {
		tok = strings.ToLower(tok)
		if IsStop(tok) {
			continue
		}

		tok = Stem(tok)
		if tok != "" {
			tokens = append(tokens, tok)
		}
	}

	return tokens
}

func IsStop(word string) bool {
	return stopWords[word]
}

var stopWords = map[string]bool{
	"a":          true,
	"about":      true,
	"above":      true,
	"after":      true,
	"again":      true,
	"against":    true,
	"all":        true,
	"am":         true,
	"an":         true,
	"and":        true,
	"any":        true,
	"are":        true,
	"as":         true,
	"at":         true,
	"be":         true,
	"because":    true,
	"been":       true,
	"before":     true,
	"being":      true,
	"below":      true,
	"between":    true,
	"both":       true,
	"but":        true,
	"by":         true,
	"can":        true,
	"did":        true,
	"do":         true,
	"does":       true,
	"doing":      true,
	"don":        true,
	"down":       true,
	"during":     true,
	"each":       true,
	"few":        true,
	"for":        true,
	"from":       true,
	"further":    true,
	"had":        true,
	"has":        true,
	"have":       true,
	"having":     true,
	"he":         true,
	"her":        true,
	"here":       true,
	"hers":       true,
	"herself":    true,
	"him":        true,
	"himself":    true,
	"his":        true,
	"how":        true,
	"i":          true,
	"if":         true,
	"in":         true,
	"into":       true,
	"is":         true,
	"it":         true,
	"its":        true,
	"itself":     true,
	"just":       true,
	"me":         true,
	"more":       true,
	"most":       true,
	"my":         true,
	"myself":     true,
	"no":         true,
	"nor":        true,
	"not":        true,
	"now":        true,
	"of":         true,
	"off":        true,
	"on":         true,
	"once":       true,
	"only":       true,
	"or":         true,
	"other":      true,
	"our":        true,
	"ours":       true,
	"ourselves":  true,
	"out":        true,
	"over":       true,
	"own":        true,
	"s":          true,
	"same":       true,
	"she":        true,
	"should":     true,
	"so":         true,
	"some":       true,
	"such":       true,
	"t":          true,
	"than":       true,
	"that":       true,
	"the":        true,
	"their":      true,
	"theirs":     true,
	"them":       true,
	"themselves": true,
	"then":       true,
	"there":      true,
	"these":      true,
	"they":       true,
	"this":       true,
	"those":      true,
	"through":    true,
	"to":         true,
	"too":        true,
	"under":      true,
	"until":      true,
	"up":         true,
	"very":       true,
	"was":        true,
	"we":         true,
	"were":       true,
	"what":       true,
	"when":       true,
	"where":      true,
	"which":      true,
	"while":      true,
	"who":        true,
	"whom":       true,
	"why":        true,
	"will":       true,
	"with":       true,
	"you":        true,
	"your":       true,
	"yours":      true,
	"yourself":   true,
	"yourselves": true,
}
