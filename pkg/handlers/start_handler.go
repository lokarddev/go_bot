package handlers

import (
	"GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ConversationHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *ConversationHandler) StartHandler() {
	switch h.triggerHandler(h.Ctx) {
	case true:
		service := services.StartService{Bot: h.Bot, Ctx: h.Ctx}
		service.SomeAction()
	case false:
		return
	}
}

func (h *ConversationHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	switch ctx.Message.Text {
	case "/start":
		return true
	default:
		return false
	}
}

func NewConversationHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *ConversationHandler {
	return &ConversationHandler{Bot: bot, Ctx: ctx}
}
