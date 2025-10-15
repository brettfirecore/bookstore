// Declare that this file belongs to the external test package "bookstore_test".
// Using the _test suffix means we import the bookstore package just like any
// other program would. This keeps our tests honest: we only use the public API
// (no reaching into unexported stuff inside package bookstore).
package bookstore_test

import (
	// "sort"                     // (example of other test helpers you might use later)
	"testing" // Go's built-in testing tools: gives us *testing.T and helpers

	// "github.com/google/go-cmp/cmp" // Handy diff lib for comparing complex values

	// The package we're testing. We import it by module path and give it a short
	// local name "bookstore" so that, below, it's obvious where types/functions
	// come from (bookstore.Book, bookstore.SomeFunc, etc).
	bookstore "github.com/brettfirecore/bookstore"
)

// TestNetPriceCents checks the method Book.NetPriceCents.
// Goal: given a Book with a price and a discount, the method should return
// the correct *final* price in cents using integer math.
// Formula we're expecting: price * (100 - discount) / 100
func TestNetPriceCents(t *testing.T) {
	t.Parallel() // Allow this test to run alongside others for speed.

	// Arrange: build a sample Book value to test with.
	// We keep money in cents (integers) to avoid float precision problems.
	// 4000 cents == $40.00; 25% discount means we expect $30.00 (3000 cents).
	b := bookstore.Book{
		Title:           "For the Love of Go", // just metadata; doesn't affect price
		PriceCents:      4000,                 // original price in cents
		DiscountPercent: 25,                   // percent off (whole number)
	}

	// This is the exact number we expect from the method.
	// 4000 * (100 - 25) / 100 == 4000 * 75 / 100 == 3000.
	want := 3000

	// Act: call the *method* on our Book value.
	// Because NetPriceCents is defined with a value receiver (b Book),
	// calling it does not modify the original Book; it just returns a number.
	got := b.NetPriceCents()

	// Assert: if the result is not what we expected, fail the test with a helpful message.
	// We use %d because these are integers (cents). If this fails, it means either:
	// 1) the method's formula is wrong, or
	// 2) the inputs/expectations in the test don't match the intended behavior.
	if want != got {
		t.Errorf("NetPriceCents(%+v): want %d, got %d", b, want, got)
		// %+v prints the struct with field names, which is great for debugging.
	}

	// Tip for future cases:
	// - Add more scenarios (0% discount, 50% on odd cents like 1999, 100% discount, etc.)
	// - If you later clamp discounts to [0,100], add tests that document that behavior too.
}
