// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

import (
	"github.com/docker/docker/pkg/namesgenerator"
)

type User struct {
	UserID      string
	RequestCode string
}

func NewUser(UserID string) ModelInterface {
	instance := new(User)
	instance.UserID = UserID
	instance.RequestCode = namesgenerator.GetRandomName(0)
	return instance
}

func (u User) Load(filter interface{}) error {
	return nil
}

func (u User) Reload() error {
	return nil
}

func (u User) Create() error {
	return nil
}

func (u User) Update() error {
	return nil
}

func (u User) Destroy() error {
	return nil
}
