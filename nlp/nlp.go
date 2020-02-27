//Package nlp provides basic NLP functionality
package nlp

//We usually have 3 import sections:
// - golang
// - my libraries
// - 3rd party libraries}

import (
	"regexp"
	"strings"
)

import (
	"git.att.com/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+") //will panic if the provided regexp is invalid
)

//Tokenize returns a slice of tokens that exist in text
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		stem := stemmer.Stem(token)
		tokens = append(tokens, stem)
	}
	stemmer.Stem("")
	return tokens
}
