package Bot

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal/Controller"
	"github.com/star-inc/xi.recv/internal/Controller/Bot/Message"
	"github.com/star-inc/xi.recv/internal/Kernel/Config"
	"os"
)

type Bot struct {
	Client *linebot.Client
}

func NewBot() Controller.Interface {
	instance := new(Bot)
	var err error
	instance.Client, err = linebot.New(
		os.Getenv(Config.LineSecret),
		os.Getenv(Config.LineToken),
	)
	if err != nil {
		panic(err)
	}
	return instance
}

func (b Bot) RouterSetup(router *gin.Engine) {
	router.GET("/bot/webhook", func(c *gin.Context) {
		events, err := b.Client.ParseRequest(c.Request)
		if err != nil {
			panic(err)
		}
		b.eventHandler(events)
	})
}

func (b Bot) eventHandler(events []*linebot.Event) {
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			b.messageHandler(event)
			break
		}
	}
}

func (b Bot) messageHandler(event *linebot.Event) {
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		Message.UserHandler(b, event)
		break
	case linebot.EventSourceTypeGroup:
		b.Client.LeaveGroup(event.Source.GroupID)
		break
	case linebot.EventSourceTypeRoom:
		b.Client.LeaveGroup(event.Source.RoomID)
		break
	}
}
