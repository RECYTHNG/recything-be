package usecase

import (
	"mime/multipart"

	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type UserTaskUsecase interface {
	GetAllTasksUsecase() ([]task.TaskChallenge, error)
	GetTaskByIdUsecase(id string) (*task.TaskChallenge, error)
	CreateUserTaskUsecase(request *dto.UserTaskRequestCreate, userId string) (*user_task.UserTaskChallenge, error)
	UploadImageTaskUsecase(request *dto.UploadImageTask, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error)
}
