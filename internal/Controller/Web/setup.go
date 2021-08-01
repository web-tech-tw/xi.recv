package Web

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal/Controller"
)

type Web struct{}

func NewWeb() Controller.Interface {
	instance := new(Web)
	return instance
}

func (w Web) RouterSetup(router *gin.Engine) {
	router.Use(Access)
	router.POST("/user")
	router.POST("/room")
	router.POST("/request")
	router.POST("/message")
}
