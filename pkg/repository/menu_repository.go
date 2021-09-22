package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"GoBot/pkg"
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

func (r *MenuRepository) GetAllTasks() *[]models.Task {
	var tasks []models.Task
	query := fmt.Sprintf("SELECT id, name, description, status, created_at, updated_at FROM %s WHERE status=$1", configs.TaskTable)
	err := r.DB.Select(&tasks, query, pkg.TaskPool)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return &tasks
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{DB: configs.DB}
}
