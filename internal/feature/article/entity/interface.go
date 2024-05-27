package entity

import "mime/multipart"

type ArticleRepositoryInterface interface {
	CreateArticle(articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search, filter string) ([]ArticleCore, pagination.PageInfo, int, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}

type ArticleServiceInterface interface {
	CreateArticle(articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search, filter string) ([]ArticleCore, pagination.PageInfo, int, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleinput ArticleCore, thumbnail *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}
