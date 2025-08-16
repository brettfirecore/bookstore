// ftl-code / 21.2 mytypes_test.go in bookstore is TestDouble.

package bookstore_test

import (
	"testing"

	mytypes "github.com/brettfirecore/bookstore"
)

func TestDouble(t *testing.T) {
	t.Parallel()
	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)
	p := &x
	p.Double()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}
