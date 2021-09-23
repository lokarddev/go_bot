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

type AllTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AllTasksHandler) StartHandler() {
	service := services.AllTasksService{
		Bot: h.Bot,
		Ctx: h.Ctx,
		DB:  repository.NewAllTasksRepository(),
	}
	if h.triggerHandler(h.Ctx) {
		switch h.Ctx.Message.Text {
		case pkg.BackKey:
			pkg.BackButtonAction(h.Bot, h.Ctx)
			state := models.State{Current: pkg.StatePosition["Menu"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		case pkg.AddTaskKey:
			service.AllTasksAddStart()
			state := models.State{Current: pkg.StatePosition["AddStart"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (h *AllTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["AllTasks"]}
	if repository.UserExists(ctx) == true {
		user, err := repository.GetUser(ctx.Message.From.ID)
		if err != nil {
			logrus.Error(err)
			return false
		}
		state, err := repository.GetState(user)
		if err != nil {
			logrus.Error(err)
			return false
		}
		if repository.IsValid(state, ValidState) {
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

type AllTasksCallbackHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AllTasksCallbackHandler) StartHandler() {
	if h.triggerHandler(h.Ctx) {
		service := services.AllTasksCallbackService{
			Bot: h.Bot,
			Ctx: h.Ctx,
			DB:  repository.NewAllTasksRepository()}
		service.ProcessCallback()
	}
}

func (h *AllTasksCallbackHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	data := strings.Split(ctx.CallbackQuery.Data, " ")[0]
	if data == pkg.AllTasksCallback {
		return true
	}
	return false
}

func NewAllTasksCallbackHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *AllTasksCallbackHandler {
	return &AllTasksCallbackHandler{
		Bot: bot,
		Ctx: ctx}
}
