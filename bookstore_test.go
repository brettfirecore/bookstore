// bookstore_test is the external test package for the bookstore module.
// Using `package bookstore_test` means we’re testing the public API of the bookstore package,
// just like any code outside the module would — this is called “black-box” testing.
// bookstore_test.go
package bookstore_test

import (
	"sort"
	"testing"

	"github.com/brettfirecore/bookstore" // import the package under test
	"github.com/google/go-cmp/cmp"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	// NEW: use the named type `Catalog` instead of a raw map[int]Book.
	// A Catalog literal looks like a map literal, because Catalog’s underlying type is a map.
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	// Expected slice of books. (Order doesn’t matter; we’ll sort `got` before comparing.)
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	// NEW: call the *method* on the catalog value instead of a free function.
	got := catalog.GetAllBooks()

	// Map iteration order is undefined, so sort by ID to make comparison deterministic.
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	// Compare the two slices; print a diff if they differ.
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
