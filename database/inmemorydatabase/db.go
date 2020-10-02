package inmemorydatabase

import (
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
func (inMemDb *InMemoryDatabase) Add(key [2]string, elem string) {
	// Get the list of words that correspont to the key
	value := inMemDb.Get(key)
	// Add the new word in the words list if he doesn't exist
	value = append(value, elem)
	inMemDb.Set(key, value)
}

// Random entry from the key subset
func (inMemDb *InMemoryDatabase) Random(key [2]string) string {
	// Get the list of words that correspont to the key
	value := inMemDb.Get(key)
	nbWords := len(value)

	// There is no key, or the key haven't any subset
	if nbWords == 0 {
		return ""
	}

	// Choose a ramdom number (idx of the word)
	idx := rand.Intn(nbWords)

	return value[idx]
}

// Get the value from the key
func (inMemDb *InMemoryDatabase) Get(key [2]string) []string {
	// Check if the key exist

	value, ok := inMemDb.data[key]
	if ok {
		return value
	}
	return make([]string, 0)
}

// Set the value to the key
func (inMemDb *InMemoryDatabase) Set(key [2]string, value []string) {
	inMemDb.data[key] = value
}

func contains(list *[]string, item string) bool {
	for _, elem := range *list {
		if elem == item {
			return true
		}
	}
	return false
}
