package jsondatabase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

// JSONDatabase struct implement Base interface
type JSONDatabase struct {
	filepath string
}

// New json db that implement the Base interface
func New(path string) JSONDatabase {
	return JSONDatabase{
		filepath: path,
	}
}

// Add an entry to the database
func (jsnDb *JSONDatabase) Add(key [2]string, elem string) {
	data := read(jsnDb.filepath)
	data = add(data, key, elem)
	write(jsnDb.filepath, data)
}

func add(
	data map[[2]string][]string,
	key [2]string,
	elem string,
) map[[2]string][]string {
	// Get the list of words that correspont to the key
	if value, ok := data[key]; !ok {
		// Create the key and his new value
		data[key] = []string{elem}
	} else if !contains(&value, elem) {
		// Add the new word in the words list if he doesn't exist
		value = append(value, elem)
		data[key] = value
	}
	return data
}

// Random entry from the key subset
func (jsnDb *JSONDatabase) Random(key [2]string) string {
	words := jsnDb.Get(key)
	return random(words)
}

func random(
	words []string,
) string {
	nbWords := len(words)
	if nbWords <= 0 {
		// The key doesn't have any subset value
		return ""
	}
	// Choose a random index
	idx := rand.Intn(nbWords)
	return words[idx]
}

// Get the value from the key
func (jsnDb *JSONDatabase) Get(key [2]string) []string {
	data := read(jsnDb.filepath)
	return get(data, key)
}

func get(
	data map[[2]string][]string,
	key [2]string,
) []string {
	// Check if the key exist
	if value, ok := data[key]; ok {
		return value
	}
	return make([]string, 0)
}

// Set the value to the key
func (jsnDb *JSONDatabase) Set(key [2]string, value []string) {
	data := read(jsnDb.filepath)
	data = set(data, key, value)
	write(jsnDb.filepath, data)
}

func set(
	data map[[2]string][]string,
	key [2]string,
	value []string,
) map[[2]string][]string {
	data[key] = value
	return data
}

func contains(list *[]string, item string) bool {
	for _, elem := range *list {
		if elem == item {
			return true
		}
	}
	return false
}

func read(path string) map[[2]string][]string {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data map[[2]string][]string

	// Parse json content
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func write(path string, data map[[2]string][]string) {
	// Stringify data
	content, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
