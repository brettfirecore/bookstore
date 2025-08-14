// ftl-code/20.1/mytypes_test.go in bookstore is the external test package for the bookstore module.

package bookstore

import "testing"

func TestStringUppercaser(t *testing.T) {
	t.Parallel()

	var su StringUppercaser
	su.Contents.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"

	got := su.ToUpper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
