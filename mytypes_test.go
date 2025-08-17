// ftl-code / 21.2 mytypes_test.go in bookstore is TestDouble.

package bookstore_test

import (
	"testing"

	"github.com/brettfirecore/bookstore"
)

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}
