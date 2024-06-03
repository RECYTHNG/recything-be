package entity

import (
	"time"

	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/user"
	"gorm.io/gorm"
)

type UserTaskChallenge struct {
	ID               string             `gorm:"primaryKey"`
	UserId           string             `gorm:"index"`
	User             user.User          `gorm:"foreignKey:UserId"`
	TaskChallengeId  string             `gorm:"index"`
	TaskChallenge    task.TaskChallenge `gorm:"foreignKey:TaskChallengeId"`
	StatusProgress   string             `gorm:"type:enum('in_progress', 'done');default:'in_progress'"`
	StatusAccept     string             `gorm:"type:enum('accept','need_rivew', 'reject');default:'need_rivew'"`
	ImageTask        []UserTaskImage    `gorm:"foreignKey:UserTaskChallengeId"`
	DescriptionImage string
	Point            int
	Reason           string
	acceptedAt       time.Time
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type UserTaskImage struct {
	ID                  int    `gorm:"primaryKey"`
	UserTaskChallengeId string `gorm:"index"`
	ImageUrl            string
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
