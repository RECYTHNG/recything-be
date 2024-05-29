package entity

import (
	"time"

	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"gorm.io/gorm"
)

type Admin struct {
	ID             string `gorm:"primaryKey"`
	Name           string
	Email          string
	Password       string
	Role           string `gorm:"type:enum('super admin', 'admin')"`
	ImageUrl       string
	TaskChallanges []task.TaskChallange `gorm:"foreignKey:AdminId"`
	CreatedAt      time.Time            `gorm:"autoCreateTime"`
	UpdatedAt      time.Time            `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt       `gorm:"index"`
}
