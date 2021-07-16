// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal/Controllers"
	"github.com/star-inc/xi.recv/internal/Controllers/Bot"
	"github.com/star-inc/xi.recv/internal/Controllers/Web"
	"log"
)

func init() {
	fmt.Println("xi.recv")
	fmt.Println("===")
	fmt.Println("The opensource direct message service for LINE OpenChat (LINE Square).")
	fmt.Println("\n(c) 2021 Star Inc.")
}

func main() {
	// Initialization
	router := gin.Default()
	// Index
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 200,
			"data": map[string]string{
				"description": "The opensource direct message service for LINE OpenChat (LINE Square).",
				"information": "https://xi.starinc.xyz/",
				"license":     "Apache License 2.0",
				"copyright":   "(c) 2021 Star Inc. and its contributors.",
			},
		})
	})
	// Triggers
	preload := []func() Controllers.Interface{Bot.NewBot, Web.NewWeb}
	var controllers []Controllers.Interface
	for _, controller := range preload {
		instance := controller()
		controllers = append(controllers, instance)
		instance.Trigger()
	}
	// Execute
	log.Println("Start")
	err := router.Run(":10101")
	if err != nil {
		panic(err)
	}
}
