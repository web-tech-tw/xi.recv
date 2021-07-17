package Message

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal/Controller/Bot"
	"github.com/star-inc/xi.recv/internal/Model"
)

func UserHandler(b Bot.Bot, event *linebot.Event) {
	if event.Message.Type() != linebot.MessageTypeText {
		return
	}
	textMessage := event.Message.(*linebot.TextMessage)
	switch textMessage.Text {
	case "issue":
		room := Model.NewRoom(b.User).(*Model.Room)
		b.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(room.UUID))
	}
}
