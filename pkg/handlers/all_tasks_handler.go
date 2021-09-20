package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type AllTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *AllTasksHandler) StartHandler() {

}

func (h *AllTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	return false
}

func NewAllTasksHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *AllTasksHandler {
	return &AllTasksHandler{Bot: bot, Ctx: ctx}
}
