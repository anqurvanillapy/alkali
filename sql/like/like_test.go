package like_test

import (
	"testing"

	"github.com/anqurvanillapy/alkali/sql/like"
)

func TestWildcardSurround(t *testing.T) {
	var val string

	val = "%"
	if like.WildcardSurround(val) != "%\\%%" {
		t.Fatal(val)
	}

	val = "_"
	if like.WildcardSurround(val) != "%\\_%" {
		t.Fatal(val)
	}

	// And use it like `tx.Where("name like ?", sql.WildcardSurround(name))`.
}
