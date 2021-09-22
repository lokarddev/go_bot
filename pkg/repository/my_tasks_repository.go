package repository

import (
	"GoBot/configs"
	"GoBot/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MyTasksRepository struct {
	DB *sqlx.DB
}

func (r *MyTasksRepository) GetTask(taskId int) *models.Task {
	var task models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", configs.TaskTable)
	err := r.DB.Get(&task, query, &taskId)
	if err != nil {
		logrus.Error(err)
	}
	return &task
}

func NewMyTasksRepository() *MyTasksRepository {
	return &MyTasksRepository{DB: configs.DB}
}
