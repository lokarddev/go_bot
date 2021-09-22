package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type DashboardHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *DashboardHandler) StartHandler() {
	if h.triggerHandler(h.Ctx) {
		switch h.Ctx.Message.Text {
		case pkg.BackKey:
			pkg.BackButtonAction(h.Bot, h.Ctx)
			state := models.State{Current: pkg.StatePosition["Menu"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (h *DashboardHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["Dashboard"]}
	state, success := PreProcess(ctx)
	if success {
		if repository.IsValid(state, ValidState) {
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
