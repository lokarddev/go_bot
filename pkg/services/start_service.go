package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StartService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *StartService) ShareService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Welcome to task bot! To start interact you should share your phone number.")
	message.BaseChat.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Share number"),
		),
	)
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}
