package markovchaingo

import (
	"github.com/keftcha/markovchaingo/database"
	"strings"
)

// MarkovChainGo struct
type MarkovChainGo struct {
	db database.Base
}

// New return a pointer to a new instance of MarkovChainGo
func New(connectionString string) *MarkovChainGo {
	mcg := new(MarkovChainGo)
	mcg.db = database.Get(connectionString)
	return mcg
}

// Learn will learn from the given text
func (mcg *MarkovChainGo) Learn(text string) {
	if text == "" {
		return
	}

	words := splitMessage(text)
	// Begin to the first value to add, so idx start at 2
	for idx := 2; idx < len(words); idx++ {
		key := [2]string{words[idx-2], words[idx-1]}
		value := words[idx]

		mcg.db.Add(key, value)
	}
}

// Talk return a random sentence
func (mcg *MarkovChainGo) Talk() string {
	key := [2]string{"", ""}
	// Our sentence
	words := make([]string, 0)

	for {
		// Get a word to add to the sentence
		word := mcg.db.Random(key)
		// Check if we have finish the sentence
		if word == "\x03" {
			break
		}

		// Add the word to the sentence
		words = append(words, word)
		// Change the key
		key[0], key[1] = key[1], word
	}

	return strings.Join(words, " ")
}

func splitMessage(text string) []string {
	words := []string{"", ""}
	if text != "" {
		words = append(words, strings.Split(text, " ")...)
	}
	words = append(words, "\x03")

	return words
}
