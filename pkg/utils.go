package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"unicode/utf8"
)

func BackButtonAction(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) {
	message := tgbotapi.NewMessage(ctx.Message.Chat.ID, "Menu")
	message.BaseChat.ReplyMarkup = MenuKeyboard
	_, err := bot.Send(message)
	if err != nil {
		return
	}
}

func ValidPhone(phone string) string {
	_, i := utf8.DecodeRuneInString(phone)
	return phone[i:]
}
