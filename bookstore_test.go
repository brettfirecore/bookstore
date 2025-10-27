// Declare that this file belongs to the external test package "bookstore_test".
// Using the _test suffix means we import the bookstore package just like any
package bookstore_test

import (
	//	"sort"    // (example of other test helpers you might use later)
	"testing" // Go's built-in testing tools: gives us *testing.T and helpers

//	"github.com/google/go-cmp/cmp" // Handy diff lib for comparing complex values

	bookstore "github.com/brettfirecore/bookstore"
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
		t.Fatal("want error for invalid catagory, got nil")
	}
}
