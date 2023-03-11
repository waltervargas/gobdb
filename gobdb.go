package gobdb

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
)

// gobdb represents a simple key-value database for Go binary objects encoded
// with encoding/gob.
type gobdb struct {
	Collection Collection
}

// Collection defines a slice of objects.
type Collection []Object

// Object defines a structure that holds any value.
type Object struct {
	Val any
}

// Open opens a file at the specified path and decodes its contents using the
// gob decoder.
// If the decoding fails, it returns an error.
// If the file is empty, it returns an empty collection.
// If the decoding is successful, it returns a gobdb object with the decoded
// collection.
func Open(path string) (gobdb, error)  {
	file, err := os.Open(path)
	if err != nil {
		return gobdb{}, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var collection Collection
	err = decoder.Decode(&collection)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return gobdb{}, err
		}
	}
	return gobdb{Collection: collection}, nil
}


// List returns the collection stored in the gobdb object.
// It takes no arguments and returns the collection and an error.
func (db gobdb) List() (Collection, error) {
	return db.Collection, nil
}
