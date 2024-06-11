package entity

import (
	"time"

	"github.com/sawalreverr/recything/internal/user"
	"gorm.io/gorm"
)

type Video struct {
	ID              int `gorm:"primaryKey"`
	Title           string
	Description     string
	Thumbnail       string
	Link            string
	Viewer          int
	VideoCategories []VideoCategory `gorm:"foreignKey:VideoID"`
	TrashCategories []TrashCategory `gorm:"foreignKey:VideoID"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt  `gorm:"index"`
}

type TrashCategory struct {
	ID        int `gorm:"primaryKey"`
	VideoID   int `gorm:"index"`
	Name      string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type VideoCategory struct {
	ID        int `gorm:"primaryKey"`
	VideoID   int `gorm:"index"`
	Name      string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	VideoID   int       `gorm:"index"`
	Video     Video     `gorm:"foreignKey:VideoID"`
	UserID    string    `gorm:"index"`
	User      user.User `gorm:"foreignKey:UserID"`
	Comment   string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
