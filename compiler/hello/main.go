// hello to Go compiler, compiled in Go version go1.15.2 darwin/amd64.
package main

import (
	"fmt"
)

// There is a very stupid bug in go1.15, if you simply dump the SSAs of `main`,
// the HTML writer will crash:
// https://github.com/golang/go/issues/41584
func hello() {
	fmt.Println("Hello, Go compiler!")
}

func main() {
	hello()
}
