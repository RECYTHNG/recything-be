package entity

import (
	"time"

	"gorm.io/gorm"
)

type Achievement struct {
	ID          int `json:"id" gorm:"primaryKey"`
	Level       string
	TargetPoint int
	BadgeUrl    string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
