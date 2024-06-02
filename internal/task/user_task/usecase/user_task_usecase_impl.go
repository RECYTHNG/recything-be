package usecase

import (
	"mime/multipart"

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
	findtask, errFindTask := usecase.ManageTaskRepository.FindTask(request.TaskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}

	if findtask.Status == false {
		return nil, pkg.ErrTaskCannotBeFollowed
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

func (usecase *UserTaskUsecaseImpl) UploadImageTaskUsecase(request *dto.UploadImageTask, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	findUserTask, errFind := usecase.ManageTaskRepository.FindUserTask(userId, userTaskId)

	if errFind != nil {
		return nil, pkg.ErrUserTaskNotFound
	}

	if findUserTask.StatusProgress == "done" {
		return nil, pkg.ErrUserTaskDone
	}

	validImages, errImages := helper.ImagesValidation(fileImage)
	if errImages != nil {
		return nil, errImages
	}

	var imageUrls []string
	for _, image := range validImages {
		imageUrl, err := helper.UploadToCloudinary(image, "task_images+"+userTaskId)
		if err != nil {
			return nil, pkg.ErrUploadCloudinary
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	data := &user_task.UserTaskChallenge{
		DescriptionImage: request.Description,
		StatusProgress:   "done",
		ImageTask:        []user_task.UserTaskImage{},
	}

	for _, image := range imageUrls {
		data.ImageTask = append(data.ImageTask, user_task.UserTaskImage{
			ImageUrl: image,
		})
	}

	userTask, err := usecase.ManageTaskRepository.UploadImageTask(data, userTaskId)
	if err != nil {
		return nil, err
	}
	return userTask, nil

}
