package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MyTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *MyTasksService) MyTasksService() {

}
