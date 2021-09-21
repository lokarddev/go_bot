package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DispatcherInterface interface {
	attach()
	callService()
}

type Dispatcher struct {
	services []BaseHandler
}

func (d *Dispatcher) attach(service BaseHandler) {
	d.services = append(d.services, service)
}

func (d *Dispatcher) callService() {
	for _, service := range d.services {
		go service.StartHandler()
	}
}

func InitHandlerDispatcher(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) *Dispatcher {
	dispatcher := &Dispatcher{}
	dispatcher.attach(NewConversationHandler(bot, ctx))
	dispatcher.attach(NewContactHandler(bot, ctx))
	dispatcher.attach(NewMenuHandler(bot, ctx))
	dispatcher.attach(NewDashboardHandler(bot, ctx))
	dispatcher.attach(NewAllTasksHandler(bot, ctx))
	dispatcher.attach(NewMyTasksHandler(bot, ctx))
	return dispatcher
}
