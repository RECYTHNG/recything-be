package repository

import (
	"github.com/sawalreverr/recything/internal/feature/article/entity"
	"gorm.io/gorm"
)

type articleRepository struct {
	db            *gorm.DB
	trashcategory trashcategory.TrashCategoryRepositoryInterface
}

func NewArticleRepository(db *gorm.DB, trashcategory trashcategory.TrashCategoryRepositoryInterface) entity.ArticleRepositoryInterface {
	return &articleRepository{
		db:            db,
		trashcategory: trashcategory,
	}
}
