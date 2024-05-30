package repository

import (
	"fmt"

	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepositoryImpl struct {
	DB database.Database
}

func NewManageTaskRepository(db database.Database) *ManageTaskRepositoryImpl {
	return &ManageTaskRepositoryImpl{DB: db}
}

func (r *ManageTaskRepositoryImpl) CreateTask(task *task.TaskChallenge) (*task.TaskChallenge, error) {
	if err := r.DB.GetDB().Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *ManageTaskRepositoryImpl) FindLastIdTaskChallange() (string, error) {
	var task *task.TaskChallenge
	if err := repository.DB.GetDB().Unscoped().Order("id desc").First(&task).Error; err != nil {
		return "TM0000", err
	}
	return task.ID, nil
}

func (repository *ManageTaskRepositoryImpl) GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error) {
	var tasks []task.TaskChallenge
	var total int64
	offset := (page - 1) * limit

	db := repository.DB.GetDB()
	err := db.Model(&task.TaskChallenge{}).Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count total tasks: %w", err)
	}

	err = db.Preload("TaskSteps").Preload("Admin").Limit(limit).Offset(offset).Order("id desc").Find(&tasks).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	return tasks, int(total), nil
}
