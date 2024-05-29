package entity

import (
	"time"

	"gorm.io/gorm"
)

type TaskChallange struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Tumbnail    string
	StartDate   time.Time
	EndDate     time.Time
	TaskSteps   []TaskStep     `gorm:"foreignKey:TaskChallangeId"`
	AdminId     string         `gorm:"index"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type TaskStep struct {
	ID              int    `gorm:"primaryKey"`
	TaskChallangeId string `gorm:"index"`
	Title           string
	Description     string
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
