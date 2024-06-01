package usecase

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/repository"
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
