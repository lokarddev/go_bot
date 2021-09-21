package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type MyTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *MyTasksHandler) StartHandler() {

}

func (h *MyTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["MyTasks"]}
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
			switch pkg.MyTasksPermissions[ctx.Message.Text] {
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

func NewMyTasksHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MyTasksHandler {
	return &MyTasksHandler{
		Bot: bot,
		Ctx: ctx}
}
