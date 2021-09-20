package pkg

// permissions is a valid input types for every keyboard

var (
	MenuPermissions = map[string]string{
		MyTasksKey:   MyTasksKey,
		DashboardKey: DashboardKey,
		AllTasksKey:  AllTasksKey,
		StartKey:     StartKey,
	}
	MyTasksPermissions = map[string]string{
		BackKey:  BackKey,
		StartKey: StartKey,
	}
	AllTasksPermissions = map[string]string{
		AddTaskKey: AddTaskKey,
		BackKey:    BackKey,
		StartKey:   StartKey,
	}
	DashbaordPermissions = map[string]string{
		BackKey:  BackKey,
		StartKey: StartKey,
	}
	SingleTaskAllPermissions = map[string]string{
		EditTaskButton:   EditTaskButton,
		RemoveTaskButton: RemoveTaskButton,
		BackKey:          BackKey,
		StartKey:         StartKey,
	}
	SingleTaskMyPermissions = map[string]string{
		StartTask:   StartTask,
		DeclineTask: DeclineTask,
		DoneTask:    DoneTask,
		BackKey:     BackKey,
		StartKey:    StartKey,
	}
	AddTaskPermissions = map[string]string{
		BackKey:  BackKey,
		StartKey: StartKey,
	}
)
