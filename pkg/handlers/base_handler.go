package handlers

import (
	"GoBot/configs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
)

func WebhookHandler(c *gin.Context) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(c.Request.Body)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}
	bot, err := tgbotapi.NewBotAPI(configs.Token)
	if err != nil {
		logrus.Error(err)
	}
	switch update.Message.Text {
	case "/start":
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello motor")
		message.BaseChat.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("test")))
		_, err = bot.Send(message)
		if err != nil {
			return
		}
	}
}
