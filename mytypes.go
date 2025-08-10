// Package bookstore defines types and functions related to managing a collection of books.
package bookstore

import "strings"

// StringUppercaser wraps strings.Builder.
type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
