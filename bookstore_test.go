package bookstore_test

import (
	"testing"

	"bookstore"
)

func TestBookFields(t *testing.T) {
	t.Parallel()

	// We expect a public Book type with public fields Title, Author, Copies.
	b := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundel",
		Copies: 99,
	}

	if b.Title == "" {
		t.Fatal("Title should be set")
	}
	if b.Author == "" {
		t.Fatal("Author should be set")
	}
	if b.Copies != 99 {
		t.Fatalf("Copies: want 99, got %d", b.Copies)
	}
}
