package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type AllTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AllTasksHandler) StartHandler() {

}

func (h *AllTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["AllTasks"]}
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
			switch pkg.AllTasksPermissions[ctx.Message.Text] {
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

func NewAllTasksHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *AllTasksHandler {
	return &AllTasksHandler{Bot: bot, Ctx: ctx}
}
