package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepository interface {
	CreateTask(task *task.TaskChallenge) (*task.TaskChallenge, error)
	FindLastIdTaskChallange() (string, error)
	GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error)
	UploadThumbnail(taskId string, thumbnail string) error
	FindTaskById(taskId string) (*task.TaskChallenge, error)
}
