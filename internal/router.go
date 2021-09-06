// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package internal

import "github.com/gin-gonic/gin"

type Router interface {
	RouterSetup(router *gin.Engine)
}
