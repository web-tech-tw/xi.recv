// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Data

import KernelSource "gopkg.in/star-inc/kaguyakernel.v2/source"

type Interface interface {
	Load(source KernelSource.Interface, filter interface{}) error
	Reload(source KernelSource.Interface) error
	Create(source KernelSource.Interface) error
	Update(source KernelSource.Interface) error
	Destroy(source KernelSource.Interface) error
}
