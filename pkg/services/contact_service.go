package services

import (
	"GoBot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ContactService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *ContactService) CheckContact() {
	// TODO: implement checking of user in base, Only registered by staff or superuser have an access to bot.
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Menu")
	message.BaseChat.ReplyMarkup = pkg.MenuKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}
