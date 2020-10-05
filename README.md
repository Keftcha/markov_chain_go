# Markov Chain Go

Markov chain implemented in go.

## Usage

We get a MarkovChainGo struct with the `New` function:

```go
package main

import "github.com/Keftcha/markovchaingo"

mcg = markogchaingo.New("<connection_string>")
```

mcg will now be of type `Base` which is an interface to your database.

For example, if you want to use a `map[[2]string][]string` data structure as
your database, your connection string must be `"in-memory://_".`  
You also can choosing to store your data un a json with the connection string
`"file:///path/to/my/file.json"`.

Then, the `mcg` struct will have two methods, `Learn` and `Talk`.

### The `Learn` method

The `Learn` method is use to learn from a sentence.

```go
func (mcg *MarkovChainGo) Learn(text string) error {
    [...]
}
```

It could return an error.

### The `Talk` method

The `Talk` method is used to generate a sentence from what it learn.

```go
func (mcg *MarkovChainGo) Talk() (string, error) {
    [...]
}
```

It return a string and may be an error.

## Connection string

| Database type | Connection string pattern | Example |
| ------------- | ------------------------- | ------- |
| In memory `map[[2]string][]string` data structure | `in-memory://<whatever you want>` | `in-memory://_` |
| Json | `file://<path to file>` | `file://./data/sentences.json`

## Contribute

If you want to add your database, be pleased.

First you need to add your package in the `./database/` directory.  
If you add a mackage for PostgreSQL (for example) please name it
`portgresqldatabase` (in a directory of the same name).  

In your package put at least two files:

- `db.go`, which is the implementation of your struct that implement the `Base`
    interface
- `db_test.go`, which is unit test of your struct and his methods

If you add other files, please, make unit test on them.

Then, don't forget to add your case in the `./database/get.go` file.  
So we will be able to choose it in term of our connection string.
