package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type StartRepository struct {
	DB *sqlx.DB
}

func (r *StartRepository) SetUser(ctx *tgbotapi.Update) error {
	var user models.BotUser
	tx, err := r.DB.Begin()
	if err != nil {
		logrus.Error(err)
		return err
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=$1", configs.BotUserTable)
	if err = r.DB.Get(&user, query, &ctx.Message.Contact.PhoneNumber); err != nil {
		query = fmt.Sprintf("INSERT INTO %s (phone) VALUES ($1)", configs.BotUserTable)
		_, err = tx.Exec(query, &ctx.Message.Contact.PhoneNumber)
		if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				return errTx
			}
			return err
		} else {
			query = fmt.Sprintf("UPDATE %s SET first_name=$1, last_name=$2, t_id=$3, phone=$4 WHERE phone=$5", configs.BotUserTable)
			_, err = tx.Exec(query, ctx.Message.Contact.FirstName, ctx.Message.Contact.LastName, ctx.Message.Contact.UserID, &ctx.Message.Contact.PhoneNumber, &ctx.Message.Contact.PhoneNumber)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
