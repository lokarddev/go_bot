package handlers

import (
	"GoBot/configs"
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
	dispatcher := InitDispatcher(bot, &update)
	dispatcher.callService()
}

//message := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to task bot, stranger!")
//message.BaseChat.ReplyMarkup = tgbotapi.NewReplyKeyboard(
//tgbotapi.NewKeyboardButtonRow(
//tgbotapi.NewKeyboardButton("test")))
//_, err = bot.Send(message)
//if err != nil {
//return
