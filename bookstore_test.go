// Declare that this file belongs to the "bookstore_test" package.
// This convention keeps the tests in a separate package from the code under test,
// so the tests can only access the public API of the bookstore package (just like a user would).
package bookstore_test

import (
	"testing" // The testing package provides the tools to write and run tests.

	// Import the code we want to test. We use its module path so the test can call its exported names.
	"github.com/brettfirecore/bookstore"
)

// TestBook checks that we can create a Book value without errors.
// It doesn’t actually assert anything yet; it’s just making sure the struct compiles and can be used.
func TestBook(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others.

	// Create a Book value with some fields set.
	// Using the blank identifier (_) so the variable is not used — we just ensure we can construct it.
	_ = bookstore.Book{
		Title:  "For the Love of Go", // Title field set
		Author: "John Arundel",       // Author field set
		Copies: 2,                    // Copies field set
	}
}

// TestBuy describes the behaviour we want: buying a book should reduce the Copies by one.
func TestBuy(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others.

	// Set up our starting Book with 2 copies available.
	b := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundel",
		Copies: 2,
	}

	// Define the expected result after buying one copy.
	want := 1

	// Call the (not yet implemented) Buy function.
	// This should return a Book with one fewer copy.
	result := bookstore.Buy(b)

	// Access the Copies field of the returned Book.
	got := result.Copies

	// Check if the result matches what we wanted.
	// If not, fail the test and print a clear message showing both values.
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
