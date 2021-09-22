package handlers

import (
	"GoBot/configs"
	"GoBot/models"
	"GoBot/pkg/repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type BaseHandler interface {
	StartHandler()
	triggerHandler(ctx *tgbotapi.Update) bool
}

func WebhookHandler(c *gin.Context) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(c.Request.Body)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		logrus.Error(err)
		return
	}
	bot, err := tgbotapi.NewBotAPI(configs.Token)
	if err != nil {
		logrus.Error(err)
	}
	if update.Message != nil {
		dispatcher := InitHandlerDispatcher(bot, &update)
		dispatcher.callService()
	} else if update.CallbackQuery != nil {
		dispatcher := InitCallbackDispatcher(bot, &update)
		dispatcher.callService()
	}
}

func PreProcess(ctx *tgbotapi.Update) (models.State, bool) {
	var state models.State
	if repository.UserExists(ctx) == true {
		user, err := repository.GetUser(ctx)
		if err != nil {
			logrus.Error(err)
			return state, false
		}
		state, err = repository.GetState(user)
		if err != nil {
			logrus.Error(err)
			return state, false
		}
	}
	return state, true
}

//	fmt.Println(update.CallbackQuery)
//	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
//	if _, err := bot.AnswerCallbackQuery(callback); err != nil {
//		logrus.Error(err)
//	}
//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
//	if _, err := bot.Send(msg); err != nil {
//		logrus.Error()
//	}
//}
