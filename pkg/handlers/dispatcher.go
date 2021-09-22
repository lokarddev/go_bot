package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DispatcherInterface interface {
	attach()
	callService()
}

type MessageDispatcher struct {
	services []BaseHandler
}

func (d *MessageDispatcher) attach(service BaseHandler) {
	d.services = append(d.services, service)
}

func (d *MessageDispatcher) callService() {
	for _, service := range d.services {
		go service.StartHandler()
	}
}

type CallBackDispatcher struct {
	services []BaseHandler
}

func (d *CallBackDispatcher) attach(service BaseHandler) {
	d.services = append(d.services, service)
}

func (d *CallBackDispatcher) callService() {
	for _, service := range d.services {
		go service.StartHandler()
	}
}

func InitHandlerDispatcher(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *MessageDispatcher {
	dispatcher := &MessageDispatcher{}
	dispatcher.attach(NewConversationHandler(bot, ctx))
	dispatcher.attach(NewContactHandler(bot, ctx))
	dispatcher.attach(NewMenuHandler(bot, ctx))
	dispatcher.attach(NewDashboardHandler(bot, ctx))
	dispatcher.attach(NewAllTasksHandler(bot, ctx))
	dispatcher.attach(NewMyTasksHandler(bot, ctx))
	return dispatcher
}

func InitCallbackDispatcher(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *CallBackDispatcher {
	dispatcher := &CallBackDispatcher{}
	dispatcher.attach(NewMyTasksCallbackHandler(bot, ctx))
	dispatcher.attach(NewAllTasksCallbackHandler(bot, ctx))
	return dispatcher
}
