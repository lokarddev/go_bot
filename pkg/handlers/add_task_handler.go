package handlers

import (
	"GoBot/models"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	"GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type AddTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AddTasksHandler) StartHandler() {
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
		default:
			service.AllTasksAddStep2()
			state := models.State{Current: pkg.StatePosition["AddStep2"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (h *AddTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["AddStep1"]}
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
			return true
		}
	}
	return false
}

func NewAddTasksHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *AddTasksHandler {
	return &AddTasksHandler{Bot: bot, Ctx: ctx}
}

type AddTasksFinishHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AddTasksFinishHandler) StartHandler() {
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
		default:
			service.AllTasksAddStep3()
			state := models.State{Current: pkg.StatePosition["AllTasks"]}
			err := repository.SetState(h.Ctx, state)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (h *AddTasksFinishHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	ValidState := models.State{Current: pkg.StatePosition["AddStep2"]}
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
			return true
		}
	}
	return false
}

func NewAddTasksFinishHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *AddTasksFinishHandler {
	return &AddTasksFinishHandler{Bot: bot, Ctx: ctx}
}
