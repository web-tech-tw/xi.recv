package Bot

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal/Controllers"
	"github.com/star-inc/xi.recv/internal/Kernel/Config"
	"os"
)

type Bot struct{}

func NewBot() Controllers.Interface {
	instance := new(Bot)
	return instance
}

func (b Bot) RouterSetup(router *gin.Engine) {
	bot, err := linebot.New(
		os.Getenv(Config.LineSecret),
		os.Getenv(Config.LineToken),
	)
	if err != nil {
		panic(err)
	}
	router.GET("/bot/webhook", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			panic(err)
		}
		eventHandler(events)
	})

}

func eventHandler(events []*linebot.Event) {
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			break
		}
	}
}
