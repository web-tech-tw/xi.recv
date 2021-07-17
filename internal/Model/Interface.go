// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Model

type ModelInterface interface {
	Load(filter interface{}) error
	Reload() error
	Create() error
	Update() error
	Destroy() error
}