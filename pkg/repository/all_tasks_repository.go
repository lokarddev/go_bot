package repository

import (
	"GoBot/configs"
	"GoBot/pkg"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AllTasksRepository struct {
	DB *sqlx.DB
}

func NewAllTasksRepository() *AllTasksRepository {
	return &AllTasksRepository{DB: configs.DB}
}

func (r *AllTasksRepository) CreateTask(taskObject *pkg.TaskCreation, userID *int) {
	tx, err := r.DB.Begin()
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("HERE")
	query := fmt.Sprintf("INSERT INTO %s (name, description, status, bot_user_id) VALUES ($1, $2, $3, $4)", configs.TaskTable)
	_, err = tx.Exec(query, taskObject.Name, taskObject.Description, pkg.TaskPool, userID)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Error(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
}
