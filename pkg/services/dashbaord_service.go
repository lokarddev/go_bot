package services

import (
	"GoBot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DashboardService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *DashboardService) DashboardService() {
	inline := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Load dashboard")
	inline.BaseChat.ReplyMarkup = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("test task 1", "test"),
		tgbotapi.NewInlineKeyboardButtonData("test task 2", "test"),
		tgbotapi.NewInlineKeyboardButtonData("test task 3", "test"))

	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Here is all active tasks")
	message.BaseChat.ReplyMarkup = pkg.DashboardKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}
