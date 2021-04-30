package checkiface

// HasFoo checks if an object has method `Foo`.
func HasFoo(v interface{}) bool {
	_, ok := v.(interface {
		Foo()
	})
	return ok
}
