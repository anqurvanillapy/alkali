package options

import (
	"time"

	"github.com/anqurvanillapy/alkali/patterns/client-options/internal"
)

type Option interface {
	Apply(settings *internal.DialSettings)
}

type withDebug struct{}

func (withDebug) Apply(settings *internal.DialSettings) {
	settings.Debug = true
}

func WithDebug() Option {
	return &withDebug{}
}

type withTimeout time.Duration

func (t withTimeout) Apply(settings *internal.DialSettings) {
	settings.Timeout = time.Duration(t)
}

func WithTimeout(t time.Duration) Option {
	return withTimeout(t)
}
