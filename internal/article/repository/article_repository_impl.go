package repository

import (
	"github.com/sawalreverr/recything/internal/article/entity"
	"gorm.io/gorm"
)

type articleRepositoryImpl struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepositoryImpl{DB: db}
}

func (repo *articleRepositoryImpl) FindArticleByTitle(title string) (*entity.Article, error) {
	var article entity.Article
	if err := repo.DB.Where("title = ?", title).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (repo *articleRepositoryImpl) FindArticleByID(id string) (*entity.Article, error) {
	var article entity.Article
	if err := repo.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (repo *articleRepositoryImpl) FindLastIdArticle() (string, error) {
	var article entity.Article
	if err := repo.DB.Last(&article).Error; err != nil {
		return "", err
	}
	return article.ID, nil
}

func (repo *articleRepositoryImpl) CreateDataArticle(article *entity.Article) (*entity.Article, error) {
	if err := repo.DB.Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (repo *articleRepositoryImpl) GetDataAllArticle(limit int, offset int) ([]entity.Article, int, error) {
	var articles []entity.Article
	var totalCount int64
	if err := repo.DB.Offset(offset).Limit(limit).Find(&articles).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	return articles, int(totalCount), nil
}

func (repo *articleRepositoryImpl) UpdateDataArticle(article *entity.Article, id string) (*entity.Article, error) {
	if err := repo.DB.Model(&entity.Article{}).Where("id = ?", id).Updates(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (repo *articleRepositoryImpl) DeleteArticle(id string) error {
	if err := repo.DB.Where("id = ?", id).Delete(&entity.Article{}).Error; err != nil {
		return err
	}
	return nil
}
