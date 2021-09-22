package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MenuRepository struct {
	DB *sqlx.DB
}

func (r *MenuRepository) GetMyTasks(userID int) *[]models.Task {
	var tasks []models.Task
	user, err := GetUser(userID)
	if err != nil {
		logrus.Error(err)
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE bot_user_id=$1", configs.TaskTable)
	err = r.DB.Select(&tasks, query, user.ID)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return &tasks
}

func (r *MenuRepository) GetAllTasks() {

}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{DB: configs.DB}
}
