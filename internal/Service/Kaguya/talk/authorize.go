// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package talk

import (
	"github.com/gin-gonic/gin"
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
)

type Authorize struct {
	identity string
	context  *gin.Context
}

type AuthorizeInterface interface {
	Kernel.AuthorizeInterface
	RelatedMessageBox(hashID string) []string
}

func getAccess(context *gin.Context) AuthorizeInterface {
	authorize := new(Authorize)
	authorize.context = context
	return authorize
}

func (authorize *Authorize) Me() string {
	return authorize.identity
}

func (authorize *Authorize) Permission(chatRoomID string) bool {
	return false
}

func (authorize *Authorize) RelatedMessageBox(hashID string) []string {
	return nil
}
