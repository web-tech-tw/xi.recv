// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

type User struct{}

func NewUser() ModelInterface {
	instance := new(User)
	return instance
}

func (u User) load() ModelInterface {
	return u
}

func (u User) reload() ModelInterface {
	return u
}

func (u User) create() bool {
	return false
}

func (u User) update() bool {
	return false
}

func (u User) destroy() bool {
	return false
}
