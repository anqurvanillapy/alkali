package reflect

import (
	"testing"
)

type Data struct {
	Key   string `form:"key"`
	Value int    `form:"value"`
}

func TestToQuery(t *testing.T) {
	data := &Data{
		"foo",
		42,
	}

	q := ToQuery(data)
	if v, ok := q["key"]; ok {
		if v[0] != "foo" {
			t.Fatal(v[0])
		}
	} else {
		t.Fatalf("Parameter 'key' not exists: %v", q)
	}

	if v, ok := q["value"]; ok {
		if v[0] != "42" {
			t.Fatal(v[0])
		}
	} else {
		t.Fatalf("Parameter 'value' not exists: %v", q)
	}
}
