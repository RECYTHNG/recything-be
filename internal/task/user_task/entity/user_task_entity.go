package entity

import (
	"time"

	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/user"
	"gorm.io/gorm"
)

type UserTaskChallenge struct {
	ID              string             `gorm:"primaryKey"`
	UserId          string             `gorm:"index"`
	User            user.User          `gorm:"foreignKey:UserId"`
	TaskChallengeId string             `gorm:"index"`
	TaskChallenge   task.TaskChallenge `gorm:"foreignKey:TaskChallengeId"`
	Status          bool
	ImageTask       []UserTaskImage `gorm:"foreignKey:TaskChallengeId"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt  `gorm:"index"`
}

type UserTaskImage struct {
	ID              string `gorm:"primaryKey"`
	TaskChallengeId string `gorm:"index"`
	ImageUrl        string
	Description     string
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
