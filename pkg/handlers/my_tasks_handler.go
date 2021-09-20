package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type MyTasksHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *MyTasksHandler) StartHandler() {

}

func (h *MyTasksHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	return false
}

func NewMyTasksHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MyTasksHandler {
	return &MyTasksHandler{
		Bot: bot,
		Ctx: ctx}
}
