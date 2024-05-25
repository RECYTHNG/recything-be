package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id          string `gorm:"primary_key"`
	Title       string
	Description string
	CreatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ArticleTrashCateory struct {
	TrashCategoryID string
	ArticleID       string
}
