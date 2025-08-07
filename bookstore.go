// Package bookstore defines types and functions related to managing a collection of books.
package bookstore

// Book represents a book with a price and optional discount.
type Book struct {
	Title           string
	PriceCents      int
	DiscountPercent int
}

// NetPriceCents returns the book's price after applying the discount.
func NetPriceCents(b Book) int {
	discount := b.PriceCents * b.DiscountPercent / 100

	return b.PriceCents - discount
}
