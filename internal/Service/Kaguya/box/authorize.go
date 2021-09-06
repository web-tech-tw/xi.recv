// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package box

import (
	"github.com/gin-gonic/gin"
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
)

type Authorize struct {
	identity string
	context  *gin.Context
}

func getAccess(context *gin.Context) Kernel.AuthorizeInterface {
	authorize := new(Authorize)
	authorize.context = context
	return authorize
}

func (authorize *Authorize) Me() string {
	return authorize.identity
}

func (authorize *Authorize) Permission(userID string) bool {
	return userID == authorize.identity
}
