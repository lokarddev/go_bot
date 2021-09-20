package handlers

import (
	"GoBot/pkg"
	services "GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type MenuHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *MenuHandler) StartHandler() {
	if h.triggerHandler(h.Ctx) {
		service := services.MenuService{Bot: h.Bot, Ctx: h.Ctx}
		switch h.Ctx.Message.Text {
		case pkg.MyTasksKey:
			service.MyTasksService()
		case pkg.AllTasksKey:
			service.AllTasksService()
		case pkg.DashboardKey:
			service.DashboardService()
		}
	} else {
		pkg.UnavailableInput(h.Bot, h.Ctx)
	}
}

func (h *MenuHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	switch pkg.MenuPermissions[ctx.Message.Text] {
	case "":
		logrus.Error(pkg.UnavailableInputMessage)
		return false
	default:
		return true
	}
}

func NewMenuHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MenuHandler {
	return &MenuHandler{Bot: bot, Ctx: ctx}
}
