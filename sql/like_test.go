package sql_test

import (
	"testing"

	"github.com/anqurvanillapy/alkali/sql"
)

func TestWildcardSurround(t *testing.T) {
	var val string

	val = "%"
	if sql.WildcardSurround(val) != "%\\%%" {
		t.Fatal(val)
	}

	val = "_"
	if sql.WildcardSurround(val) != "%\\_%" {
		t.Fatal(val)
	}

	// And use it like `tx.Where("name like ?", sql.WildcardSurround(name))`.
}
