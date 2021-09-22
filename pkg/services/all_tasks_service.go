package services

import (
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type AllTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.AllTasksRepository
}

type AllTasksCallbackService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.AllTasksRepository
}

func (s *AllTasksService) AllTasksService() {

}

func (s *AllTasksService) processCallback() {

}
