package entity

import (
	"mime/multipart"

	"github.com/sawalreverr/recything/pagination"
)

type ArticleRepositoryInterface interface {
	CreateArticle(articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search, filter string) ([]ArticleCore, pagination.Pageinfo, int, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}

type ArticleServiceInterface interface {
	CreateArticle(articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search, filter string) ([]ArticleCore, pagination.Pageinfo, int, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}
