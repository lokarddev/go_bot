package repository

import (
	"GoBot/configs"
	"github.com/jmoiron/sqlx"
)

type MyTasksRepository struct {
	DB *sqlx.DB
}

func NewMyTasksRepository() *MyTasksRepository {
	return &MyTasksRepository{DB: configs.DB}
}
