// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Active

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal/Data"
	"github.com/star-inc/xi.recv/internal/Service/Web/Hook"
)

func UserHandler(b Hook.Bot, event *linebot.Event) {
	user := Data.NewUser(event.Source.UserID)
	if event.Message.Type() != linebot.MessageTypeText {
		return
	}
	textMessage := event.Message.(*linebot.TextMessage)
	switch textMessage.Text {
	case "issue":
		room := Data.NewRoom(user.(*Data.User)).(*Data.Room)
		b.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(room.UUID))
	}
}
