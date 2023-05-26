// Package gobdb provides a simple key-value database for Go binary objects
// encoded with encoding/gob.
package gobdb

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
)

// Gobdb represents a simple key-value database for Go binary objects encoded
// with encoding/gob.
type Gobdb [T any] struct {
	Data []T
	path string
}

// Open opens a file at the specified path and decodes its contents using the
// gob decoder.
//
// If the decoding fails, it returns an error.
// If the decoding is successful, it returns a gobdb object with the decoded
// data.
func Open[T any](path string) (Gobdb[T], error)  {
	var t T
	gob.Register(t)

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return Gobdb[T]{}, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var data[] T
	err = decoder.Decode(&data)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return Gobdb[T]{}, err
		}
	}
	return Gobdb[T]{
		Data: data,
		path: path,
	}, nil
}

// List returns the entire database as a map.
func (db Gobdb[T]) List() []T {
	return db.Data
}

// Add append a new key-value pair of data to the db and persist data to disk
// using the given `path` to Open()
func (db *Gobdb[T]) Add(d ...T) error {
	file, err := os.Create(db.path)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	db.Data = append(db.Data, d...)
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&db.Data)
	if err != nil {
		return fmt.Errorf("unable to encode collection: %w", err)
	}

	return nil
}
