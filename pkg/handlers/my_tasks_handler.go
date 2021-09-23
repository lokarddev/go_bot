package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	"GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"strings"
)

type MyTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *MyTasksHandler) StartHandler() {
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

func (h *MyTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["MyTasks"]}
	state, success := PreProcess(ctx)
	if success {
		if repository.IsValid(state, ValidState) {
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

type MyTasksCallbackHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *MyTasksCallbackHandler) StartHandler() {
	if h.triggerHandler(h.Ctx) {
		service := services.MyTasksCallbackService{
			Bot: h.Bot,
			Ctx: h.Ctx,
			DB:  repository.NewMyTasksRepository()}
		service.ProcessCallback()
	}
}

func (h *MyTasksCallbackHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	data := strings.Split(ctx.CallbackQuery.Data, " ")[0]
	if data == pkg.MyTasksCallback {
		return true
	}
	return false
}

func NewMyTasksCallbackHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MyTasksCallbackHandler {
	return &MyTasksCallbackHandler{
		Bot: bot,
		Ctx: ctx}
}
