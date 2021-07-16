package Bot

import (
	"github.com/gin-gonic/gin"
	"github.com/star-inc/xi.recv/internal/Controllers"
	"log"
)

type Bot struct{}

func NewBot() Controllers.Interface {
	instance := new(Bot)
	return instance
}

func (b Bot) Trigger() {
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
	// Execute
	log.Println("Start")
	err := router.Run(":10101")
	if err != nil {
		panic(err)
	}
}
