package repository

import (
	"GoBot/configs"
	"github.com/jmoiron/sqlx"
)

type MenuRepository struct {
	DB *sqlx.DB
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{DB: configs.DB}
}
