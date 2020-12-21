package types

import (
	"fmt"
)

type Data struct{}
type IData interface{ Func() }

func (e Data) Func() { // receiver is a struct
	fmt.Println("func")
}

func main() {
	var p *Data
	Do(p)
}

func Do(data IData) {
	if data != nil { // it goes here
		data.Func() // panics
	} else {
		fmt.Println("no data")
	}
}
