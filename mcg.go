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
func (mcg *MarkovChainGo) Learn(text string) error {
	if text == "" {
		return nil
	}

	words := splitMessage(text)
	// Begin to the first value to add, so idx start at 2
	for idx := 2; idx < len(words); idx++ {
		key := [2]string{words[idx-2], words[idx-1]}
		value := words[idx]

		err := mcg.db.Add(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// Talk return a random sentence
func (mcg *MarkovChainGo) Talk() (string, error) {
	key := [2]string{"", ""}
	// Our sentence
	words := make([]string, 0)

	for {
		// Get a word to add to the sentence
		word, err := mcg.db.Random(key)
		if err != nil {
			return "", err
		}
		// Check if we have finish the sentence
		if word == "\x03" {
			break
		}

		// Add the word to the sentence
		words = append(words, word)
		// Change the key
		key[0], key[1] = key[1], word
	}

	return strings.Join(words, " "), nil
}

func splitMessage(text string) []string {
	words := []string{"", ""}
	if text != "" {
		words = append(words, strings.Split(text, " ")...)
	}
	words = append(words, "\x03")

	return words
}
