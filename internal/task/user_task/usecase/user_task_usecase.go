package usecase

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type UserTaskUsecase interface {
	GetAllTasksUsecase() ([]task.TaskChallenge, error)
	GetTaskByIdUsecase(id string) (*task.TaskChallenge, error)
}
