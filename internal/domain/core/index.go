package core

import "github.com/Akanibekuly/tsarka-tz/internal/interfaces"

type St struct {
	lg interfaces.Logger
}

func New(lg interfaces.Logger) *St {
	return &St{
		lg: lg,
	}
}
