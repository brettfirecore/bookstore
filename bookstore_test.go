// Declare that this file belongs to the "bookstore_test" package (note the _test suffix).
// Keeping tests in an *external* package means tests can only touch the public API
// (exported identifiers) of the bookstore package—exactly like real users would.
// This helps catch accidental reliance on unexported internals.
// bookstore_test.go
package bookstore_test

import (
	"testing" // The Go testing framework: provides *testing.T and helpers.

	// cmp is a small library that compares Go values and prints readable diffs.
	// It’s great for tests because it shows *what* differs, not just “not equal”.
	"github.com/google/go-cmp/cmp"

	// Import the package under test. We alias it to "bookstore" so we can write
	// bookstore.Book, bookstore.GetBook, etc. (This matches the module name
	// in your go.mod: module github.com/brettfirecore/bookstore)
	bookstore "github.com/brettfirecore/bookstore"
)

// TestGetBook verifies the behavior: given a catalog with unique IDs, looking up
// by ID should return the matching Book value. This test also protects against a
// bogus implementation that always returns the first element.
func TestGetBook(t *testing.T) {
	t.Parallel() // Allow this test to run concurrently with others in the package.

	// Arrange: build a tiny in-memory catalog. Using a slice literal keeps the
	// “world setup” close to the test. Each Book has a unique ID so we can look it up.
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	catalog[2] = bookstore.Book{ID: 2, Title: "The Power of Go: Tools (2nd ed.)"}
	// We expect the *second* book to be returned when we ask for ID=2. Picking
	// the second element (rather than the first) ensures the test would fail if the
	// implementation incorrectly returned catalog[0] for every request.


	// Act: call the function under test. The API returns a single Book value; if the
	// ID didn’t exist, we’d get the zero value (Book{}). Here, ID=2 *does* exist.
    got := bookstore.GetBook(catalog, 2)

    want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools (2nd ed.)"}
    if diff := cmp.Diff(want, got); diff != "" {
        t.Fatalf("mismatch (-want +got):\n%s", diff)
    }
}

	// Assert: compare the structs. We prefer cmp.Equal + cmp.Diff because it
	// produces a focused, line-by-line diff on failure (much nicer than printing
	// whole structs). If the values differ, - lines are the expected (want) value,
	// and + lines are the actual (got) value.

		// t.Error marks the test as failed but continues; t.Fatalf would stop immediately.
		// Either is fine here—Error is used so the diff shows in output without aborting the test file.
		