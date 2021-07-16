// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Models

type Model interface {
	load() Model
	reload() Model
	create() bool
	update() bool
	destroy() bool
}
