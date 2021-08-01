// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Model

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Message struct {
	source *redis.Client
	UUID   string
	From   *User
	To     *Room
	Text   string
}

func NewMessage(from *User, to *Room, text string) Interface {
	instance := new(Message)
	instance.UUID = uuid.New().String()
	instance.From = from
	instance.To = to
	instance.Text = text
	return instance
}

func (m *Message) Load(filter interface{}) error {
	ctx := context.Background()
	values := m.source.Get(ctx, filter.(string))
	data, err := values.Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Message) Reload() error {
	return m.Load(m.UUID)
}

func (m *Message) Create() error {
	return nil
}

func (m *Message) Update() error {
	return nil
}

func (m *Message) Destroy() error {
	return nil
}
