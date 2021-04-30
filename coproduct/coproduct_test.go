package coproduct

import (
	"fmt"
	"reflect"
	"testing"
)

type I interface{ I() }

type Items []Item

func (Items) I() {}

type Item struct {
	Name string `json:"name"`
}

func (Item) I() {}

var types = []reflect.Type{
	reflect.TypeOf(Item{}),
	reflect.TypeOf(Items{}),
}

func TestDecodeUntaggedUnion(t *testing.T) {
	var i I

	d := []byte(`{"name":"foo"}`)
	if err := DecodeUntaggedUnion(d, &i, types...); err != nil {
		t.Fatal(err)
	}
	fmt.Println(i)

	d = []byte(`[{"name":"foo"},{"name":"bar"}]`)
	if err := DecodeUntaggedUnion(d, &i, types...); err != nil {
		t.Fatal(err)
	}
	fmt.Println(i)
}

type U struct {
	I    `json:"value"`
	Type string `json:"type"`
}

type A struct {
	A string `json:"a"`
}

func (A) I() {}

type B struct {
	B string `json:"b"`
}

func (B) I() {}

var typeMap = map[string]reflect.Type{
	"a": reflect.TypeOf(A{}),
	"b": reflect.TypeOf(B{}),
}

func TestDecodeTaggedUnion(t *testing.T) {
	var u U
	tf, _ := reflect.TypeOf(u).FieldByName("Type")
	vf, _ := reflect.TypeOf(u).FieldByName("I")

	d := []byte(`{"type":"a","value":{"a":"a"}}`)
	if err := DecodeTaggedUnion(d, tf, vf, &u, typeMap); err != nil {
		t.Fatal(err)
	}
	fmt.Println(u.Type, u.I)

	d = []byte(`{"type":"b","value":{"b":"b"}}`)
	if err := DecodeTaggedUnion(d, tf, vf, &u, typeMap); err != nil {
		t.Fatal(err)
	}
	fmt.Println(u.Type, u.I)
}
