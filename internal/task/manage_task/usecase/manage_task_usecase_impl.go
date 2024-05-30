package usecase

import (
	"io"

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
	findLastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallange()
	id := helper.GenerateCustomID(findLastId, "TM")

	taskChallange := &task.TaskChallenge{
		ID:          id,
		AdminId:     adminId,
		Title:       request.Title,
		Description: request.Description,
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

func (usecase *ManageTaskUsecaseImpl) UploadThumbnailUsecae(taskId string, thumbnail io.Reader) (string, error) {
	if _, err := usecase.ManageTaskRepository.FindTaskById(taskId); err != nil {
		return "", pkg.ErrTaskNotFound
	}
	thumbnailUrl, errUpload := helper.UploadToCloudinary(thumbnail, "task_thumbnail")
	if errUpload != nil {
		return "", pkg.ErrUploadCloudinary
	}

	if err := usecase.ManageTaskRepository.UploadThumbnail(taskId, thumbnailUrl); err != nil {
		return "", err
	}
	return thumbnailUrl, nil
}
