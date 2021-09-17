package pkg

import (
	"GoBot/configs"
	"encoding/json"
	"fmt"
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

	fmt.Println(update.Message.Text)
	if update.Message.Text == "test" {
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
		_, err := bot.Send(message)
		if err != nil {
			return
		}
	}
}
