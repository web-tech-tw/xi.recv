// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package talk

import (
	"github.com/gin-gonic/gin"
	Kernel "gopkg.in/star-inc/kaguyakernel.v2"
	"gopkg.in/star-inc/kaguyakernel.v2/service/talk"
	"log"
	"net/http"
	"unicode/utf8"
)

func Service(c *gin.Context) {
	relationID := c.Param("relation_id")
	containerSource.RelationID = relationID
	salt := ""

	Access := getAccess(c)

	if !Access.Permission(relationID) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if !containerSource.CheckReady() {
		if err := containerSource.CreateTerm(); err != nil {
			log.Panicln(err)
		}
	}

	contentValidator := func(contentType int, content string) bool {
		if contentType == 0 {
			if utf8.RuneCountInString(content) <= 2000 {
				return true
			}
		}
		return false
	}

	middlewares := newMiddlewares(relationID, Access, messageboxSource)
	service := talk.NewServiceInterface(containerSource, contentValidator)

	handler := Kernel.Run(service, Access, middlewares, salt)
	if err := handler.HandleRequest(c.Writer, c.Request); err != nil {
		log.Panicln(err)
	}
}
