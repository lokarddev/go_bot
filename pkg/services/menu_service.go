package services

import (
	"GoBot/pkg"
	"GoBot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"strconv"
)

type MenuService struct {
	Bot *tgbotapi.BotAPI
	Ctx *tgbotapi.Update
	DB  *repository.MenuRepository
}

func (s *MenuService) AllTasksService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "All tasks")
	inline := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "TaskList")

	baseRecords := []string{"record1", "record2"}

	var buttons [][]tgbotapi.InlineKeyboardButton

	for _, value := range baseRecords {
		butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Test all task", value))
		buttons = append(buttons, butt)
	}

	inline.BaseChat.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	_, err := s.Bot.Send(inline)
	if err != nil {
		logrus.Error(err)
	}

	message.BaseChat.ReplyMarkup = pkg.AllTasksKeyboard
	_, err = s.Bot.Send(message)
	if err != nil {
		logrus.Error(err)
	}
}

func (s *MenuService) MyTasksService() {
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Your tasks")
	inline := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "TaskList")

	taskList := s.DB.GetMyTasks(s.Ctx.Message.From.ID)
	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, task := range *taskList {
		taskID := strconv.Itoa(task.ID)
		butt := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(task.Name, taskID))
		buttons = append(buttons, butt)
	}
	inline.BaseChat.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	_, err := s.Bot.Send(inline)
	if err != nil {
		logrus.Error(err)
	}
	message.BaseChat.ReplyMarkup = pkg.MyTasksKeyboard
	_, err = s.Bot.Send(message)
	if err != nil {
		logrus.Error(err)
	}
}

func (s *MenuService) DashboardService() {
	inline := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Load dashboard")
	// TODO: displaying of active tasks per user
	inline.BaseChat.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Test task", "wq")))
	_, err := s.Bot.Send(inline)
	if err != nil {
		logrus.Error(err)
	}
	message := tgbotapi.NewMessage(s.Ctx.Message.Chat.ID, "Here is all active tasks")
	message.BaseChat.ReplyMarkup = pkg.DashboardKeyboard
	_, err = s.Bot.Send(message)
	if err != nil {
		logrus.Error(err)
	}
}
