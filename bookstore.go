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

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	books := make([]Book, 0, len(c))
	for _, b := range c {
		books = append(books, b)
	}
	return books
}

func (c Catalog) GetBook(id int) (Book, bool) {
	b, ok := c[id]
	return b, ok
}
