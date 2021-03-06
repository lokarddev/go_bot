package services

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type ContactService struct {
	Bot  *tgbotapi.BotAPI
	Ctx  *tgbotapi.Update
	Repo *repository.StartRepository
}

func (s *ContactService) CheckContact() {
	// TODO: implement checking of user in base, Only registered by staff or superuser have an access to bot.
	err := s.Repo.SetUser(s.Ctx)
	if err != nil {
		logrus.Error(err)
		return
	}
	state := models.State{Current: pkg.StatePosition["Menu"]}
	err = repository.SetState(s.Ctx, state)
	if err != nil {
		logrus.Error(err)
		return
	}
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Menu")
	message.BaseChat.ReplyMarkup = pkg.MenuKeyboard
	_, err = s.Bot.Send(message)
	if err != nil {
		return
	}
}
