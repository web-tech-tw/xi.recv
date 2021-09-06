// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package box

import (
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
)

type Middlewares struct{}

func newMiddlewares() Kernel.MiddlewareInterface {
	instance := new(Middlewares)
	return instance
}

func (m Middlewares) OnRequestBefore() Kernel.OnRequestBefore {
	return nil
}

func (m Middlewares) OnRequestAfter() Kernel.OnRequestAfter {
	return nil
}

func (m Middlewares) OnResponseBefore() Kernel.OnResponseBefore {
	return nil
}

func (m Middlewares) OnResponseAfter() Kernel.OnResponseAfter {
	return nil
}
