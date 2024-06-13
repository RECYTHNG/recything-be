package usecase

import (
	"mime/multipart"
	"time"

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

func (usecase *UserTaskUsecaseImpl) CreateUserTaskUsecase(taskChallengeId string, userId string) (*user_task.UserTaskChallenge, error) {
	findtask, errFindTask := usecase.ManageTaskRepository.FindTask(taskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}

	if findtask.Status == false {
		return nil, pkg.ErrTaskCannotBeFollowed
	}

	if _, err := usecase.ManageTaskRepository.FindUserTask(userId, taskChallengeId); err == nil {
		return nil, pkg.ErrUserTaskExist
	}

	if _, err := usecase.ManageTaskRepository.FindUserHasSameTask(userId, taskChallengeId); err == nil {
		return nil, pkg.ErrUserTaskExist
	}

	lastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(lastId, "UT")
	userTask := &user_task.UserTaskChallenge{
		ID:              id,
		UserId:          userId,
		TaskChallengeId: taskChallengeId,
		AcceptedAt:      time.Now(),
		ImageTask:       []user_task.UserTaskImage{},
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

	findTask, errFindTask := usecase.ManageTaskRepository.FindTask(findUserTask.TaskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}
	countImage := len(fileImage)
	countTaskSteps := len(findTask.TaskSteps) * 3
	if countImage > countTaskSteps {
		return nil, pkg.ErrImagesExceed
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
		Point:            findTask.Point,
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

func (usecase *UserTaskUsecaseImpl) GetUserTaskByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error) {
	userTask, err := usecase.ManageTaskRepository.GetUserTaskByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(userTask) == 0 {
		return nil, pkg.ErrUserNoHasTask
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskDoneByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error) {
	userTask, err := usecase.ManageTaskRepository.GetUserTaskDoneByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(userTask) == 0 {
		return nil, pkg.ErrUserNoHasTask
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) UpdateUserTaskUsecase(request *dto.UpdateUserTaskRequest, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	findUserTask, errFind := usecase.ManageTaskRepository.FindUserTask(userId, userTaskId)

	if errFind != nil {
		return nil, pkg.ErrUserTaskNotFound
	}

	if findUserTask.StatusAccept != "reject" {
		return nil, pkg.ErrUserTaskNotReject
	}

	findTask, errFindTask := usecase.ManageTaskRepository.FindTask(findUserTask.TaskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}
	countImage := len(fileImage)
	lenTaskSteps := len(findTask.TaskSteps)
	countTaskSteps := lenTaskSteps * 3

	if countImage > countTaskSteps {
		return nil, pkg.ErrImagesExceed
	}

	validImages, errImages := helper.ImagesValidation(fileImage)
	if errImages != nil {
		return nil, errImages
	}

	var imageUrls []string
	for _, image := range validImages {
		imageUrl, err := helper.UploadToCloudinary(image, "task_images_update+"+userTaskId)
		if err != nil {
			return nil, pkg.ErrUploadCloudinary
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	data := &user_task.UserTaskChallenge{
		DescriptionImage: request.Description,
		StatusAccept:     "need_rivew",
		ImageTask:        []user_task.UserTaskImage{},
	}

	for _, image := range imageUrls {
		data.ImageTask = append(data.ImageTask, user_task.UserTaskImage{
			ImageUrl: image,
		})
	}

	userTask, err := usecase.ManageTaskRepository.UpdateUserTask(data, userTaskId)
	if err != nil {
		return nil, err
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskDetailsUsecase(userTaskId string, userId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	userTask, imageTask, err := usecase.ManageTaskRepository.GetUserTaskDetails(userTaskId, userId)
	if err != nil {
		return nil, nil, pkg.ErrUserTaskNotFound
	}
	return userTask, imageTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetHistoryPointByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, int, error) {

	userTask, err := usecase.ManageTaskRepository.GetHistoryPointByUserId(userId)
	if err != nil {
		return nil, 0, err
	}
	if len(userTask) == 0 {
		return nil, 0, pkg.ErrUserNoHasTask
	}

	var totalPoint int
	for _, task := range userTask {
		totalPoint += task.Point
	}
	return userTask, totalPoint, nil
}

func (usecase *UserTaskUsecaseImpl) UpdateTaskStepUsecase(request *dto.UpdateTaskStepRequest, userId string) (*user_task.UserTaskChallenge, error) {
	userTask, errUserTask := usecase.ManageTaskRepository.FindUserTask(userId, request.UserTask)

	if errUserTask != nil {
		return nil, pkg.ErrUserTaskNotFound
	}
	if userTask.StatusAccept != "in-progress" {
		return nil, pkg.ErrUserTaskDone
	}
	errStatus := usecase.ManageTaskRepository.FindTaskStep(request.StepId, userTask.TaskChallengeId)

	if errStatus != nil {
		return nil, pkg.ErrTaskStepNotFound
	}

	errTaskStep := usecase.ManageTaskRepository.UpdateTaskStep(request.StepId, userTask.TaskChallengeId)
	if errTaskStep != nil {
		return nil, errTaskStep
	}

	return userTask, nil

}
