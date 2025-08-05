// oink
package bookstore

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

// Broken on purpose: returns a map instead of a slice
func GetAllBooks(catalog map[int]Book) []Book {
	return catalog // ❌ compile-time type error: map[int]Book ≠ []Book
}

