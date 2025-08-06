// bookstore_test is the external test package for the bookstore module.
// Using `package bookstore_test` allows us to test the public API like an outside user would.
package bookstore_test // (Listing 17.3)

import (
	"bookstore"                // Import the bookstore package being tested
	"sort"                     // Used to sort slices before comparison
	"testing"                 // Go's built-in testing package

	"github.com/google/go-cmp/cmp" // Used to compare structs and print readable diffs
)

// TestGetAllBooks verifies that the GetAllBooks function returns the expected list of books.
func TestGetAllBooks(t *testing.T) {
	t.Parallel() // Allow this test to run in parallel with others to speed up testing

	// Arrange: Set up a sample catalog (map[int]Book).
	// The keys are book IDs; the values are Book structs.
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	// Define the expected result as a slice of Book structs.
	// Since maps don't preserve order, we'll sort both slices before comparison.
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	// Act: Call the function under test.
	got := bookstore.GetAllBooks(catalog)

	// Sort the resulting slice by ID to ensure consistent ordering for comparison.
	// This is necessary because map iteration order is random in Go.
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	// Assert: Compare the expected and actual result using cmp.Equal.
	// If they differ, print a unified diff to help diagnose the problem.
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got)) // Show only what's different between the slices
	}
}
