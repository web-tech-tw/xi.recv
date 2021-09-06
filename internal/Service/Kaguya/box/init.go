// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package box

import (
	"github.com/star-inc/xi.recv/internal/Config"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
	KernelSource "gopkg.in/star-inc/kaguyakernel.v2/source"
	"log"
)

var (
	containerSource  *KernelSource.ContainerSource
	messageboxSource *KernelSource.MessageboxSource
)

func init() {
	source, err := KernelSource.NewContainerSource(
		rethinkdb.ConnectOpts{Address: Config.RethinkdbAddr},
		"message",
	)
	if err != nil {
		log.Panicln(err)
	}
	containerSource = source.(*KernelSource.ContainerSource)
}

func init() {
	source, err := KernelSource.NewMessageboxSource(
		rethinkdb.ConnectOpts{Address: Config.RethinkdbAddr},
		"messagebox",
	)
	if err != nil {
		log.Panicln(err)
	}
	messageboxSource = source.(*KernelSource.MessageboxSource)
}
