package models

import "time"

type BotUser struct {
	ID      int       `json:"id"`
	Phone   string    `json:"phone"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

type DashBoard struct {
	*BotUser
	*Task
}

type State struct {
	ID       int       `json:"id"`
	Previous string    `json:"previous"`
	Current  string    `json:"current"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}
