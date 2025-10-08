// Declare that this file belongs to the "bookstore_test" package (note the _test suffix).
package bookstore_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	bookstore "github.com/brettfirecore/bookstore"
)

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	got, err := bookstore.GetBook(catalog, 2)

	if err != nil {
		t.Fatal(err)
	}

	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
