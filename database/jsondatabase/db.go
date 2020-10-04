package jsondatabase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
func (jsnDb *JSONDatabase) Add(key [2]string, elem string) error {
	data, err := read(jsnDb.filepath)
	if err != nil {
		return err
	}

	data, err = add(data, key, elem)
	if err != nil {
		return err
	}

	err = write(jsnDb.filepath, data)
	if err != nil {
		return err
	}

	return nil
}

func add(
	data map[string][]string,
	keyArr [2]string,
	elem string,
) (map[string][]string, error) {
	key, err := stringifyKey(keyArr)
	if err != nil {
		return nil, err
	}

	// Get the list of words that correspont to the key
	if value, ok := data[key]; !ok {
		// Create the key and his new value
		data[key] = []string{elem}
	} else if !contains(&value, elem) {
		// Add the new word in the words list if he doesn't exist
		value = append(value, elem)
		data[key] = value
	}

	return data, nil
}

// Random entry from the key subset
func (jsnDb *JSONDatabase) Random(key [2]string) (string, error) {
	words, err := jsnDb.Get(key)
	if err != nil {
		return "", nil
	}
	return random(words)
}

func random(words []string) (string, error) {
	nbWords := len(words)
	if nbWords <= 0 {
		// The key doesn't have any subset value
		return "", errors.New("The key haven't any words in his subset")
	}
	// Choose a random index
	idx := rand.Intn(nbWords)
	return words[idx], nil
}

// Get the value from the key
func (jsnDb *JSONDatabase) Get(key [2]string) ([]string, error) {
	data, err := read(jsnDb.filepath)
	if err != nil {
		return nil, err
	}
	return get(data, key)
}

func get(data map[string][]string, keyArr [2]string) ([]string, error) {
	key, err := stringifyKey(keyArr)
	if err != nil {
		return nil, err
	}
	// Check if the key exist
	if value, ok := data[key]; ok {
		return value, nil
	}
	return nil, errors.New("Key not found")
}

// Set the value to the key
func (jsnDb *JSONDatabase) Set(key [2]string, value []string) error {
	data, err := read(jsnDb.filepath)
	if err != nil {
		return err
	}

	data, err = set(data, key, value)
	if err != nil {
		return err
	}

	err = write(jsnDb.filepath, data)
	if err != nil {
		return err
	}
	return nil
}

func set(
	data map[string][]string,
	keyArr [2]string,
	value []string,
) (map[string][]string, error) {
	key, err := stringifyKey(keyArr)
	if err != nil {
		return nil, err
	}

	data[key] = value
	return data, nil
}

func contains(list *[]string, item string) bool {
	for _, elem := range *list {
		if elem == item {
			return true
		}
	}
	return false
}

func read(path string) (map[string][]string, error) {
	// Read file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string][]string

	// Parse json content
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func write(path string, data map[string][]string) error {
	// Stringify data
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func stringifyKey(keyArr [2]string) (string, error) {
	k, err := json.Marshal(keyArr)
	key := string(k)
	if err != nil {
		return "", err
	}
	return key, nil
}
