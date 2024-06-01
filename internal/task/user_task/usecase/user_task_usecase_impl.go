package usecase

import (
	"github.com/sawalreverr/recything/internal/helper"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
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

func (usecase *UserTaskUsecaseImpl) CreateUserTaskUsecase(request *dto.UserTaskRequestCreate, userId string) (*user_task.UserTaskChallenge, error) {
	if _, err := usecase.ManageTaskRepository.FindTask(request.TaskChallengeId); err != nil {
		return nil, pkg.ErrTaskNotFound
	}

	if _, err := usecase.ManageTaskRepository.FindUserTask(userId, request.TaskChallengeId); err == nil {
		return nil, pkg.ErrUserTaskExist
	}

	lastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(lastId, "UT")
	userTask := &user_task.UserTaskChallenge{
		ID:              id,
		UserId:          userId,
		TaskChallengeId: request.TaskChallengeId,
	}

	userTaskData, err := usecase.ManageTaskRepository.CreateUserTask(userTask)
	if err != nil {
		return nil, err
	}
	return userTaskData, nil

}
