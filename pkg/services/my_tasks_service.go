package services

import (
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MyTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.MyTasksRepository
}

type MyTasksCallbackService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.MyTasksRepository
}

func (s *MyTasksService) MyTasksService() {

}

func (s *MyTasksCallbackService) ProcessCallback() {
	//taskId := s.Ctx.CallbackQuery.Data

}
