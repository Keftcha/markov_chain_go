package inmemorydatabase

import (
	"errors"
	"math/rand"
)

// InMemoryDatabase struct implement Base interface
type InMemoryDatabase struct {
	data map[[2]string][]string
}

// New in-memory db that implement the Base interface
func New() InMemoryDatabase {
	return InMemoryDatabase{
		data: make(map[[2]string][]string),
	}
}

// Add an entry to the database
func (inMemDb *InMemoryDatabase) Add(key [2]string, elem string) error {
	// Get the list of words that correspont to the key
	value, _ := inMemDb.Get(key)
	// Add the new word in the words list if he doesn't exist
	if !contains(&value, elem) {
		value = append(value, elem)
		inMemDb.Set(key, value)
	}

	return nil
}

// Random entry from the key subset
func (inMemDb *InMemoryDatabase) Random(key [2]string) (string, error) {
	// Get the list of words that correspont to the key
	value, err := inMemDb.Get(key)
	if err != nil {
		return "", err
	}
	nbWords := len(value)

	// There is no key, or the key haven't any subset
	if nbWords == 0 {
		return "", errors.New("The key haven't any words in his subset")
	}

	// Choose a ramdom number (idx of the word)
	idx := rand.Intn(nbWords)

	return value[idx], nil
}

// Get the value from the key
func (inMemDb *InMemoryDatabase) Get(key [2]string) ([]string, error) {
	// Check if the key exist
	if value, ok := inMemDb.data[key]; ok {
		return value, nil
	}
	return nil, errors.New("Key not found")

}

// Set the value to the key
func (inMemDb *InMemoryDatabase) Set(key [2]string, value []string) error {
	inMemDb.data[key] = value
	return nil
}

func contains(list *[]string, item string) bool {
	for _, elem := range *list {
		if elem == item {
			return true
		}
	}
	return false
}
