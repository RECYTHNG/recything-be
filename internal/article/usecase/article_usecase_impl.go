package usecase

import (
	"errors"
	"io"

	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/entity"
	"github.com/sawalreverr/recything/internal/article/repository"
	"github.com/sawalreverr/recything/pkg"
)

type articleUsecaseImpl struct {
	repo repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecaseImpl{repo: repo}
}

func (u *articleUsecaseImpl) AddArticleUsecase(request dto.ArticleRequestCreate, file io.Reader) (entity.Article, error) {
	// Handle file upload to Cloudinary
	imageUrl, err := pkg.UploadToCloudinary(file)
	if err != nil {
		return entity.Article{}, pkg.ErrUploadCloudinary
	}

	// Create article
	article := entity.Article{
		ID:       pkg.GenerateID(),
		Title:    request.Title,
		Content:  request.Content,
		ImageUrl: imageUrl,
	}

	if err := u.repo.Create(article); err != nil {
		if errors.Is(err, pkg.ErrTitleAlreadyExists) {
			return article, pkg.ErrTitleAlreadyExists
		}
		return article, err
	}

	return article, nil
}

func (u *articleUsecaseImpl) GetDataAllArticleUsecase(limit, page int) ([]entity.Article, int, error) {
	offset := (page - 1) * limit
	articles, total, err := u.repo.GetAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

func (u *articleUsecaseImpl) GetDataArticleByIdUsecase(id string) (entity.Article, error) {
	article, err := u.repo.GetByID(id)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (u *articleUsecaseImpl) UpdateArticleUsecase(request dto.ArticleUpdateRequest, id string, file io.Reader) (entity.Article, error) {
	// Handle file upload to Cloudinary
	imageUrl, err := pkg.UploadToCloudinary(file)
	if err != nil {
		return entity.Article{}, pkg.ErrUploadCloudinary
	}

	// Update article
	article := entity.Article{
		ID:       id,
		Title:    request.Title,
		Content:  request.Content,
		ImageUrl: imageUrl,
	}

	if err := u.repo.Update(article); err != nil {
		if errors.Is(err, entity.ErrArticleNotFound) {
			return article, entity.ErrArticleNotFound
		}
		return article, err
	}

	return article, nil
}

func (u *articleUsecaseImpl) DeleteArticleUsecase(id string) error {
	if err := u.repo.Delete(id); err != nil {
		if errors.Is(err, entity.ErrArticleNotFound) {
			return entity.ErrArticleNotFound
		}
		return err
	}
	return nil
}
