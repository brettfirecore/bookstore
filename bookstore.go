// Package bookstore defines tiny domain types and functions for a bookstore.
package bookstore

import "fmt"

type Book struct {
	Title  string // Readable title (e.g., "For the Love of Go")
	Author string // Author’s name
	Copies int    // Copies in stock (0 means sold out)
	ID     int    // Unique identifier used for lookups (e.g., 1, 2, 42)
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID &d doessn't exist", ID)
	}
	return b, nil
}
