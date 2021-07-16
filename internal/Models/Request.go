// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

type Request struct{}

func NewRequest() ModelInterface {
	instance := new(Request)
	return instance
}

func (r Request) load() ModelInterface {
	return r
}

func (r Request) reload() ModelInterface {
	return r
}

func (r Request) create() bool {
	return false
}

func (r Request) update() bool {
	return false
}

func (r Request) destroy() bool {
	return false
}
