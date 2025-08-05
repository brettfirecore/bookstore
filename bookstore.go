// Package bookstore oink is the name of a pig
package bookstore

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}
