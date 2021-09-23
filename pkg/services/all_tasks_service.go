package services

import (
	"GoBot/pkg"
	"GoBot/pkg/repository"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

type AllTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.AllTasksRepository
}

type AllTasksCallbackService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.AllTasksRepository
}

func (s *AllTasksService) AllTasksAddStart() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Type task name")
	message.BaseChat.ReplyMarkup = pkg.AllTasksKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}

func (s *AllTasksService) AllTasksAddEnd() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Type task description")
	message.BaseChat.ReplyMarkup = pkg.AllTasksKeyboard
	_, err := s.Bot.Send(message)
	if err != nil {
		return
	}
}

func (s *AllTasksCallbackService) ProcessCallback() {
	rawId := strings.Split(s.Ctx.CallbackQuery.Data, " ")[1]
	taskId, _ := strconv.Atoi(rawId)
	task := repository.GetTask(taskId)
	msgText := fmt.Sprintf("Task: %s\n Description: %s", task.Name, task.Description)
	msg := tgbotapi.NewMessage(s.Ctx.CallbackQuery.Message.Chat.ID, msgText)
	msg.BaseChat.ReplyMarkup = pkg.TaskAllKeyboard
	if _, err := s.Bot.Send(msg); err != nil {
		panic(err)
	}
}
