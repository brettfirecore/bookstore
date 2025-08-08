// Package bookstore defines types and functions related to managing a collection of books.
// This is the top-level package declaration. Every Go file must begin with a `package` line.
package bookstore

// Book represents a single book in the catalog.
// It contains an ID, title, price in cents, and an optional discount percentage.
type Book struct {
	Title           string // The book's title
	Author          string //
	Copies          int    //
	ID              int    // Unique identifier for the book
	PriceCents      int    // Price stored as cents to avoid floating-point issues
	DiscountPercent int    // Discount percentage (0–100)
}

// Catalog is a custom type that represents the bookstore's inventory.
// Internally, it's a map where the key is the book's ID and the value is the Book struct.
// We use a custom type instead of map[int]Book directly so we can add methods to it.
type Catalog map[int]Book

// GetAllBooks is a method on the Catalog type that returns a slice of all books in the catalog.
func (c Catalog) GetAllBooks() []Book {
	// Create a slice with an initial capacity equal to the catalog size.
	// This avoids unnecessary allocations when appending.
	books := make([]Book, 0, len(c))

	// Loop through the map values (books) in the catalog.
	for _, b := range c {
		// Append each book to the slice.
		books = append(books, b)
	}

	// Return the complete slice of books.
	// The order will be random because Go maps are unordered.
	return books
}
