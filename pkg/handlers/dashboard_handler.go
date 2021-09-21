package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
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
	ValidState := models.State{Current: pkg.StatePosition["Dashboard"]}
	if ctx.Message.Contact != nil {
		user, err := repository.GetUser(ctx)
		if err != nil {
			logrus.Error(err)
			return false
		}
		state, err := repository.GetState(user)
		if err != nil {
			logrus.Error(err)
			return false
		}
		valid, err := repository.IsValid(state, ValidState)
		if valid {
			switch pkg.DashbaordPermissions[ctx.Message.Text] {
			case "":
				return false
			default:
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func NewDashboardHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *DashboardHandler {
	return &DashboardHandler{Bot: bot, Ctx: ctx}
}
