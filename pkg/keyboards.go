package pkg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// keys text
const (
	StartKey = "/start"

	MyTasksKey   = "My tasks"
	DashboardKey = "Dashboard"
	AllTasksKey  = "All tasks"
	BackKey      = "Back"
	AddTaskKey   = "Add task"

	EditTaskButton   = "Edit task"
	RemoveTaskButton = "Remove task"

	StartTask   = "Start"
	DeclineTask = "Decline"
	DoneTask    = "Done"
)

// common message text

const (
	UnavailableInputMessage = "Unavailable input, please try again"
)

// menu naming for user state position management

var StatePosition = map[string]string{
	"Start":        "Start",
	"Dashboard":    "Dashboard",
	"ShareContact": "ShareContact",
	"Menu":         "Menu",
	"MyTasks":      "MyTasks",
	"AllTasks":     "AllTasks",
	"TaskAll":      "TaskAll",
	"TaskMy":       "TaskMy",
}

// some common keyboards
var (
	MenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(MyTasksKey),
			tgbotapi.NewKeyboardButton(DashboardKey),
			tgbotapi.NewKeyboardButton(AllTasksKey),
		),
	)
	AllTasksKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(AddTaskKey),
			tgbotapi.NewKeyboardButton(BackKey),
		),
	)
	MyTasksKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BackKey),
		),
	)
	DashboardKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BackKey),
		),
	)
	TaskMyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(StartTask),
			tgbotapi.NewKeyboardButton(DeclineTask),
			tgbotapi.NewKeyboardButton(DoneTask),
			tgbotapi.NewKeyboardButton(BackKey),
		),
	)
	TaskAllKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(EditTaskButton),
			tgbotapi.NewKeyboardButton(RemoveTaskButton),
			tgbotapi.NewKeyboardButton(BackKey),
		),
	)
)