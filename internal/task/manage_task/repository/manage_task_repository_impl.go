package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepositoryImpl struct {
	DB database.Database
}

func NewManageTaskRepository(db database.Database) *ManageTaskRepositoryImpl {
	return &ManageTaskRepositoryImpl{DB: db}
}

func (r *ManageTaskRepositoryImpl) CreateTask(task *task.TaskChallange) (*task.TaskChallange, error) {
	if err := r.DB.GetDB().Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *ManageTaskRepositoryImpl) FindLastIdTaskChallange() (string, error) {
	var task *task.TaskChallange
	if err := repository.DB.GetDB().Order("id desc").First(&task).Error; err != nil {
		return "TM0000", err
	}
	return task.ID, nil
}
