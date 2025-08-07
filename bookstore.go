// Package bookstore defines types and functions related to managing a collection of books. 
// This is the top-level package declaration. Every Go file must begin with a `package` line.
package bookstore

// Book represents a book with a price and optional discount.
// It's a struct — a custom data type made of named fields.
type Book struct {
	Title           string // The title of the book (e.g., "For the Love of Go")
	PriceCents      int    // The price in cents (e.g., 4000 = $40.00)
	DiscountPercent int    // The discount percentage (e.g., 25 = 25% off)
}

// NetPriceCents returns the book's price after applying the discount.
// It takes a Book as input and returns an int (final price in cents).
func NetPriceCents(b Book) int {
	// Compute the amount of money to subtract as a discount
	// Multiply first to avoid integer division errors (like 25/100 = 0)
	saving := b.PriceCents * b.DiscountPercent / 100

	// Subtract the discount from the original price to get the net price
	return b.PriceCents - saving
}
