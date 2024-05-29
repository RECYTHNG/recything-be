package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepository interface {
	CreateTask(task *task.TaskChallange) (*task.TaskChallange, error)
	FindLastIdTaskChallange() (string, error)
}
