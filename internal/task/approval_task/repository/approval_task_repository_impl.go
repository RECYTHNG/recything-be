package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskRepositoryImpl struct {
	DB database.Database
}

func NewApprovalTaskRepository(db database.Database) *ApprovalTaskRepositoryImpl {
	return &ApprovalTaskRepositoryImpl{DB: db}
}

func (repository *ApprovalTaskRepositoryImpl) GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error) {
	var tasks []*user_task.UserTaskChallenge
	var total int64
	offset = (offset - 1) * limit

	if err := repository.DB.GetDB().Model(&user_task.UserTaskChallenge{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := repository.DB.GetDB().
		Preload("TaskChallenge.TaskSteps").
		Preload("User").
		Limit(limit).
		Offset(offset).Order("id desc").Find(&tasks).Error; err != nil {
		return nil, 0, err
	}
	return tasks, int(total), nil
}
