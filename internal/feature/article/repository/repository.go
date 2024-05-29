package repository

import (
	"errors"
	"mime/multipart"

	"github.com/sawalreverr/recything/internal/feature/article/entity"
	"github.com/sawalreverr/recything/internal/feature/article/model"
	"github.com/sawalreverr/recything/storage"
	"gorm.io/gorm"
)

type articleRepository struct {
	db            *gorm.DB
	trashcategory trashcategory.TrashCategoryRepositoryInterface
}

func NewArticleRepository(db *gorm.DB, trashcategory trashcategory.TrashCategoryRepositoryInterface) entity.ArticleRepositoryInterface {
	return &articleRepository{
		db:            db,
		trashcategory: trashcategory,
	}
}

func (article *articleRepository) DeleteArticle(id string) error {
	checkId := model.Article{}

	tx := article.db.Where("id = ?", id).Delete(&checkId)
	if tx.Error != nil {
		return tx.Error
	}

	categoryId := model.ArticleTrashCategory{}
	categoryDel := article.db.Where("article_id = ?", id).Delete(&categoryId)
	if categoryDel.Error != nil {
		return categoryDel.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("tidak ada data yang dihapus")
	}

	return nil
}

// UpdateArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) UpdateArticle(idArticle string, articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	input := entity.ArticleCoreToArticleModel(articleInput)
	var articleData model.Article

	check := article.db.Where("id = ?", idArticle).First(&articleData)
	if check.Error != nil {
		return entity.ArticleCore{}, check.Error
	}

	if image != nil {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return entity.ArticleCore{}, errUpload
		}
		articleData.Image = imageURL

	} else {
		input.Image = articleData.Image
	}

	articleData.Title = articleInput.Title
	articleData.Description = articleInput.Description

	tx := article.db.Begin()

	// Hapus kategori yang terkait dengan artikel
	categoryId := model.ArticleTrashCategory{}
	categoryDel := tx.Where("article_id = ?", idArticle).Delete(&categoryId)
	if categoryDel.Error != nil {
		return entity.ArticleCore{}, categoryDel.Error
	}

	if err := tx.Save(&articleData).Error; err != nil {
		tx.Rollback()
		return entity.ArticleCore{}, err
	}

	// Tambahkan kategori yang baru
	for _, categoryId := range articleInput.Category_id {
		categories := new(model.ArticleTrashCategory)
		categories.ArticleID = idArticle
		categories.TrashCategoryID = categoryId

		txLink := tx.Create(&categories)
		if txLink.Error != nil {
			tx.Rollback()
			return entity.ArticleCore{}, errors.New("kategori tidak ditemukan")
		}
	}

	tx.Commit()

	articleUpdate := entity.ArticleModelToArticleCore(articleData)

	return articleUpdate, nil
}
