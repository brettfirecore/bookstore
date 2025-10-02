// Package bookstore defines tiny domain types and functions for a bookstore.
// This file contains the production code exercised by bookstore_test.go.
//
// Design notes (for this chapter’s exercises):
//   • We model a book as a simple struct of exported fields so tests (in an external
//     *_test package) and other packages can construct/read values directly.
//   • We look books up by a unique integer ID. For now, IDs are assumed to be
//     assigned already (e.g., by fixtures in tests).
//   • GetBook returns a single Book value. If the ID isn’t found, it returns the
//     zero value (Book{}). That keeps the API minimal for now; an alternative
//     you might adopt later is (Book, bool) to signal “not found” explicitly.
package bookstore

// Book represents a single title in our catalog.
//
// Field visibility:
//   • Fields start with capital letters, so they’re exported (public) and can be set
//     and read from other packages, including external tests.
//   • All fields are plain value types (string/int), so copying Book values is cheap
//     and safe with no internal pointers to worry about.
//
// Invariants (by convention, not enforced here):
//   • Copies should never go negative.
//   • ID should be unique within a given catalog slice.
type Book struct {
	Title  string // Human-readable title (e.g., "For the Love of Go")
	Author string // Author’s name (optional for these exercises)
	Copies int    // How many copies are in stock (0 means sold out)
	ID     int    // Unique identifier used for lookups (e.g., 1, 2, 42)
}

// GetBook scans the provided catalog (a slice of Book) and returns the Book whose
// ID matches the requested id. If no match exists, it returns the zero value: Book{}.
//
// Semantics:
//   • Lookup strategy is a linear scan (O(n)) using a for-range loop. That’s the
//     simplest correct approach for small in-memory catalogs used in tests.
//   • Return type is a single Book value. “Not found” is represented by Book{}.
//     This keeps the function signature minimal for now and matches the chapter’s
//     goal. If you later want an explicit signal for absence, change the signature to
//     (Book, bool) and return (Book{}, false) when not found.
//
// Implementation details:
//   • `for _, b := range catalog` iterates over each element. `b` is a *copy* of the
//     slice element at that position; returning `b` therefore returns a copy of the
//     catalog element and does not mutate the caller’s slice.
//   • If duplicate IDs exist (they shouldn’t), the first match wins.
// The underscore (_) means "ignore this value" (INDEX)
func GetBook(catalog []Book, id int) Book {
	for _, b := range catalog { // visit each book in order (index is irrelevant here)
		if b.ID == id { // is this the one we’re looking for?
			return b // found: return a copy of this element
		}
	}
	// Not found: return the zero value. Callers can compare against (Book{}) to
	// detect this case if they need to. (Alternatively, switch to (Book, bool).)
	return Book{}
}
