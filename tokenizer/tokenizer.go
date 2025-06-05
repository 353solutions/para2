package tokenizer

import (
	"regexp"
	"slices"
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
	return slices.Contains(stopWords, word)
}

var stopWords = []string{
	"a",
	"about",
	"above",
	"after",
	"again",
	"against",
	"all",
	"am",
	"an",
	"and",
	"any",
	"are",
	"as",
	"at",
	"be",
	"because",
	"been",
	"before",
	"being",
	"below",
	"between",
	"both",
	"but",
	"by",
	"can",
	"did",
	"do",
	"does",
	"doing",
	"don",
	"down",
	"during",
	"each",
	"few",
	"for",
	"from",
	"further",
	"had",
	"has",
	"have",
	"having",
	"he",
	"her",
	"here",
	"hers",
	"herself",
	"him",
	"himself",
	"his",
	"how",
	"i",
	"if",
	"in",
	"into",
	"is",
	"it",
	"its",
	"itself",
	"just",
	"me",
	"more",
	"most",
	"my",
	"myself",
	"no",
	"nor",
	"not",
	"now",
	"of",
	"off",
	"on",
	"once",
	"only",
	"or",
	"other",
	"our",
	"ours",
	"ourselves",
	"out",
	"over",
	"own",
	"s",
	"same",
	"she",
	"should",
	"so",
	"some",
	"such",
	"t",
	"than",
	"that",
	"the",
	"their",
	"theirs",
	"them",
	"themselves",
	"then",
	"there",
	"these",
	"they",
	"this",
	"those",
	"through",
	"to",
	"too",
	"under",
	"until",
	"up",
	"very",
	"was",
	"we",
	"were",
	"what",
	"when",
	"where",
	"which",
	"while",
	"who",
	"whom",
	"why",
	"will",
	"with",
	"you",
	"your",
	"yours",
	"yourself",
	"yourselves",
}
