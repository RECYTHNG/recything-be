package repository

import (
	"time"

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

func (repository *ApprovalTaskRepositoryImpl) FindUserTask(userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().Where("id = ?", userTaskId).First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

func (repository *ApprovalTaskRepositoryImpl) ApproveUserTask(status string, userTaskId string) error {
	var userTask user_task.UserTaskChallenge
	acceptedAt := time.Now()
	if err := repository.DB.GetDB().Model(&userTask).Updates(map[string]interface{}{
		"status_accept": status,
		"accepted_at":   acceptedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (repository *ApprovalTaskRepositoryImpl) RejectUserTask(data *user_task.UserTaskChallenge, userTaskId string) error {
	if err := repository.DB.GetDB().Where("id = ?", userTaskId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ApprovalTaskRepositoryImpl) GetUserTaskDetails(userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("UserTaskImage").
		Preload("User").
		Preload("TaskChallenge").
		Where("id = ?", userTaskId).First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}
