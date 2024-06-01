package usecase

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/repository"
	"github.com/sawalreverr/recything/pkg"
)

type UserTaskUsecaseImpl struct {
	ManageTaskRepository repository.UserTaskRepository
}

func NewUserTaskUsecase(repository repository.UserTaskRepository) *UserTaskUsecaseImpl {
	return &UserTaskUsecaseImpl{ManageTaskRepository: repository}
}

func (usecase *UserTaskUsecaseImpl) GetAllTasksUsecase() ([]task.TaskChallenge, error) {
	userTask, err := usecase.ManageTaskRepository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetTaskByIdUsecase(id string) (*task.TaskChallenge, error) {
	userTask, err := usecase.ManageTaskRepository.GetTaskById(id)
	if err != nil {
		return nil, pkg.ErrTaskNotFound
	}

	return userTask, nil
}
