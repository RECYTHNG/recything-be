package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
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
