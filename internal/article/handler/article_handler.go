package handler

import "github.com/sawalreverr/recything/internal/article/entity"

type ArticleUsecase interface {
	UploadArticle(article entity.Article) (*entity.Article, error)
	UpdateArticle(articleID string, article entity.Article) error
	DeleteArticle(articleID string) error

	FindArticleByID(articleID string) (*entity.Article, error)
	FindAllArticles(page int, limit int, sortBy string, sortType string) (*[]entity.Article, error)
}
