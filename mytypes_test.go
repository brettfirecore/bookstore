// ftl-code / 21.1 mytypes_test.go in bookstore is TestDouble.

package bookstore

import "testing"

func TestDouble(t *testing.T) {
	t.Parallel()

	x := 12
	want := 24

	got := Double(x)
	if got != want {
		t.Errorf("Double(%d) = %d; want %d", x, got, want)
	}

	// x is unchanged because Go passes by valoe
	if x != 12 {
		t.Errorf("x was mutated: got %d; wan 12", x)
	}
}
