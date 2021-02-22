package client_test

import (
	"time"

	"github.com/anqurvanillapy/alkali/patterns/client-options"
	"github.com/anqurvanillapy/alkali/patterns/client-options/options"
)

func ExampleNewClient() {
	c := client.NewClient(
		options.WithDebug(),
		options.WithTimeout(time.Second),
	)
	c.Run()
	// Output: debug
	// timeout: 1s
	// running
}
