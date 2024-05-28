package model

import (
	"time"

	"github.com/sawalreverr/recything/internal/feature/trash_category/model"
	"gorm.io/gorm"
)

type Article struct {
	Id          string `gorm:"primary_key"`
	Title       string
	Description string
	Categories  []model.TrashCategory `gorm:"many2many:ArticleTrashCategory"`
	category_id []string
	Thumbnail   string
	CreatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ArticleTrashCateory struct {
	TrashCategoryID string
	ArticleID       string
}
