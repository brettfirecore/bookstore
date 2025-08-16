// ftl-code / 21.2 mytypes_test.go in bookstore is TestDouble.

package bookstore_test

import (
	"testing"

	mytypes "github.com/brettfirecore/bookstore"
)

func TestDouble(t *testing.T) {
	t.Parallel()

	x := 12
	mytypes.Double(&x) // <- passing *int to a func that expects int
}
