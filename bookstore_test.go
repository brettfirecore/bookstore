// ftl-code / 21.2 mytypes_test.go in bookstore is TestDouble.

package bookstore_test

import (
	"testing"

	"github.com/brettfirecore/bookstore"
)

func TestSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()
	if want != got {
		t.Errorf("want category %q, got %q", want, got)
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory("bogus")
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}
