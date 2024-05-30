package usecase

import (
	"io"

	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskUsecase interface {
	CreateTaskUsecase(request *dto.CreateTaskResquest, adminId string) (*task.TaskChallenge, error)
	GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error)
	UploadThumbnailUsecae(taskId string, thumbnail io.Reader) (string, error)
}
