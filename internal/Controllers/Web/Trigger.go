package Web

import (
	"github.com/star-inc/xi.recv/internal/Controllers"
)

type Web struct{}

func NewWeb() Controllers.Interface {
	instance := new(Web)
	return instance
}

func (w Web) Trigger() {
}
