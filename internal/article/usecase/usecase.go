package article

import (
	admin "github.com/sawalreverr/recything/internal/admin/repository"
	art "github.com/sawalreverr/recything/internal/article"
)

type articleUsecase struct {
	articleRepo art.ArticleRepository
	adminRepo   admin.AdminRepository
}

func NewArticleUsecase(articleRepo art.ArticleRepository, adminRepo admin.AdminRepository) art.ArticleUsecase {
	return &articleUsecase{articleRepo: articleRepo, adminRepo: adminRepo}
}

func (uc *articleUsecase) NewArticle(article art.ArticleInput) (*art.ArticleDetail, error) {

	return nil, nil
}

func (uc *articleUsecase) GetArticleByID(articleID string) (*art.ArticleDetail, error) {
	articleFound, err := uc.articleRepo.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	articleDetail := uc.GetArticleDetail(*articleFound)

	return articleDetail, nil
}

func (uc *articleUsecase) GetAllArticle(page, limit uint) (*art.ArticleResponsePagination, error) {

	return nil, nil
}

func (uc *articleUsecase) GetArticleByKeyword(keyword string) (*[]art.ArticleDetail, error) {

	return nil, nil
}

func (uc *articleUsecase) GetArticleByCategory(categoryName string) (*[]art.ArticleDetail, error) {

	return nil, nil
}

func (uc *articleUsecase) Update(articleID string, article art.ArticleInput) error {

	return nil
}

func (uc *articleUsecase) Delete(articleID string) error {

	return nil
}

func (uc *articleUsecase) GetArticleDetail(article art.Article) *art.ArticleDetail {
	adminDetail, _ := uc.GetDetailAuthor(article.AuthorID)
	categories, _ := uc.articleRepo.FindCategories(article.ID)

	articleDetail := art.ArticleDetail{
		ID:           article.ID,
		AuthorID:     *adminDetail,
		Title:        article.Title,
		Description:  article.Description,
		ThumbnailURL: article.ThumbnailURL,
		CreatedAt:    article.CreatedAt,
		Categories:   *categories,
		Sections:     article.Sections,
	}

	return &articleDetail
}

func (uc *articleUsecase) GetDetailAuthor(authorID string) (*art.AdminDetail, error) {
	adminFound, err := uc.adminRepo.FindAdminByID(authorID)
	if err != nil {
		return nil, err
	}

	adminDetail := art.AdminDetail{
		ID:       adminFound.ID,
		Name:     adminFound.Name,
		ImageURL: adminFound.ImageUrl,
	}

	return &adminDetail, nil
}
