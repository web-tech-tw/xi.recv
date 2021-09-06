// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package talk

import (
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
	KernelSource "gopkg.in/star-inc/kaguyakernel.v2/source"
	"log"
	"os"
)

var (
	containerSource  *KernelSource.ContainerSource
	messageboxSource *KernelSource.MessageboxSource
)

func init() {
	source, err := KernelSource.NewContainerSource(
		rethinkdb.ConnectOpts{Address: os.Getenv("RETHINKDB_ADDR")},
		"message",
	)
	if err != nil {
		log.Panicln(err)
	}
	containerSource = source.(*KernelSource.ContainerSource)
}

func init() {
	source, err := KernelSource.NewMessageboxSource(
		rethinkdb.ConnectOpts{Address: os.Getenv("RETHINKDB_ADDR")},
		"messagebox",
	)
	if err != nil {
		log.Panicln(err)
	}
	messageboxSource = source.(*KernelSource.MessageboxSource)
}
