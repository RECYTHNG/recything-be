package repository

import (
	"log"
	"time"

	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user_entity "github.com/sawalreverr/recything/internal/user"
)

type ApprovalTaskRepositoryImpl struct {
	DB database.Database
}

func NewApprovalTaskRepositoryImpl(db database.Database) *ApprovalTaskRepositoryImpl {
	return &ApprovalTaskRepositoryImpl{
		DB: db,
	}
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
	tx := repository.DB.GetDB().Begin()

	// Load the user task first to ensure we have all its fields
	if err := tx.Where("id = ?", userTaskId).First(&userTask).Error; err != nil {
		tx.Rollback()
		return err
	}

	acceptedAt := time.Now()
	if err := tx.Model(&userTask).Where("id = ?", userTaskId).Updates(map[string]interface{}{
		"status_accept": status,
		"accepted_at":   acceptedAt,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	point := userTask.Point
	log.Println("user id: ", userTask.UserId)

	var user user_entity.User
	if err := tx.Where("id = ?", userTask.UserId).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	pointUpdate := int(user.Point) + point
	if err := tx.Model(&user_entity.User{}).Where("id = ?", userTask.UserId).Update("point", pointUpdate).Error; err != nil {
		tx.Rollback()
		return err
	}

	var achievements []achievement.Achievement

	if err := tx.Model(&achievement.Achievement{}).Find(&achievements).Error; err != nil {
		tx.Rollback()
		return err
	}

	log.Println("achievements: ", achievements)

	if len(achievements) == 0 {
		tx.Rollback()
		return nil
	}
	var badge string
	for _, ach := range achievements {
		if point >= ach.TargetPoint {
			badge = ach.Level
			break
		}
	}

	if err := tx.Model(&user_entity.User{}).Where("id = ?", userTask.UserId).Update("badge", badge).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
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

func (repository *ApprovalTaskRepositoryImpl) GetUserTaskDetails(userTaskId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("User").
		Preload("TaskChallenge").
		Where("id = ?", userTaskId).
		First(&userTask).Error; err != nil {
		return nil, nil, err
	}

	var images []*user_task.UserTaskImage
	if err := repository.DB.GetDB().
		Where("user_task_challenge_id = ?", userTaskId).
		Find(&images).Error; err != nil {
		return nil, nil, err
	}

	return &userTask, images, nil
}
