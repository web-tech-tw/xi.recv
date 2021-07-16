// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

type Message struct{}

func NewMessage() ModelInterface {
	instance := new(Message)
	return instance
}

func (m Message) load() ModelInterface {
	return m
}

func (m Message) reload() ModelInterface {
	return m
}

func (m Message) create() bool {
	return false
}

func (m Message) update() bool {
	return false
}

func (m Message) destroy() bool {
	return false
}
