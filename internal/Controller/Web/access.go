// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Web

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal/Model"
)

func Access(c *gin.Context) {
	identity := c.GetHeader("xi-jwt")
	user := Model.NewUser(identity)
	err := user.Reload()
	if err != nil {
		panic(err)
	}
	c.Set("user", user)
}
