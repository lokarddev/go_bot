package repository

import (
	"GoBot/configs"
	"github.com/jmoiron/sqlx"
)

type AllTasksRepository struct {
	DB *sqlx.DB
}

func NewAllTasksRepository() *AllTasksRepository {
	return &AllTasksRepository{DB: configs.DB}
}
