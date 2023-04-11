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

// gobdb represents a simple key-value database for Go binary objects encoded
// with encoding/gob.
type gobdb struct {
	Data Data
	path string
	customTypes []any
}

// Data is a type alias for a map with keys and values of any type.
type Data map[any]any

func WithType(t any) func(*gobdb) {
	return func(db *gobdb) {
		db.customTypes = append(db.customTypes, t)
	}
}

type option func(*gobdb)

// Open opens a file at the specified path and decodes its contents using the
// gob decoder.
//
// If the decoding fails, it returns an error.
// If the decoding is successful, it returns a gobdb object with the decoded
// data.
func Open(path string, opts ...option) (gobdb, error)  {
	db := &gobdb{}
	for _, optf := range opts {
		optf(db)
	}
	for _, t := range db.customTypes {
		gob.Register(t)
	}

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return gobdb{}, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	data := make(Data)
	err = decoder.Decode(&data)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return gobdb{}, err
		}
	}
	return gobdb{
		Data: data,
		path: path,
	}, nil
}

// List returns the entire database as a map.
func (db gobdb) List() Data {
	return db.Data
}

// Add append a new key-value pair of data to the db and persist data to disk
// using the given `path` to Open()
func (db *gobdb) Add(d Data) error {
	for k, v := range d {
		// if the caller uses a custom type such as struct as value, then we
		// need to register this value to gob so it an encode it, otherwise:
		//
		//    unable to encode collection: gob: type not registered for interface: []gobdb_test.Todo
		gob.Register(v)
		db.Data[k] = v

	}
	file, err := os.Create(db.path)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&db.Data)
	if err != nil {
		return fmt.Errorf("unable to encode collection: %w", err)
	}

	return nil
}
