package types

import "testing"

type A struct{}
type B struct{}

func (B) Foo() {}

func TestHasFoo(t *testing.T) {
	a := A{}
	b := B{}
	if HasFoo(a) {
		t.Fatal(a)
	}
	if !HasFoo(b) {
		t.Fatal(b)
	}
}
