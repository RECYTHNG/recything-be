package article

import (
	"errors"

	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type articleRepository struct {
	DB database.Database
}

func NewArticleRepository(db database.Database) art.ArticleRepository {
	return &articleRepository{DB: db}
}

func (r *articleRepository) Create(article art.Article) (*art.Article, error) {
	if err := r.DB.GetDB().Create(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *articleRepository) FindByID(articleID string) (*art.Article, error) {
	var article art.Article
	if err := r.DB.GetDB().Preload("Categories").Preload("Sections").First(&article, "id = ?", articleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrArticleNotFound
		}
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) FindAll(page, limit uint) (*[]art.Article, error) {
	var articles []art.Article
	offset := (page - 1) * limit
	if err := r.DB.GetDB().Preload("Categories").Preload("Sections").Limit(int(limit)).Offset(int(offset)).Find(&articles).Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func (r *articleRepository) FindLastID() (string, error) {
	var article art.Article
	if err := r.DB.GetDB().Unscoped().Order("id DESC").First(&article).Error; err != nil {
		return "ART0000", err
	}

	return article.ID, nil
}

func (r *articleRepository) FindByKeyword(keyword string) (*[]art.Article, error) {
	var articles []art.Article
	query := "%" + keyword + "%"
	if err := r.DB.GetDB().Preload("Categories").Preload("Sections").
		Where("title LIKE ? OR description LIKE ?", query, query).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func (r *articleRepository) FindByCategory(categoryName string) (*[]art.Article, error) {
	var articles []art.Article
	if err := r.DB.GetDB().Preload("Categories").Preload("Sections").
		Joins("JOIN article_categories ON articles.id = article_categories.article_id").
		Joins("JOIN categories ON article_categories.category_id = categories.id").
		Where("categories.name = ?", categoryName).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func (r *articleRepository) Update(article art.Article) error {
	if err := r.DB.GetDB().Save(&article).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) Delete(articleID string) error {
	var article art.Article
	if err := r.DB.GetDB().Delete(&article, "id = ?", articleID).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) FindCategories(articleID string) (*[]art.WasteCategory, *[]art.VideoCategory, error) {
	var articleCategories []art.ArticleCategories
	var wasteCategories []art.WasteCategory
	var contentCategories []art.VideoCategory

	if err := r.DB.GetDB().Where("article_id = ?", articleID).Find(&articleCategories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, pkg.ErrArticleNotFound
		}
		return nil, nil, err
	}

	var wasteCategoryIDs []uint
	var contentCategoryIDs []uint
	for _, ac := range articleCategories {
		if ac.WasteCategoryID != 0 {
			wasteCategoryIDs = append(wasteCategoryIDs, ac.WasteCategoryID)
		}
		if ac.ContentCategoryID != 0 {
			contentCategoryIDs = append(contentCategoryIDs, uint(ac.ContentCategoryID))
		}
	}

	if len(wasteCategoryIDs) > 0 {
		if err := r.DB.GetDB().Where("id IN (?)", wasteCategoryIDs).Find(&wasteCategories).Error; err != nil {
			return nil, nil, err
		}
	}

	if len(contentCategoryIDs) > 0 {
		if err := r.DB.GetDB().Where("id IN (?)", contentCategoryIDs).Find(&contentCategories).Error; err != nil {
			return nil, nil, err
		}
	}

	return &wasteCategories, &contentCategories, nil
}
