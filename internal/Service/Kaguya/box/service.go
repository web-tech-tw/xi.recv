// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package box

import (
	"github.com/gin-gonic/gin"
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
	"gopkg.in/star-inc/kaguyakernel.v2/data"
	"gopkg.in/star-inc/kaguyakernel.v2/service/box"
	"log"
	"net/http"
	"time"
)

func Service(c *gin.Context) {
	messageboxSource.ClientID = c.Param("client_id")
	salt := ""

	Access := getAccess(c)

	if !Access.Permission(messageboxSource.ClientID) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if !messageboxSource.CheckReady() {
		if err := messageboxSource.CreateTerm(); err != nil {
			log.Panicln(err)
		}
	}

	syncExtraDataAssigner := func(container data.SyncMessagebox) interface{} {
		// Add prefix
		container.LastSeen += int64(3 * time.Millisecond)
		containerSource.RelationID = container.Target
		unreadCount := data.CountContainersByTimestamp(containerSource, container.LastSeen)
		return map[string]interface{}{
			"unreadCount": unreadCount,
		}
	}

	middlewares := newMiddlewares()
	service := box.NewServiceInterface(messageboxSource, syncExtraDataAssigner)

	handler := Kernel.Run(service, Access, middlewares, salt)
	if err := handler.HandleRequest(c.Writer, c.Request); err != nil {
		log.Panicln(err)
	}
}
