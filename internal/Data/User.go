// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Data

import (
	"github.com/docker/docker/pkg/namesgenerator"
	XISource "github.com/star-inc/xi.recv/internal/Source"
	KernelSource "gopkg.in/star-inc/kaguyakernel.v2/source"
)

type User struct {
	UserID      string
	RequestCode string
}

func NewUser(UserID string) Interface {
	instance := new(User)
	instance.UserID = UserID
	instance.RequestCode = namesgenerator.GetRandomName(0)
	return instance
}

func (u *User) Load(_ KernelSource.Interface, filter interface{}) error {
	return nil
}

func (u *User) Reload(source KernelSource.Interface) error {
	return u.Load(source, u.UserID)
}

func (u *User) Create(source KernelSource.Interface) error {
	sourceInstance := source.(*XISource.UserSource)
	return source.GetTerm().Table(sourceInstance.RelationID).Insert(u).Exec(source.GetSession())
}

func (u *User) Update(source KernelSource.Interface) error {
	sourceInstance := source.(*XISource.UserSource)
	return source.GetTerm().Table(sourceInstance.RelationID).Replace(u).Exec(source.GetSession())
}

func (u *User) Destroy(source KernelSource.Interface) error {
	sourceInstance := source.(*XISource.UserSource)
	return source.GetTerm().Table(sourceInstance.ClientID).Get(u.Target).Delete().Exec(source.GetSession())
}
