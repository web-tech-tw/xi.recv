// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package talk

import (
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
	"gopkg.in/star-inc/kaguyakernel.v2/source"
)

type Middlewares struct {
	RelationID       string
	Access           AuthorizeInterface
	MessageboxSource *source.MessageboxSource
}

func newMiddlewares(
	RelationID string,
	Access AuthorizeInterface,
	MessageboxSource *source.MessageboxSource,
) Kernel.MiddlewareInterface {
	instance := new(Middlewares)
	instance.RelationID = RelationID
	instance.Access = Access
	instance.MessageboxSource = MessageboxSource
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
