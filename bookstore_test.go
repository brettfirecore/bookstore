// bookstore_test is the external test package for the bookstore module.
// Using `package bookstore_test` means we’re testing the public API of the bookstore package,
// just like any code outside the module would — this is called “black-box” testing.
// bookstore_test.go
// bookstore_test.go
package bookstore_test

import (
	"testing"

	"github.com/brettfirecore/bookstore"
	"github.com/google/go-cmp/cmp"
)

// Valid ID → found
func TestCatalog_GetBook_Found(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	// NEW: call the *method* on Catalog, and expect (Book, bool)
	got, ok := catalog.GetBook(2)
	if !ok {
		t.Fatalf("GetBook(2): want ok=true")
	}

	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	if !cmp.Equal(want, got) {
		t.Fatalf("GetBook(2) mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

// Invalid ID → not found
func TestCatalog_GetBook_NotFound(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
	}

	got, ok := catalog.GetBook(99) // NEW: method call + comma-ok
	if ok {
		t.Fatalf("GetBook(99): want ok=false, got ok=true with %+v", got)
	}
}
