package usecase

import (
	"io"

	// "github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/entity"
)

type ArticleUsecase interface {
	AddArticleUsecase(request dto.ArticleRequestCreate, file io.Reader) (entity.Article, error)
	GetDataAllArticleUsecase(limit, page int) ([]entity.Article, int, error)
	GetDataArticleByIdUsecase(id string) (entity.Article, error)
	UpdateArticleUsecase(request dto.ArticleUpdateRequest, id string, file io.Reader) (entity.Article, error)
	DeleteArticleUsecase(id string) error
}
