// Package bookstore defines types and functions related to managing a collection of books.
package bookstore

// Book represents the details of a single book in the catalog.
type Book struct {
	Title  string // The title of the book
	Author string // The author's name
	Copies int    // The number of copies available
	ID     int    // A unique identifier for the book
}

// GetAllBooks returns a slice of all Book values in the given catalog map.
//
// Parameters:
// - catalog: a map where the key is an int (usually the book ID) and
//   the value is a Book struct.
//
// Returns:
// - A slice of Book structs, containing every Book from the map.
func GetAllBooks(catalog map[int]Book) []Book {
	// Start with an empty slice to hold the result.
	// This slice will grow as we append each book from the map.
	result := []Book{}

	// Use a for loop to iterate over the map.
	// We ignore the key (`_`) since we're only interested in the Book values.
	for _, b := range catalog {
		// Append each Book (value in the map) to the result slice.
		result = append(result, b)
	}

	// After collecting all the books, return the final slice.
	return result
}
