package services

import (
	"GoBot/pkg"
	"GoBot/pkg/repository"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

type MyTasksService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.MyTasksRepository
}

type MyTasksCallbackService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.MyTasksRepository
}

func (s *MyTasksService) MyTasksService() {

}

func (s *MyTasksCallbackService) ProcessCallback() {
	fmt.Println(s.Ctx.CallbackQuery.Data)
	taskId, _ := strconv.Atoi(s.Ctx.CallbackQuery.Data)
	task := s.DB.GetTask(taskId)
	msgText := fmt.Sprintf("Task: %s\n Description: %s", task.Name, task.Description)
	msg := tgbotapi.NewMessage(s.Ctx.CallbackQuery.Message.Chat.ID, msgText)
	msg.BaseChat.ReplyMarkup = pkg.TaskMyKeyboard
	if _, err := s.Bot.Send(msg); err != nil {
		panic(err)
	}
}
