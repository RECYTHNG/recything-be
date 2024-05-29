package service_test

import (
	"errors"
	"mime/multipart"
	"testing"
	"time"

	"github.com/sawalreverr/recything/internal/feature/article/entity"
	"github.com/sawalreverr/recything/internal/feature/article/service"
	"github.com/sawalreverr/recything/mocks"
	"github.com/sawalreverr/recything/pagination"
	"github.com/stretchr/testify/assert"
)

func TestDeleteArticleSucces(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	articleID := "1234ABDC"

	repoData.On("DeleteArticle", articleID).Return(nil)
	err := articleService.DeleteArticle(articleID)

	assert.NoError(t, err)
}

func TestDeleteArticleEmptyID(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	err := articleService.DeleteArticle("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id artikel tidak ditemukan")
}

func TestDeleteArticleRepositoryError(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	articleID := "12345abc"

	repoData.On("DeleteArticle", articleID).Return(errors.New("repository error"))
	err := articleService.DeleteArticle(articleID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gagal menghapus artikel")
}

func TestGetSpecificArticleSuccess(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticle := entity.ArticleCore{
		ID:          "12345abc",
		Title:       "Sample Article",
		Description: "ini isi dari artikel",
	}

	repoData.On("GetSpecificArticle", mockArticle.ID).Return(mockArticle, nil)
	articleData, err := articleService.GetSpecificArticle(mockArticle.ID)

	assert.NoError(t, err)
	assert.Equal(t, mockArticle, articleData)
}

func TestGetSpecificArticleEmptyID(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	articleData, err := articleService.GetSpecificArticle("")

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, articleData)
	assert.Contains(t, err.Error(), "id tidak cocok")
}

func TestGetSpecificArticleRepositoryError(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"

	repoData.On("GetSpecificArticle", mockArticleID).Return(entity.ArticleCore{}, errors.New("repository error"))
	articleData, err := articleService.GetSpecificArticle(mockArticleID)

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, articleData)
	assert.Contains(t, err.Error(), "gagal membaca data")
}

func TestUpdateArticleSuccess(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"
	mockArticleInput := entity.ArticleCore{
		Title:       "Updated Article",
		Description: "Updated Article Description",
		Category_id: []string{"123", "456"},
	}

	repoData.On("UpdateArticle", mockArticleID, mockArticleInput, (*multipart.FileHeader)(nil)).Return(mockArticleInput, nil)
	updatedArticle, err := articleService.UpdateArticle(mockArticleID, mockArticleInput, nil)

	assert.NoError(t, err)
	assert.Equal(t, mockArticleInput, updatedArticle)
}

func TestUpdateArticleEmptyID(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	updatedArticle, err := articleService.UpdateArticle("", entity.ArticleCore{}, nil)

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, updatedArticle)
	assert.Contains(t, err.Error(), "id tidak ditemukan")
}

func TestUpdateArticleInvalidData(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"
	invalidArticleInput := entity.ArticleCore{
		Title: "Updated Article",
	}

	updatedArticle, err := articleService.UpdateArticle(mockArticleID, invalidArticleInput, nil)

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, updatedArticle)
	assert.Contains(t, err.Error(), "artikel tidak boleh kosong")
}

func TestUpdateArticleInvalidCategory(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"
	invalidCategoryInput := entity.ArticleCore{
		Title:       "Updated Article",
		Description: "Updated Article Description",
		Category_id: []string{},
	}

	updatedArticle, err := articleService.UpdateArticle(mockArticleID, invalidCategoryInput, nil)

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, updatedArticle)
	assert.Contains(t, err.Error(), "kategori tidak boleh kosong")
}

func TestUpdateLargeImage(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"
	mockArticleInput := entity.ArticleCore{
		Title:       "Updated Article",
		Description: "Updated Article Description",
		Category_id: []string{"123", "456"},
	}

	//large image size
	largeImage := &multipart.FileHeader{
		Size: 6 * 1024 * 1024,
	}

	updateArticle, err := articleService.UpdateArticle(mockArticleID, mockArticleInput, largeImage)

	assert.Error(t, err)
	assert.Equal(t, entity.ArticleCore{}, updateArticle)
	assert.Contains(t, err.Error(), "ukuran gambar tidak boleh lebih dari 5 MB")

}

func TestUpdateArticleRepositoryError(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleID := "12345abc"
	mockArticleInput := entity.ArticleCore{
		Title:       "Updated Article",
		Description: "Updated Article Description",
		Category_id: []string{"123", "456"},
	}

	mockImage := &multipart.FileHeader{
		Size: 6 * 1024 * 1024,
	}

	repoData.On("UpdateArticle", mockArticleID, mockArticleInput, mockImage).Return(entity.ArticleCore{}, errors.New("repository error"))
	updatedArticle, err := articleService.UpdateArticle(mockArticleID, mockArticleInput, mockImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "repository error")
	assert.Equal(t, entity.ArticleCore{}, updatedArticle)
}

