// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Model

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Request struct {
	HashId      string
	From        *User
	To          *Room
	CreatedTime int64
}

func NewRequest(from *User, to *Room) ModelInterface {
	instance := new(Request)
	sha256Sum := sha256.Sum256([]byte(strings.Join([]string{from.UserID, to.UUID}, "\x1e")))
	instance.HashId = fmt.Sprintf("%x", sha256Sum)
	instance.From = from
	instance.To = to
	instance.CreatedTime = time.Now().Unix()
	return instance
}

func (r Request) Load(filter interface{}) error {
	return nil
}

func (r Request) Reload() error {
	return r.Load(r.HashId)
}

func (r Request) Create() error {
	return nil
}

func (r Request) Update() error {
	return nil
}

func (r Request) Destroy() error {
	return nil
}
