package Message

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/star-inc/xi.recv/internal/Controller/Bot"
	"github.com/star-inc/xi.recv/internal/Model"
)

func UserHandler(b Bot.Bot, event *linebot.Event) {
	user := Model.NewUser(event.Source.UserID)
	if event.Message.Type() != linebot.MessageTypeText {
		return
	}
	textMessage := event.Message.(*linebot.TextMessage)
	switch textMessage.Text {
	case "issue":
		room := Model.NewRoom(user.(*Model.User)).(*Model.Room)
		b.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(room.UUID))
	}
}
