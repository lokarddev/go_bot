package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DashboardService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}
