// Package bookstore defines types and functions related to managing a collection of books.
// ftl-21.2
package bookstore

type MyInt int

func (input *MyInt) Double() {

	*input *= 2
}
