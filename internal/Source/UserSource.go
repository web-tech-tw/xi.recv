// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Source

import (
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
	KernelSource "gopkg.in/star-inc/kaguyakernel.v2/source"
)

type UserSource struct {
	KernelSource.Base
	RelationID string
}

func NewUserSource(config rethinkdb.ConnectOpts, databaseName string) (KernelSource.Interface, error) {
	instance := new(UserSource)
	return instance, nil
}

func (u UserSource) CheckReady() bool {
	panic("implement me")
}

func (u UserSource) checkTerm() bool {
	panic("implement me")
}

func (u UserSource) CreateTerm() error {
	panic("implement me")
}

func (u UserSource) DropTerm() error {
	panic("implement me")
}
