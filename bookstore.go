// Package bookstore defines a Book type and methods for setting and getting book information.
// This file lives in the 'bookstore' folder and contains logic related to a single Book.
package bookstore

import "fmt" // Import the fmt package for formatted error messages.

// Book represents information about a single book in the store.
type Book struct {
	Title    string // Title is exported (public) so it can be accessed outside the package.
	Author   string // Author is exported (public) for the same reason.
	category string // category is unexported (private) and only accessible within this package.
	// PriceCents int // commented out for now; could store price in cents.
}

// SetCategory validates and assigns a category to the book.
// Because it modifies the receiver, it uses a *pointer receiver (*Book).
func (b *Book) SetCategory(category string) error {
	// If the given category is not "Autobiography", return an error.
	if category != "Autobiography" {
		return fmt.Errorf("unknown category %q", category)
	}

	// Otherwise, set the category field on the actual book.
	b.category = category

	// Return nil to indicate success (no error).
	return nil
}

// Category returns the current category of the book.
// This uses a value receiver because it only *reads* data, not modifies it.
func (b Book) Category() string {
	return b.category // Return the book's category string.
}