func TestGetAllArticleInvalidLimit(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockPage := 1
	mockLimit := 15
	mockSearch := "Sample Search"
	mockFilter := "Article"

	articles, pageInfo, count, err := articleService.GetAllArticle(mockPage, mockLimit, mockSearch, mockFilter)

	assert.Error(t, err)
	assert.Empty(t, articles)
	assert.Equal(t, pagination.Pageinfo{}, pageInfo)
	assert.Equal(t, 0, count)
	assert.Contains(t, err.Error(), "limit tidak boleh lebih dari 10")
}

func TestGetAllArticleInvalidCategory(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockPage := 1
	mockLimit := 5
	mockSearch := "Sample Search"
	mockFilter := "Invalid Category"

	articles, pageInfo, count, err := articleService.GetAllArticle(mockPage, mockLimit, mockSearch, mockFilter)

	assert.Error(t, err)
	assert.Empty(t, articles)
	assert.Equal(t, pagination.Pageinfo{}, pageInfo)
	assert.Equal(t, 0, count)
	assert.Contains(t, err.Error(), "error : kategori tidak valid")
}

func TestGetAllArticleRepositoryError(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockPage := 1
	mockLimit := 5
	mockSearch := "Sample Search"
	mockFilter := "Article"

	repoData.On("GetAllArticle", mockPage, mockLimit, mockSearch, mockFilter).Return([]entity.ArticleCore{}, pagination.Pageinfo{}, 0, errors.New("repository error"))

	articles, pageInfo, count, err := articleService.GetAllArticle(mockPage, mockLimit, mockSearch, mockFilter)

	assert.Error(t, err)
	assert.Empty(t, articles)
	assert.Equal(t, pagination.Pageinfo{}, pageInfo)
	assert.Equal(t, 0, count)
	assert.Contains(t, err.Error(), "kategori tidak valid")
}

func TestCreateArticleSuccess(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleInput := entity.ArticleCore{
		Title:       "Sample Title",
		Description: "Sample Description",
		Category_id: []string{"1", "2"},
		Image:       "sample_image.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockImage := &multipart.FileHeader{
		Size: 1 * 1024 * 1024,
	}

	repoData.On("CreateArticle", mockArticleInput, mockImage).Return(mockArticleInput, nil)
	article, err := articleService.CreateArticle(mockArticleInput, mockImage)

	assert.NoError(t, err)
	assert.Equal(t, mockArticleInput, article)
}

func TestCreateArticleEmptyTitleAndContent(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleInput := entity.ArticleCore{}

	mockImage := &multipart.FileHeader{
		Size: 1 * 1024 * 1024,
	}

	article, err := articleService.CreateArticle(mockArticleInput, mockImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "judul dan konten artikel tidak boleh kosong")
	assert.Equal(t, entity.ArticleCore{}, article)
}

func TestCreateArticleEmptyCategory(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleInput := entity.ArticleCore{
		Title:       "Sample Title",
		Description: "Sample Description",
	}

	mockImage := &multipart.FileHeader{
		Size: 1 * 1024 * 1024,
	}

	article, err := articleService.CreateArticle(mockArticleInput, mockImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "kategori tidak boleh kosong")
	assert.Equal(t, entity.ArticleCore{}, article)
}

func TestCreateArticleLargeImage(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleInput := entity.ArticleCore{
		Title:       "Sample Title",
		Description: "Sample Description",
		Category_id: []string{"1", "2"},
	}

	mockImage := &multipart.FileHeader{
		Size: 6 * 1024 * 1024,
	}

	article, err := articleService.CreateArticle(mockArticleInput, mockImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ukuran file tidak boleh lebih dari 5 MB")
	assert.Equal(t, entity.ArticleCore{}, article)
}

func TestCreateArticleRepositoryError(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	mockArticleInput := entity.ArticleCore{
		Title:       "Updated Title",
		Description: "Updated Description",
		Category_id: []string{"1", "2"},
	}

	mockImage := &multipart.FileHeader{
		Size: 1 * 1024 * 1024,
	}

	repoData.On("CreateArticle", mockArticleInput, mockImage).Return(entity.ArticleCore{}, errors.New("repository error"))
	CreateArticle, err := articleService.CreateArticle(mockArticleInput, mockImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "repository error")
	assert.Equal(t, entity.ArticleCore{}, CreateArticle)
}
