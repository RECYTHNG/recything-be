package service

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/sawalreverr/recything/internal/feature/article/entity"
)

type articleService struct {
	ArticleRepository entity.ArticleRepositoryInterface
}

func NewArticleService(article entity.ArticleRepositoryInterface) entity.ArticleServiceInterface {
	return &articleService{
		ArticleRepository: article,
	}
}

// DeleteArticle implements entity.ArticleServiceInterface.
func (c *articleService) DeleteArticle(id string) error {
	if id == "" {
		return errors.New("id artikel tidak ditemukan")
	}

	errArticle := c.ArticleRepository.DeleteArticle(id)
	if errArticle != nil {
		return errors.New("gagal menghapus artikel " + errArticle.Error())
	}

	return nil
}

// GetSpecificArticle implements entity.ArticleServiceInterface.
func (c *articleService) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak cocok")
	}

	articleData, err := c.ArticleRepository.GetSpecificArticle(idArticle)
	if err != nil {
		fmt.Println("service", err)
		return entity.ArticleCore{}, errors.New("gagal membaca data")
	}

	return articleData, nil
}

func (article *articleService) UpdateArticle(idArticle string, articleInput entity.ArticleCore, thumbnail *multipart.FileHeader) (entity.ArticleCore, error) {

	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak ditemukan")
	}

	if articleInput.Title == "" || articleInput.Description == "" {
		return entity.ArticleCore{}, errors.New("artikel tidak boleh kosong")
	}

	if len(articleInput.Category_id) == 0 {
		return entity.ArticleCore{}, errors.New("kategori tidak boleh kosong")
	}

	if thumbnail != nil && thumbnail.Size > 5*1024*1024 {
		return entity.ArticleCore{}, errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	articleUpdate, errinsert := article.ArticleRepository.UpdateArticle(idArticle, articleInput, thumbnail)
	if errinsert != nil {
		return entity.ArticleCore{}, errinsert
	}

	return articleUpdate, nil
}

// GetPopularArticle implements entity.ArticleServiceInterface.
// func (article *articleService) GetPopularArticle(search string) ([]entity.ArticleCore, error) {
// 	articleData, err := article.ArticleRepository.GetPopularArticle(search)
// 	if err != nil {
// 		return []entity.ArticleCore{}, err
// 	}

// 	return articleData, nil
// }
