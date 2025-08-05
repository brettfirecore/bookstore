// oink
package bookstore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := map[int]Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	got := GetAllBooks(catalog) // ← This line will fail if GetAllBooks returns map[int]Book
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

