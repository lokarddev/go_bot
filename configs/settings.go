package configs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
)

var (
	Token       string
	Address     string
	Port        string
	TelegramUrl string
	Debug       bool
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("Error loading .env file")
	}
	_, err = godotenv.Read()
	if err != nil {
		logrus.Info(".env file cannot be read")
	}

	Address = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	Token = os.Getenv("TOKEN")
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	TelegramUrl = "https://api.telegram.org/bot"
}

func SetHook() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Println(err)
		return
	}
	url := Address + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}
