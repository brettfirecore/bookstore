// Declare that this file belongs to the external test package "bookstore_test".
// Using the _test suffix means our tests import the bookstore package the same
// way real users would. This helps us avoid accidentally touching unexported
// details inside the package we’re testing.
package bookstore_test

import (
	"testing" // Go's built-in testing framework

	"github.com/google/go-cmp/cmp" // Handy library to print readable diffs

	// We import the package under test. Giving it the local name "bookstore"
	// makes it clear where types and functions come from in this file.
	bookstore "github.com/brettfirecore/bookstore"
)

// TestGetBook checks the "happy path": asking for an ID that exists
// should return the correct Book and no error.
func TestGetBook(t *testing.T) {
	// t.Parallel marks this test as safe to run alongside others.
	// Go's test runner can then schedule multiple tests at once to speed things up.
	t.Parallel()

	// Arrange: build a tiny "catalog" (a map from ID -> Book) to use as test data.
	// This lives only for the duration of the test.
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
		3: {ID: 3, Title: "Old title"},
	}
	catalog[3] = bookstore.Book{ID: 3, Title: "Spark Joy"}

	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got)) // Helpful diff showing which fields differ
	}
}
