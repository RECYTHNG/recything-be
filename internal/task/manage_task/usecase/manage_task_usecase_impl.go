package usecase

import (
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/manage_task/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ManageTaskUsecaseImpl struct {
	ManageTaskRepository repository.ManageTaskRepository
}

func NewManageTaskUsecase(repository repository.ManageTaskRepository) *ManageTaskUsecaseImpl {
	return &ManageTaskUsecaseImpl{ManageTaskRepository: repository}
}

func (usecase *ManageTaskUsecaseImpl) CreateTaskUsecase(request *dto.CreateTaskResquest, adminId string) (*task.TaskChallenge, error) {
	if len(request.Steps) == 0 {
		return nil, pkg.ErrTaskStepsNull

	}
	findLastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(findLastId, "TM")

	taskChallange := &task.TaskChallenge{
		ID:          id,
		AdminId:     adminId,
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   request.Thumbnail,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		TaskSteps:   []task.TaskStep{},
		DeletedAt:   gorm.DeletedAt{},
	}

	for _, step := range request.Steps {
		taskStep := task.TaskStep{
			TaskChallangeId: id,
			Title:           step.Title,
			Description:     step.Description,
		}
		taskChallange.TaskSteps = append(taskChallange.TaskSteps, taskStep)
	}

	if _, err := usecase.ManageTaskRepository.CreateTask(taskChallange); err != nil {
		return nil, err
	}
	return taskChallange, nil
}

func (usecase *ManageTaskUsecaseImpl) GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error) {
	tasks, total, err := usecase.ManageTaskRepository.GetTaskChallengePagination(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}
