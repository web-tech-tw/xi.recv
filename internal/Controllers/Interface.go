// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Controllers

import "github.com/gin-gonic/gin"

type Interface interface {
	RouterSetup(router *gin.Engine)
}
