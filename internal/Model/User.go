// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Model

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/go-redis/redis/v8"
)

type User struct {
	source      *redis.Client
	UserID      string
	RequestCode string
}

func NewUser(UserID string) Interface {
	instance := new(User)
	instance.UserID = UserID
	instance.RequestCode = namesgenerator.GetRandomName(0)
	return instance
}

func (u *User) Load(filter interface{}) error {
	ctx := context.Background()
	values := u.source.Get(ctx, filter.(string))
	data, err := values.Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Reload() error {
	return u.Load(u.UserID)
}

func (u *User) Create() error {
	return nil
}

func (u *User) Update() error {
	return nil
}

func (u *User) Destroy() error {
	return nil
}
