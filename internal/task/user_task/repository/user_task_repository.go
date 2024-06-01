package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type UserTaskRepository interface {
	GetAllTasks() ([]task.TaskChallenge, error)
}
