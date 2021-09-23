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

func IsValid(userState, validState models.State) bool {
	if userState.Current == validState.Current {
		return true
	}
	return false
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
	user, err := GetUser(ctx.Message.From.ID)
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
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetUser(ID int) (models.BotUser, error) {
	var user models.BotUser
	query := fmt.Sprintf("SELECT * FROM %s WHERE t_id = $1", configs.BotUserTable)
	err := configs.DB.Get(&user, query, ID)
	if err != nil {
		logrus.Error(err)
		return user, err
	}
	return user, nil
}

func UserExists(ctx *tgbotapi.Update) bool {
	var user models.BotUser
	Id := ctx.Message.From.ID
	query := fmt.Sprintf("SELECT * FROM %s WHERE t_id = $1", configs.BotUserTable)
	err := configs.DB.Get(&user, query, Id)
	if err != nil {
		return false
	}
	return true
}
func GetTask(taskId int) *models.Task {
	var task models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", configs.TaskTable)
	err := configs.DB.Get(&task, query, &taskId)
	if err != nil {
		logrus.Error(err)
	}
	return &task
}
