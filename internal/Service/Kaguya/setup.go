// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Kaguya

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal"
	"github.com/star-inc/xi.recv/internal/Middleware/Web/Access"
	"github.com/star-inc/xi.recv/internal/Service/Kaguya/box"
	"github.com/star-inc/xi.recv/internal/Service/Kaguya/talk"
)

type Kaguya struct{}

func NewKaguya() internal.Router {
	instance := new(Kaguya)
	return instance
}

func (w *Kaguya) RouterSetup(router *gin.Engine) {
	router.Use(Access.Access)
	// Websocket Channels
	router.GET("/message/:relation_id", talk.Service)
	router.GET("/messagebox/:client_id", box.Service)
}
