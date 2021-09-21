package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func IsValid(userState, validState models.State) (bool, error) {

	return true, nil
}

func GetState(user models.BotUser) (models.State, error) {
	var state models.State
	query := fmt.Sprintf("SELECT * FROM %s WHERE bot_user_id = $1", configs.StateTable)
	err := configs.DB.Get(&state, query, user.ID)
	if err != nil {
		logrus.Error(err)
		return state, err
	}
	return state, nil
}

func SetState(ctx *tgbotapi.Update, state models.State) error {
	user, err := GetUser(ctx)
	var newState models.State
	tx, err := configs.DB.Begin()
	if err != nil {
		logrus.Error(err)
		return err
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE bot_user_id = $1", configs.StateTable)
	err = configs.DB.Get(&newState, query, user.ID)
	if err != nil {
		query = fmt.Sprintf("INSERT INTO %s (current, bot_user_id) VALUES ($1, $2)", configs.StateTable)
		_, err = tx.Exec(query, state.Current, user.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		query = fmt.Sprintf("UPDATE %s SET current = $1 WHERE bot_user_id = $2", configs.StateTable)
		_, err = tx.Exec(query, state.Current, user.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func GetUser(ctx *tgbotapi.Update) (models.BotUser, error) {
	var user models.BotUser
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone = $1", configs.BotUserTable)
	err := configs.DB.Get(&user, query, ctx.Message.Contact.PhoneNumber)
	if err != nil {
		logrus.Error(err)
		return user, err
	}
	return user, nil
}
