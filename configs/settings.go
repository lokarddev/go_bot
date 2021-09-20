package configs

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
)

var (
	Token       string
	Address     string
	Port        string
	TelegramUrl string
	Debug       bool
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	DbSsl       string
	DB          *sqlx.DB
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// db table names

var (
	BotUserTable   = "bot_user"
	DashboardTable = "dashboard"
	StateTable     = "state"
	TaskTable      = "task"
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("Error loading .env file")
	}
	_, err = godotenv.Read()
	if err != nil {
		logrus.Info(".env file cannot be read")
	}

	Address = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	Token = os.Getenv("TOKEN")
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	TelegramUrl = "https://api.telegram.org/bot"

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbSsl = os.Getenv("DB_SSL")
}

func SetHook() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Println(err)
		return
	}
	url := Address + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func NewPostgresDB(cfg DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() {
	database, err := NewPostgresDB(DbConfig{
		Host:     DbHost,
		Port:     DbPort,
		Username: DbUser,
		Password: DbPassword,
		DBName:   DbName,
		SSLMode:  DbSsl,
	})
	if err != nil {
		log.Fatalf("failed to initialize, %s", err)
	}
	DB = database
}
