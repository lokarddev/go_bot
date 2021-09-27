package pkg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// callback signatures
const (
	MyTasksCallback  = "MyTasks"
	AllTasksCallback = "AllTasks"
)

// task status list
const (
	Pending    = "pending"
	InProgress = "in_progress"
	TaskPool   = "task_pool"
	Done       = "done"
	Paused     = "paused"
)

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

	"AddStep1": "AddStep1",
	"AddStep2": "AddStep2",
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

type TaskCreation struct {
	Name        string
	Description string
}

// TaskMap A state like object for creating tasks step by step
var TaskMap = make(map[int]TaskCreation)
