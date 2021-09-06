// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Data

import "github.com/google/uuid"

type Room struct {
	UUID     string
	Creator  *User
	Requests []*Request
}

func NewRoom(creator *User) Interface {
	instance := new(Room)
	instance.UUID = uuid.New().String()
	instance.Creator = creator
	instance.Requests = make([]*Request, 0)
	return instance
}

func (r *Room) Load(filter interface{}) error {
	return nil
}

func (r *Room) Reload() error {
	return r.Load(r.UUID)
}

func (r *Room) Create() error {
	return nil
}

func (r *Room) Update() error {
	return nil
}

func (r *Room) Destroy() error {
	return nil
}
