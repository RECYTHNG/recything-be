package entity

import (
	"time"

	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"gorm.io/gorm"
)

type ApprovalTask struct {
	ID                  string `gorm:"primaryKey"`
	UserTaskChallengeId string `gorm:"index"`
	UserTaskChallenge   user_task.UserTaskChallenge

	Point     int
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
