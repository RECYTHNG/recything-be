package repository

import "github.com/sawalreverr/recything/internal/article/entity"

type ArticleRepository interface {
	CreateArticleRepository(article *entity.Article)
	GetAllArticleRepository(limit, offset int) ([]entity.Article, int, error)
	GetByIDArticleRepository(id string) (entity.Article, error)
	UpdateArticleRepository(article *entity.Article) error
	DeleteArticleRepository(id string) error
}
