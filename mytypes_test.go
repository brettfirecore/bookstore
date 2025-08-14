// ftl-code / 21.1 mytypes_test.go in bookstore is TestDouble.

package bookstore

import "testing"

func TestDouble(t *testing.T) {
	t.Parallel()

	var x int = 12
	want := 24

	Double(x)

	if want != x {
		t.Errorf("want %d, want %d", want, x)
	}
}
