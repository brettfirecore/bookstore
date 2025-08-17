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
	err := b.SetPriceCents(-1)
	if err != nil {
		t.Fatal("wanr error setting invalid price -1, got nil")
	}
}
