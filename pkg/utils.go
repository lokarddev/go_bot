package pkg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func UnavailableInput(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) {
	message := tgbotapi.NewMessage(ctx.Message.Chat.ID, UnavailableInputMessage)
	_, err := bot.Send(message)
	if err != nil {
		return
	}
}

func CheckState() {
	// TODO: check for user state status. Only after this we can get available valid inputs for him.
}

func BackButtonAction(bot *tgbotapi.BotAPI, ctx *tgbotapi.Update) {
	message := tgbotapi.NewMessage(ctx.Message.Chat.ID, "Menu")
	message.BaseChat.ReplyMarkup = MenuKeyboard
	_, err := bot.Send(message)
	if err != nil {
		return
	}
}
