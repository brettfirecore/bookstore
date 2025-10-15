// Package bookstore stores info about a book in a folder Bookstore.go
package bookstore

// Book stores info about a book. We keep prices in *cents* (integers)
// so we never run into floating-point rounding problems (e.g., 0.1+0.2 issues).
type Book struct {
	Title           string // The book’s title
	Author          string // The author’s name
	Copies          int    // Copies in stock
	ID              int    // Unique ID for lookups
	PriceCents      int    // Price in cents (e.g., $12.34 -> 1234)
	DiscountPercent int    // Discount as a whole number percent (0..100)
}

// NetPriceCents is a *method* on Book.
// The bit in front of the name — (b Book) — is called the *receiver*.
// Think of it like a hidden first parameter named "b" that holds the Book
// we’re calling the method on. Because it’s a *value receiver*, we get a copy
// of the Book and we DON’T change the original.
//
// Example call site:
//
//	b := Book{PriceCents: 2000, DiscountPercent: 25}
//	n := b.NetPriceCents()  // n == 1500
func (b Book) NetPriceCents() int {
	// Math, all in integers (cents), so it’s precise:
	//   net = price * (100 - discount) / 100
	// Note: integer division truncates toward zero (no rounding up).
	return b.PriceCents * (100 - b.DiscountPercent) / 100
}
