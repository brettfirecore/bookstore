// Package bookstore defines types and functions related to managing a collection of books.
// ftl-21.2
package bookstore

// NOTE: takes int (not *int) on purpose.
func Double(input *int) {
	// no-op; body doesn't matter for the error
	*input *= 2
}
