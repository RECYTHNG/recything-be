package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/task/manage_task/entity"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"gorm.io/gorm"
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

func (repository *ManageTaskRepositoryImpl) FindLastIdTaskChallenge() (string, error) {
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

func (repository *ManageTaskRepositoryImpl) GetTaskById(id string) (*task.TaskChallenge, error) {
	var task *task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Preload("Admin").
		First(&task, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *ManageTaskRepositoryImpl) FindTask(id string) (*task.TaskChallenge, error) {
	log.Println("Finding task with ID:", id)
	var task task.TaskChallenge
	if err := repository.DB.GetDB().Where("id = ?", id).First(&task).Error; err != nil {
		log.Println("Error finding task:", err)
		return nil, err
	}
	return &task, nil
}

func (repository *ManageTaskRepositoryImpl) UpdateTaskChallenge(taskChallenge *task.TaskChallenge, taskId string) (*task.TaskChallenge, error) {
	log.Println("Updating task with ID:", taskId)
	tx := repository.DB.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Transaction rollback due to panic:", r)
		}
	}()
	if err := tx.Where("task_challenge_id = ?", taskId).Delete(&task.TaskStep{}).Error; err != nil {
		log.Println("Error deleting task steps:", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(taskChallenge).Error; err != nil {
		log.Println("Error updating task challenge:", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, err
	}

	return taskChallenge, nil
}

func (repository *ManageTaskRepositoryImpl) DeleteTaskChallenge(taskId string) error {
	if err := repository.DB.GetDB().Where("id = ?", taskId).Delete(&task.TaskChallenge{}).Error; err != nil {
		return err
	}
	return nil
}

// cronjob for update task challenge status
func UpdateTaskChallengeStatus(db *ManageTaskRepositoryImpl) {
	now := time.Now()
	result := db.DB.GetDB().Model(&entity.TaskChallenge{}).
		Where("end_date < ? AND status = ?", now, true).Update("status", false)

	if result.Error != nil {
		log.Printf("Error updating task challenge status: %v", result.Error)
	} else {
		log.Printf("Updated %d task challenge(s) status to false", result.RowsAffected)
	}
}
