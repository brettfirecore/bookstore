// bookstore_test is the external test package for the bookstore module.
// Using `package bookstore_test` means we’re testing the public API of the bookstore package,
// just like any code outside the module would — this is called “black-box” testing.
package bookstore_test // (Listing 17.3)

import (
	"testing"    // Go’s standard library testing package

	"github.com/brettfirecore/bookstore"
)

// TestNetPriceCents verifies that NetPriceCents correctly calculates the final price
// after applying a percentage discount.
func TestNetPriceCents(t *testing.T) {
	t.Parallel() // Allows this test to run in parallel with others, speeding up test execution.

	// Arrange: create a Book with a known price and discount
	b := bookstore.Book{
		Title:           "For the love of Go", // title is just metadata
		PriceCents:      4000,                 // base price = 4000 cents = $40.00
		DiscountPercent: 25,                   // 25% discount
	}

	// Expectation: 25% of 4000 = 1000 → 4000 - 1000 = 3000
	want := 3000 // this is the value we expect after discount is applied

	// Act: call the function under test
	got := bookstore.NetPriceCents(b)

	// Assert: check if the result is what we expected
	if want != got {
		// If not, fail the test and print the difference
		t.Errorf("want %d, got %d", want, got)
	}
}
