// Package bookstore defines types and functions related to managing a collection of books.
// ftl-21.2
package bookstore

import "fmt"

const (
	CategoryAutobiography     = "Autobiography"
	CategoryLargePrintRomance = "Large Print Romance"
	CategoryParticlePhysics   = "Particle Physics"
)

var validCategory = map[string]bool{
	"Autobiography":       true,
	"Large Print Romance": true,
	"Particle Physics":    true,
}

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
