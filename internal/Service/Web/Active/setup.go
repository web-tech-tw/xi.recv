// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Active

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal"
	"github.com/star-inc/xi.recv/internal/Middleware/Web/Access"
)

type Web struct{}

func NewWeb() internal.Router {
	instance := new(Web)
	return instance
}

func (w *Web) RouterSetup(router *gin.Engine) {
	router.Use(Access.Access)
}
