package markovchaingo

import (
	"strings"
)

// MarkovChainGo struct
type MarkovChainGo struct {
}

// Learn will learn from the given text
func (mcg *MarkovChainGo) Learn(text string) {
}

// Talk return a random sentence
func (mcg *MarkovChainGo) Talk() string {
	return ""
}

// New return a pointer to a new instance of MarkovChainGo
func New() *MarkovChainGo {
	return &MarkovChainGo{}
}

func splitMessage(text string) []string {
	words := []string{"", ""}
	if text != "" {
		words = append(words, strings.Split(text, " ")...)
	}
	words = append(words, "\x03")

	return words
}
