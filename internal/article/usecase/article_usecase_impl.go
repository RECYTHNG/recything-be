package usecase

import (
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/entity"
	"github.com/sawalreverr/recything/internal/article/repository"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ArticleUsecaseImpl struct {
	Repository repository.ArticleRepository
	Validate   *validator.Validate
}

func NewArticleUsecase(articleRepo repository.ArticleRepository) *ArticleUsecaseImpl {
	return &ArticleUsecaseImpl{Repository: articleRepo}
}

func (usecase *ArticleUsecaseImpl) AddArticleUsecase(request dto.ArticleRequestCreate, file io.Reader) (*entity.Article, error) {
	findArticle, _ := usecase.Repository.FindArticleByTitle(request.Title)
	if findArticle != nil {
		return nil, pkg.ErrTitleAlreadyExists
	}

	imageUrl, errUpload := helper.UploadToCloudinary(file, "article_image")
	if errUpload != nil {
		return nil, pkg.ErrUploadCloudinary
	}

	findLastId, _ := usecase.Repository.FindLastIdArticle()
	id := helper.GenerateCustomID(findLastId, "AR")

	article := &entity.Article{
		ID:        id,
		Title:     request.Title,
		Content:   request.Content,
		ImageUrl:  imageUrl,
		DeletedAt: gorm.DeletedAt{},
	}

	if _, err := usecase.Repository.CreateDataArticle(article); err != nil {
		return nil, err
	}
	return article, nil
}

func (usecase *ArticleUsecaseImpl) GetDataAllArticleUsecase(limit int, offset int) ([]entity.Article, int, error) {
	articles, totalCount, err := usecase.Repository.GetDataAllArticle(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	fmt.Println("data article", articles)
	return articles, totalCount, nil
}

func (usecase *ArticleUsecaseImpl) GetDataArticleByIdUsecase(id string) (*entity.Article, error) {
	article, err := usecase.Repository.FindArticleByID(id)
	if err != nil {
		return nil, pkg.ErrArticleNotFound
	}
	return article, nil
}

func (usecase *ArticleUsecaseImpl) UpdateArticleUsecase(request dto.ArticleUpdateRequest, id string, file io.Reader) (*entity.Article, error) {
	findArticle, _ := usecase.Repository.FindArticleByID(id)
	if findArticle == nil {
		return nil, pkg.ErrArticleNotFound
	}

	imageUrl, errUpload := helper.UploadToCloudinary(file, "article_image_update")
	if errUpload != nil {
		return nil, pkg.ErrUploadCloudinary
	}

	article, error := usecase.Repository.UpdateDataArticle(&entity.Article{
		Title:     request.Title,
		Content:   request.Content,
		ImageUrl:  imageUrl,
		DeletedAt: gorm.DeletedAt{},
	}, id)
	if error != nil {
		return nil, error
	}
	return article, nil
}

func (usecase *ArticleUsecaseImpl) DeleteArticleUsecase(id string) error {
	findArticle, _ := usecase.Repository.FindArticleByID(id)
	if findArticle == nil {
		return pkg.ErrArticleNotFound
	}
	if err := usecase.Repository.DeleteArticle(id); err != nil {
		return err
	}
	return nil
}
