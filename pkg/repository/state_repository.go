package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
)

type CheckStateRepo struct {
	DB *sqlx.DB
}

func (r *CheckStateRepo) IsValid(ctx tgbotapi.Update) error {

	return nil
}

func (r *CheckStateRepo) SetState(ctx tgbotapi.Update, user models.BotUser) error {

	return nil
}

func (r *CheckStateRepo) GetUser(ctx tgbotapi.Update) (models.BotUser, error) {
	var user models.BotUser
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone = $1", configs.BotUserTable)
	err := r.DB.Get(&user, query, ctx.Message.Contact.PhoneNumber)
	if err != nil {
		return user, err
	}
	return user, nil
}
