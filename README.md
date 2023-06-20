[![Go Reference](https://pkg.go.dev/badge/github.com/waltervargas/gobdb.svg)](https://pkg.go.dev/github.com/waltervargas/gobdb)[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)[![Go Report Card](https://goreportcard.com/badge/github.com/waltervargas/gobdb)](https://goreportcard.com/report/github.com/waltervargas/gobdb)

# gobdb

gobdb is a simple key-value database for Go binary objects encoded with
encoding/gob.

## Features

- Open: Opens a file at the specified path and decodes its contents using the
  gob decoder.
- List: Returns the entire database as an array.
- Add: Appends a new value to the database and persist data to disk.

## Usage
### Open a Database


```go
db, err := gobdb.Open[YourType]("path/to/yourfile.gob")
if err != nil {
    // handle error
}
```

### List the Database Contents

```go
data := db.List()
for _, item := range data {
    // process item
}
```

### Add Data to the Database

```go
err := db.Add(yourData)
if err != nil {
    // handle error
}
```

## Use Cases

This package is ideal for simple use-cases where you need to persist and
retrieve Go binary objects without the overhead of a full-fledged database
system. Here are a few scenarios where it can be handy:

- Caching: You can use gobdb to cache complex data structures that take time to
  generate or fetch from a slow source. This way, you can quickly retrieve the
  data from the disk whenever needed.

- Local Data Persistence: If you're building a Go application that needs to
  store data locally, you can use gobdb as a lightweight, file-based database.

- Data Serialization: The package can also be used for serializing and
  deserializing Go data structures for network communication or other purposes
  where you need to convert Go objects to a binary format and vice versa.

## TODO

While gobdb provides a convenient way to work with gob encoded data, there are
several features that could improve its utility:

- Delete Functionality: Currently, there's no way to delete data from the
  database. This feature is crucial for managing the data efficiently.

- Update Functionality: Another missing feature is the ability to update an
  existing value in the database.

- Search Functionality: The package currently lacks a way to search for a
  specific data object in the database. Implementing this would provide quicker
  access to data.

- Concurrency Safety: The package isn't currently safe for concurrent use. It
  would be beneficial to add some form of locking to ensure that concurrent
  reads and writes don't cause data corruption.

- Error Handling: The current error handling could be improved, for instance by
  providing more detailed error messages or different types of errors for
  different failure cases.

- Batch Operations: Adding support for batch operations could improve
  performance when working with large amounts of data.

Contributions are welcome to help implement these features and improve this
package.
