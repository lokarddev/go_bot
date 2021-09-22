package models

import "time"

type BotUser struct {
	ID         int       `json:"id" db:"id"`
	Firstname  string    `json:"first_name" db:"first_name"`
	Lastname   string    `json:"last_name" db:"last_name"`
	TelegramID string    `json:"t_id" db:"t_id"`
	Phone      string    `json:"phone" db:"phone"`
	Created    time.Time `json:"created_at" db:"created_at"`
	Updated    time.Time `json:"updated_at" db:"updated_at"`
}

type Task struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Created     time.Time `json:"created_at" db:"created_at"`
	Updated     time.Time `json:"updated_at" db:"updated_at"`
	BotUserId   int       `json:"bot_user_id" db:"bot_user_id"`
}

type DashBoard struct {
	*BotUser
	*Task
}

type State struct {
	ID        int       `json:"id" db:"id"`
	Previous  string    `json:"previous" db:"previous"`
	Current   string    `json:"current" db:"current"`
	Created   time.Time `json:"created_at" db:"created_at"`
	Updated   time.Time `json:"updated_at" db:"updated_at"`
	BotUserId int       `json:"bot_user_id" db:"bot_user_id"`
}
