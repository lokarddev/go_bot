package services

import (
	"GoBot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MenuService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *MenuService) AllTasksService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "All tasks")
	message.BaseChat.ReplyMarkup = pkg.AllTasksKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}

func (s *MenuService) MyTasksService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Your tasks")
	message.BaseChat.ReplyMarkup = pkg.MyTasksKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}

func (s *MenuService) DashboardService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Dashboard")
	message.BaseChat.ReplyMarkup = pkg.DashboardKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}
