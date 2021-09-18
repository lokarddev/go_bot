package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StartService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
}

func (s *StartService) SomeAction() {
	fmt.Println("Hello world")
}
