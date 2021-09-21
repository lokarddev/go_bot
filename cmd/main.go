package main

import (
	"GoBot/configs"
	"GoBot/pkg/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.InitEnvVariables()
	configs.InitDB()
	port := configs.Port
	if port == "" {
		logrus.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	configs.SetHook()
	router.POST("/"+configs.Token, handlers.WebhookHandler)
	err := router.Run(":" + port)
	if err != nil {
		logrus.Error(err)
	}
}
