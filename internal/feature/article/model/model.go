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
	Image       string
	Categories  []model.TrashCategory `gorm:"many2many:ArticleTrashCategory"`
	category_id []string
	CreatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Section struct {
	Id          string
	SubTitle    string `json:"sub_title"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type NewArticle struct {
	Id       string
	Title    string    `json:"title"`
	Image    string    `json:"image"`
	Sections []Section `json:"sections"`
}

type ArticleTrashCategory struct {
	TrashCategoryID string
	ArticleID       string
}
