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
		service.ShareService()
	case false:
		return
	}
}

func (h *ConversationHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	if ctx.Message != nil {
		if ctx.Message.Text == "/start" {
			return true
		}
	}
	return false
}

func NewConversationHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *ConversationHandler {
	return &ConversationHandler{Bot: bot, Ctx: ctx}
}
