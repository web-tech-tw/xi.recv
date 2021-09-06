// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Access

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal/Data"
	"log"
)

// Access - the middleware to get access to connect talk service.
func Access(c *gin.Context) {
	idToken := c.GetHeader("xi-it")
	status, profile := getProfileFromLINE(idToken)
	if resp, ok := profile.(successResponse); status == 200 && ok {
		user := Data.NewUser(resp.Sub)
		err := user.Reload()
		if err != nil {
			panic(err)
		}
		c.Set("user", user)
	} else if resp, ok := profile.(errorResponse); ok {
		log.Panicf(
			"Status Code: %d\nError: %s\nDescription: %s",
			status,
			resp.Error,
			resp.ErrorDescription,
		)
	} else {
		message := "Unexpected response from LINE"
		log.Panicf("Status: %d\nError: %s", status, message)
	}
}
