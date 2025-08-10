// bookstore_test is the external test package for the bookstore module.

package bookstore

import (
	"strings"
	"testing"
)

func TestStringsBuilder(t *testing.T) {
	t.Parallel()

	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")

	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", sb.String(), wantLen, gotLen)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()

	var mb MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
