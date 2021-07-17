package Model

import "github.com/google/uuid"

type Room struct {
	UUID     string
	Creator  *User
	Requests []*Request
}

func NewRoom(creator *User) ModelInterface {
	instance := new(Room)
	instance.UUID = uuid.New().String()
	instance.Creator = creator
	instance.Requests = make([]*Request, 0)
	return instance
}

func (r Room) Load(filter interface{}) error {
	return nil
}

func (r Room) Reload() error {
	return r.Load(r.UUID)
}

func (r Room) Create() error {
	return nil
}

func (r Room) Update() error {
	return nil
}

func (r Room) Destroy() error {
	return nil
}
