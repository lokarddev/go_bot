package handlers

import (
	"GoBot/pkg"
	"GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type DashboardHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *DashboardHandler) StartHandler() {
	if h.triggerHandler(h.Ctx) {
		service := services.DashboardService{Bot: h.Bot, Ctx: h.Ctx}
		service.DashboardService()
	}
}

func (h *DashboardHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	switch pkg.DashbaordPermissions[ctx.Message.Text] {
	case "":
		logrus.Error(pkg.UnavailableInputMessage)
		pkg.UnavailableInput(h.Bot, h.Ctx)
		return false
	default:
		return true
	}
}

func NewDashboardHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *DashboardHandler {
	return &DashboardHandler{Bot: bot, Ctx: ctx}
}
