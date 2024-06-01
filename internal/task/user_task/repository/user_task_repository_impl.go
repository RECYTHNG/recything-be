package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"gorm.io/gorm"
)

type UserTaskRepositoryImpl struct {
	DB database.Database
}

func NewUserTaskRepository(db database.Database) *UserTaskRepositoryImpl {
	return &UserTaskRepositoryImpl{DB: db}
}

func (repository *UserTaskRepositoryImpl) GetAllTasks() ([]task.TaskChallenge, error) {
	var tasks []task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Order("id desc").
		Find(&tasks, "status = ?", true).
		Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repository *UserTaskRepositoryImpl) GetTaskById(id string) (*task.TaskChallenge, error) {
	var task task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		First(&task, "id = ? and status = ?", id, true).
		Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (repository *UserTaskRepositoryImpl) FindLastIdTaskChallenge() (string, error) {
	var task user_task.UserTaskChallenge
	if err := repository.DB.GetDB().Unscoped().Order("id desc").First(&task).Error; err != nil {
		return "UT0000", err
	}
	return task.ID, nil
}

func (repository *UserTaskRepositoryImpl) CreateUserTask(userTask *user_task.UserTaskChallenge) (*user_task.UserTaskChallenge, error) {
	var result user_task.UserTaskChallenge
	err := repository.DB.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userTask).Error; err != nil {
			return err
		}
		if err := tx.Preload("TaskChallenge.TaskSteps").
			Where("user_task_challenges.id = ?", userTask.ID).
			First(&result).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository *UserTaskRepositoryImpl) FindUserTask(userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().Where("user_id = ? and task_challenge_id = ?", userId, userTaskId).First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

func (repository *UserTaskRepositoryImpl) FindTask(taskId string) (*task.TaskChallenge, error) {
	var task task.TaskChallenge
	if err := repository.DB.GetDB().First(&task, "id = ?", taskId).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
