// bookstore_test is the external test package for the bookstore module.
// Using `package bookstore_test` allows us to test the public API like an outside user would.
package bookstore_test // (Listing 17.3)

import (
	"testing"

	"bookstore" // Import the bookstore package being teste                   // Used to sort slices before comparison
)

// TestGetAllBooks verifies that the GetAllBooks function returns the expected list of books.
func TestNetPriceCents(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others to speed up testing
	b := bookstore.Book{
		Title:           "For the love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	want := 3000
	got := bookstore.NetPriceCents(b)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
