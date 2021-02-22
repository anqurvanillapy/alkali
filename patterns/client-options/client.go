package client

import (
	"fmt"

	"github.com/anqurvanillapy/alkali/patterns/client-options/internal"
	"github.com/anqurvanillapy/alkali/patterns/client-options/options"
)

type Client struct {
	settings *internal.DialSettings
}

func NewClient(opts ...options.Option) (client *Client) {
	client = &Client{settings: &internal.DialSettings{}}
	for _, opt := range opts {
		opt.Apply(client.settings)
	}
	return
}

func (c *Client) Run() {
	if c.settings.Debug {
		fmt.Println("debug")
	}
	if timeout := c.settings.Timeout; timeout > 0 {
		fmt.Println("timeout:", timeout)
	}
	fmt.Println("running")
}
