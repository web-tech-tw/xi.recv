// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Hook

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal"
	"github.com/star-inc/xi.recv/internal/Config"
	"github.com/star-inc/xi.recv/internal/Service/Web/Active"
)

type Bot struct {
	Client *linebot.Client
}

func NewBot() internal.Router {
	instance := new(Bot)
	var err error
	instance.Client, err = linebot.New(
		Config.LineSecret,
		Config.LineToken,
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
		Active.UserHandler(b, event)
		break
	case linebot.EventSourceTypeGroup:
		b.Client.LeaveGroup(event.Source.GroupID)
		break
	case linebot.EventSourceTypeRoom:
		b.Client.LeaveGroup(event.Source.RoomID)
		break
	}
}
