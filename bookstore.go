// Package bookstore does simple calculations and types for a tiny bookstore domain.
// This is the code under test — each exported identifier here is exercised by bookstore_test.go.
package bookstore

import "errors" // We use errors.New to produce a descriptive error when stock is empty.

// Book represents one book in inventory.
// Because Title, Author, and Copies start with capital letters, they are exported
// and therefore part of the package's public API (visible to other packages).
type Book struct {
	Title  string // Human-readable title of the book.
	Author string // Author's name.
	Copies int    // How many copies are currently in stock (>= 0).
}

// Buy attempts to purchase a single copy of b and returns the updated Book and an error.
// Behaviour:
//   - If no copies remain (Copies == 0), it returns the zero Book value and a non-nil error.
//   - Otherwise it decrements the Copies count by 1 and returns the updated Book and a nil error.
//
// Note:
//   - The parameter b is passed BY VALUE, so this function works on a copy. The caller's
//     original variable is not mutated; use the returned Book.
//   - Returning the zero value on error is a common Go convention: the data result is
//     meaningless when an error is present.
func Buy(b Book) (Book, error) {
	// Guard: prevent negative inventory by refusing the purchase if stock is empty.
	// Consider changing to "b.Copies <= 0" if you also want to reject malformed inputs.
	if b.Copies == 0 {
		// On error, return the zero Book value and a descriptive error.
		return Book{}, errors.New("no copies left")
	}

	// Happy path: there is stock. Decrement the (copied) count and return it.
	b.Copies--

	// Return the updated Book and a nil error to signal success.
	return b, nil
}
