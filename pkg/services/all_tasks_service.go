package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type AllTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *AllTasksService) AllTasksService() {

}
