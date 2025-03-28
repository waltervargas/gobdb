// Package gobdb provides a simple key-value database for Go binary objects
// encoded with encoding/gob.
package gobdb

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slices"
)

// Gobdb represents a simple key-value database for Go binary objects encoded
// with encoding/gob. It is a generic type that can store any type of data in
// its Data field. The path field is unexported and contains the path to the
// file where the database is stored.
//
// The Gobdb type parameter T is constrained by the 'comparable' type
// constraint, meaning it can be any comparable Go type. See
// https://go.dev/ref/spec#Comparison_operators for more information.
//
// Fields:
//
//	Data: A slice of type T. It contains the values of the database.
//	path: A string representing the path to the file where the database is stored.
type Gobdb[T comparable] struct {
	Data []T
	path string
}

// Open is a generic function that opens a file at the specified path and
// decodes its contents using gob decoder. The type of the data in the
// file should be specified as the type argument T when calling this function.
//
// The type parameter T is constrained by the 'comparable' type constraint, meaning
// it can be any comparable Go type.
//
// If the decoding fails, it returns an error. If the decoding is successful,
// it returns a Gobdb object with the decoded data.
//
// Parameters:
//
//	path: A string representing the path to the file to open and decode.
//
// Returns:
//
//	A Gobdb object containing the decoded data, and a nil error if the
//	operation was successful. If an error occurred, the Gobdb object will
//	be empty and the error will contain details about what went wrong.
func Open[T comparable](path string) (Gobdb[T], error) {
	var t T
	gob.Register(t)

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return Gobdb[T]{}, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var data []T
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

// List is a method on the Gobdb type that returns all the data in the
// database as a slice. The type of the data returned is the same as the type
// parameter T of the Gobdb object.
//
// The type parameter T is constrained by the 'any' type constraint, meaning
// it can be any Go type.
//
// This method does not take any parameters.
//
// Returns:
//
//	A slice of type T containing all the data in the database.
func (db Gobdb[T]) List() []T {
	return db.Data
}

// Add is a method on the Gobdb type that adds new data to the database and
// persists the updated database to disk. The data to be added should be of
// the same type as the type parameter T of the Gobdb object.
//
// The type parameter T is constrained by the 'any' type constraint, meaning
// it can be any Go type.
//
// The method uses the path stored in the Gobdb object to open the file and
// encodes the updated data using the gob encoder. If an error occurs during
// this process, it returns an error.
//
// Parameters:
//
//	d: A variadic parameter of type T representing the data to add to the
//	   database.
//
// Returns:
//
//	If the operation is successful, the method returns nil. If an error occurs
//	during the process, the method returns the error with details about what
//	went wrong.
func (db *Gobdb[T]) Add(d ...T) error {
	db.Data = append(db.Data, d...)
	return db.Sync()
}

func (db *Gobdb[T]) Sync() error {
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

func (db *Gobdb[T]) Delete(vals ...T) error {
	var datatmp []T
	// we keep the db.Data values that are not in vals.
	for _, v := range db.Data {
		if !slices.Contains(vals, v) {
			datatmp = append(datatmp, v)
		}
	}
	db.Data = datatmp
	return db.Sync()
}

// DeleteAll deletes all values from the database and syncs the changes to
// the file. It clears the Data slice and writes an empty slice to the file.
func (db *Gobdb[T]) DeleteAll() error {
	db.Data = []T{}
	return db.Sync()
}
