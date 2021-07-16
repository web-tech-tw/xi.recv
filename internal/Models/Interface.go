// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

type ModelInterface interface {
	load() ModelInterface
	reload() ModelInterface
	create() bool
	update() bool
	destroy() bool
}
