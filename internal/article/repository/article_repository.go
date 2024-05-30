package repository

import "github.com/sawalreverr/recything/internal/article/entity"

type ArticleRepository interface {
	FindArticleByTitle(title string) (*entity.Article, error)
	FindArticleByID(id string) (*entity.Article, error)
	FindLastIdArticle() (string, error)
	CreateDataArticle(article *entity.Article) (*entity.Article, error)
	GetDataAllArticle(limit int, offset int) ([]entity.Article, int, error)
	UpdateDataArticle(article *entity.Article, id string) (*entity.Article, error)
	DeleteArticle(id string) error
}
