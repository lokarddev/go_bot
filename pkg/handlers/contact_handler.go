package handlers

import (
	"GoBot/configs"
	"GoBot/pkg"
	"GoBot/pkg/repository"
	"GoBot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ContactHandler struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (h *ContactHandler) StartHandler() {
	switch h.triggerHandler(h.Ctx) {
	case true:
		service := services.ContactService{
			Bot:  h.Bot,
			Ctx:  h.Ctx,
			Repo: &repository.StartRepository{DB: configs.DB}}
		service.CheckContact()
	}
}

func (h *ContactHandler) triggerHandler(ctx *tgbotapi.Update) bool {
	switch ctx.Message.Contact {
	case nil:
		pkg.UnavailableInput(h.Bot, h.Ctx)
		return false
	default:
		return true
	}
}

func NewContactHandler(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *ContactHandler {
	return &ContactHandler{Bot: bot, Ctx: ctx}
}
