package repository

import (
	"github.com/sawalreverr/recything/internal/article/entity"
	"github.com/sawalreverr/recything/internal/database"
)

// ArticleRepository defines the interface for article repository

// articleRepositoryImpl implements the ArticleRepository interface
type articleRepositoryImpl struct {
	DB database.Database
}

// NewArticleRepository returns a new instance of articleRepositoryImpl
func NewArticleRepository(db database.Database) ArticleRepository {
	return &articleRepositoryImpl{DB: db}
}

// FindArticleByTitle retrieves an article by its title
func (repo *articleRepositoryImpl) FindArticleByTitle(title string) (*entity.Article, error) {
	var article entity.Article
	if err := repo.DB.GetDB().Where("title = ?", title).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// FindArticleByID retrieves an article by its ID
func (repo *articleRepositoryImpl) FindArticleByID(id string) (*entity.Article, error) {
	var article entity.Article
	if err := repo.DB.GetDB().Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// FindLastIdArticle retrieves the last inserted article's ID
func (repo *articleRepositoryImpl) FindLastIdArticle() (string, error) {
	var article entity.Article
	if err := repo.DB.GetDB().Order("id DESC").First(&article).Error; err != nil {
		return "", err
	}
	return article.ID, nil
}

// CreateDataArticle creates a new article
func (repo *articleRepositoryImpl) CreateDataArticle(article *entity.Article) (*entity.Article, error) {
	if err := repo.DB.GetDB().Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// GetDataAllArticle retrieves all articles with pagination
func (repo *articleRepositoryImpl) GetDataAllArticle(limit int, offset int) ([]entity.Article, int, error) {
	var articles []entity.Article
	var totalCount int64
	if err := repo.DB.GetDB().Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	if err := repo.DB.GetDB().Model(&entity.Article{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	return articles, int(totalCount), nil
}

// UpdateDataArticle updates an existing article
func (repo *articleRepositoryImpl) UpdateDataArticle(article *entity.Article, id string) (*entity.Article, error) {
	if err := repo.DB.GetDB().Model(&entity.Article{}).Where("id = ?", id).Updates(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// DeleteArticle deletes an article by its ID
func (repo *articleRepositoryImpl) DeleteArticle(id string) error {
	if err := repo.DB.GetDB().Where("id = ?", id).Delete(&entity.Article{}).Error; err != nil {
		return err
	}
	return nil
}
