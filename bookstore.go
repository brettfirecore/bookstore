// Package bookstore defines types and functions related to managing a collection of books.
// ftl-21.2
package bookstore

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

func (b *Book) SetPriceCents(price int) error {
	b.PriceCents = price // nope
	return nil
}
