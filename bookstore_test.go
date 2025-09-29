// Declare that this file belongs to the "bookstore_test" package.
// Keeping tests in a separate *_test package means they can only use the
// public API of the bookstore package (just like external users would).
package bookstore_test

import (
	"testing" // Provides T (testing.T) and helpers like t.Run/t.Parallel/t.Fatal.

	// Import the package under test by its module path so we can use exported names:
	//   bookstore.Book, bookstore.Buy, etc.
	bookstore "github.com/brettfirecore/bookstore"
)

// TestBook verifies we can construct a public Book value.
// It’s a lightweight “does this public type compile and instantiate” check.
func TestBook(t *testing.T) {
	t.Parallel() // Allow this test to run alongside others to speed up the suite.

	// Construct a Book. Assigning to the blank identifier (_) avoids an “unused variable” error.
	// This test doesn’t need to read the value; it just ensures the type is usable.
	_ = bookstore.Book{
		Title:  "Spark Joy",  // Example title
		Author: "Marie Kondō", // Example author (note: public/exported field)
		Copies: 2,            // Arbitrary example inventory
	}
}

// TestBuy checks the happy path: with available stock, Buy should decrement Copies by one
// and return no error.
func TestBuy(t *testing.T) {
	t.Parallel() // Safe because this test doesn’t touch shared state.

	// Arrange: start with two copies so one purchase is valid.
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 2,
	}

	// We expect the returned Book to have one fewer copy.
	want := 1

	// Act: call Buy. It returns an updated Book plus an error.
	result, err := bookstore.Buy(b)
	if err != nil {
		// Fatal stops this test immediately; no point comparing values if Buy failed.
		t.Fatal(err)
	}

	// Extract the field we’re asserting on (clearer failure messages).
	got := result.Copies

	// Assert: one copy was decremented.
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	// (Optional defensive check — uncomment to verify pass-by-value behavior:)
	// The input 'b' should remain unchanged because Buy takes/returns values, not pointers.
	// if b.Copies != 2 {
	// 	t.Errorf("original Book mutated: want Copies=2, got %d", b.Copies)
	// }
}

// TestBuyErrorsIfNoCopiesLeft ensures the error branch executes when stock is zero.
// This increases coverage and protects behavior for the “sold out” case.
func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()

	// Arrange: zero stock should cause Buy to fail.
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 0,
	}

	// Act: attempt to buy with no stock.
	_, err := bookstore.Buy(b)

	// Assert: we expect a non-nil error.
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}

	// (Optional enhancement — if the package exposes a sentinel error like ErrNoCopies,
	// you can assert precisely using errors.Is. Example:
	//
	//   _, err := bookstore.Buy(b)
	//   if !errors.Is(err, bookstore.ErrNoCopies) {
	//       t.Fatalf("want ErrNoCopies, got %v", err)
	//   }
	//
	// This helps catch error-message changes while preserving behavior.)
}
