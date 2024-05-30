package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	Category  string
	ImageUrl  string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
