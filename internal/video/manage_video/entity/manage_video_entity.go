package entity

import (
	"time"

	"github.com/sawalreverr/recything/internal/user"
)

type Video struct {
	ID              string `gorm:"primaryKey"`
	Title           string
	Description     string
	Thumbnail       string
	Link            string
	View            int
	Comments        []Comment     `gorm:"foreignKey:VideoID"`
	VideoCategoryID string        `gorm:"index"`
	Category        VideoCategory `gorm:"foreignKey:VideoCategoryID"`
	CreatedAt       time.Time     `gorm:"autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"autoUpdateTime"`
	DeletedAt       time.Time     `gorm:"index"`
}

type VideoCategory struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}

type Comment struct {
	ID        string `gorm:"primaryKey"`
	VideoID   string `gorm:"index"`
	UserID    string `gorm:"index"`
	User      user.User
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
