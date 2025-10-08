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
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got, err := bookstore.GetBook(catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}
