package database

// Base interface is a database interface
// to support multiple database
type Base interface {
	// Add an entry to the database
	Add(key [2]string, elem string) error

	// Pick a random entry from the key subset
	Random(key [2]string) (string, error)

	// Get the value form the key
	Get(key [2]string) ([]string, error)

	// Set the value to the key
	Set(key [2]string, value []string) error
}
