package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
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
		service := services.MenuService{
			Bot: h.Bot,
			Ctx: h.Ctx,
			DB:  repository.NewMenuRepository()}

		switch h.Ctx.Message.Text {
		case pkg.MyTasksKey:
			service.MyTasksService()
			state := models.State{Current: pkg.StatePosition["MyTasks"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		case pkg.AllTasksKey:
			service.AllTasksService()
			state := models.State{Current: pkg.StatePosition["AllTasks"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		case pkg.DashboardKey:
			service.DashboardService()
			state := models.State{Current: pkg.StatePosition["Dashboard"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (h *MenuHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["Menu"]}
	state, success := PreProcess(ctx)
	if success {
		if repository.IsValid(state, ValidState) {
			switch pkg.MenuPermissions[ctx.Message.Text] {
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

func NewMenuHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MenuHandler {
	return &MenuHandler{Bot: bot, Ctx: ctx}
}
