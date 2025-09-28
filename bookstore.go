// Package bookstore does simple calculations.
// This   is the code under test — each function here will be tested in bookstore_test.go.
package bookstore

// Book represents a book in the bookstore.
// It has fields (Title, Author, Copies) that describe one book.
type Book struct {
	Title  string // The title of the book
	Author string // The author of the book
	Copies int    // How many copies of this book are currently in stock
}

// Buy takes a Book and returns a new Book value.
// Right now it just returns an empty Book — a "zero value" — so our test will compile
// but fail, which is the first step in test-driven development.
func Buy(b Book) Book {
	return Book{} // return an empty Book (zero value) for now
}
