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
	DB                   *gorm.DB
}

func NewManageTaskUsecase(repository repository.ManageTaskRepository) *ManageTaskUsecaseImpl {
	return &ManageTaskUsecaseImpl{ManageTaskRepository: repository}
}

func (usecase *ManageTaskUsecaseImpl) CreateTaskUsecase(request *dto.CreateTaskResquest, adminId string) (*task.TaskChallenge, error) {
	if len(request.Steps) == 0 {
		return nil, pkg.ErrTaskStepsNull

	}
	// now := time.Now()
	// oneMinute := now.Add(time.Minute * 1)
	findLastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(findLastId, "TM")

	taskChallange := &task.TaskChallenge{
		ID:          id,
		AdminId:     adminId,
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   request.ThumbnailUrl,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		TaskSteps:   []task.TaskStep{},
		DeletedAt:   gorm.DeletedAt{},
	}

	for _, step := range request.Steps {
		taskStep := task.TaskStep{
			TaskChallengeId: id,
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

func (usecase *ManageTaskUsecaseImpl) GetTaskByIdUsecase(id string) (*task.TaskChallenge, error) {
	task, err := usecase.ManageTaskRepository.GetTaskById(id)
	if err != nil {
		return nil, pkg.ErrTaskNotFound
	}
	return task, nil
}

func (usecase *ManageTaskUsecaseImpl) UpdateTaskChallengeUsecase(request *dto.UpdateTaskRequest, id string) (*task.TaskChallenge, error) {
	findTask, _ := usecase.ManageTaskRepository.FindTask(id)

	if findTask == nil {
		return nil, pkg.ErrTaskNotFound
	}
	if len(request.Steps) == 0 {
		return nil, pkg.ErrTaskStepsNull
	}

	taskChallenge := &task.TaskChallenge{
		ID:          id,
		AdminId:     findTask.AdminId,
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   request.ThumbnailUrl,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
	}

	// Add new steps
	for _, step := range request.Steps {
		taskStep := task.TaskStep{
			TaskChallengeId: id,
			Title:           step.Title,
			Description:     step.Description,
		}
		taskChallenge.TaskSteps = append(taskChallenge.TaskSteps, taskStep)
	}

	updatedTaskChallenge, err := usecase.ManageTaskRepository.UpdateTaskChallenge(taskChallenge, id)
	if err != nil {
		return nil, err
	}

	return updatedTaskChallenge, nil
}

func (usecase *ManageTaskUsecaseImpl) DeleteTaskChallengeUsecase(id string) error {
	if _, err := usecase.ManageTaskRepository.FindTask(id); err != nil {
		return pkg.ErrTaskNotFound
	}
	if err := usecase.ManageTaskRepository.DeleteTaskChallenge(id); err != nil {
		return err
	}
	return nil
}
