// Package bookstore defines types and functions related to managing a collection of books.
package bookstore

// MyBuilder is a custom type for experimenting
type MyBuilder struct{}

// Hello returns a fixed greeting.
func (MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
