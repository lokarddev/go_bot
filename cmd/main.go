package main

import (
	"GoBot/configs"
	"GoBot/pkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.InitEnvVariables()
	port := configs.Port
	if port == "" {
		logrus.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	configs.SetHook()
	router.POST("/"+configs.Token, pkg.WebhookHandler)
	err := router.Run(":" + port)
	if err != nil {
		logrus.Error(err)
	}
}
