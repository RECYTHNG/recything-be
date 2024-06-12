package article

import (
	admin "github.com/sawalreverr/recything/internal/admin/repository"
	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/helper"
	user "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type articleUsecase struct {
	articleRepo art.ArticleRepository
	adminRepo   admin.AdminRepository
	userRepo    user.UserRepository
}

func NewArticleUsecase(articleRepo art.ArticleRepository, adminRepo admin.AdminRepository, userRepo user.UserRepository) art.ArticleUsecase {
	return &articleUsecase{articleRepo: articleRepo, adminRepo: adminRepo, userRepo: userRepo}
}

func (u *articleUsecase) NewArticle(article art.ArticleInput, authorId string) (*art.ArticleDetail, error) {
	lastID, _ := u.articleRepo.FindLastID()
	newID := helper.GenerateCustomID(lastID, "ART")

	newArticle := art.Article{
		ID:           newID,
		Title:        article.Title,
		Description:  article.Description,
		ThumbnailURL: article.ThumbnailURL,
		AuthorID:     authorId,
	}

	createdArticle, err := u.articleRepo.Create(newArticle)
	if err != nil {
		return nil, err
	}

	for _, section := range article.Sections {
		section.ArticleID = createdArticle.ID
		if err := u.articleRepo.CreateSection(section); err != nil {
			_ = u.articleRepo.Delete(createdArticle.ID)
			return nil, err
		}
	}

	for _, categoryName := range article.WasteCategories {
		categoryID, err := u.articleRepo.FindCategoryByName(categoryName, "waste")
		if err != nil {
			_ = u.articleRepo.Delete(createdArticle.ID)
			return nil, err
		}

		articleCategory := art.ArticleCategories{
			ArticleID:       createdArticle.ID,
			WasteCategoryID: categoryID,
		}

		if err := u.articleRepo.CreateArticleCategory(articleCategory); err != nil {
			_ = u.articleRepo.Delete(createdArticle.ID)
			return nil, err
		}
	}

	for _, categoryName := range article.ContentCategories {
		categoryID, err := u.articleRepo.FindCategoryByName(categoryName, "content")
		if err != nil {
			_ = u.articleRepo.Delete(createdArticle.ID)
			return nil, err
		}

		articleCategory := art.ArticleCategories{
			ArticleID:         createdArticle.ID,
			ContentCategoryID: int(categoryID),
		}

		if err := u.articleRepo.CreateArticleCategory(articleCategory); err != nil {
			_ = u.articleRepo.Delete(createdArticle.ID)
			return nil, err
		}
	}

	articleFound, _ := u.articleRepo.FindByID(createdArticle.ID)
	return u.GetArticleDetail(*articleFound), nil
}

func (uc *articleUsecase) GetArticleByID(articleID string) (*art.ArticleDetail, error) {
	articleFound, err := uc.articleRepo.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	return uc.GetArticleDetail(*articleFound), nil
}

func (u *articleUsecase) GetAllArticle(page, limit int, sortBy string, sortType string) (*art.ArticleResponsePagination, error) {
	articles, total, err := u.articleRepo.FindAll(uint(page), uint(limit), sortBy, sortType)
	if err != nil {
		return nil, err
	}

	articleDetails := make([]art.ArticleDetail, len(*articles))
	for i, article := range *articles {
		articleDetails[i] = *u.GetArticleDetail(article)
	}

	return &art.ArticleResponsePagination{
		Total:    total,
		Articles: articleDetails,
		Page:     uint(page),
		Limit:    uint(limit),
	}, nil
}

func (u *articleUsecase) GetArticleByKeyword(keyword string) (*[]art.ArticleDetail, error) {
	articles, err := u.articleRepo.FindByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	articleDetails := make([]art.ArticleDetail, len(*articles))
	for i, article := range *articles {
		articleDetails[i] = *u.GetArticleDetail(article)
	}

	return &articleDetails, nil
}

func (u *articleUsecase) GetArticleByCategory(categoryName string, categoryType string) (*[]art.ArticleDetail, error) {
	articles, err := u.articleRepo.FindByCategory(categoryName, categoryType)
	if err != nil {
		return nil, err
	}

	articleDetails := make([]art.ArticleDetail, len(*articles))
	for i, article := range *articles {
		articleDetails[i] = *u.GetArticleDetail(article)
	}

	return &articleDetails, nil
}

func (u *articleUsecase) Update(articleID string, article art.ArticleInput) error {
	articleFound, err := u.articleRepo.FindByID(articleID)
	if err != nil {
		return err
	}

	articleToUpdate := art.Article{
		ID:           articleFound.ID,
		Title:        article.Title,
		Description:  article.Description,
		ThumbnailURL: article.ThumbnailURL,
		AuthorID:     articleFound.AuthorID,
	}

	if err := u.articleRepo.DeleteAllSection(articleID); err != nil {
		return err
	}

	for _, section := range article.Sections {
		section.ArticleID = articleID
		if err := u.articleRepo.CreateSection(section); err != nil {
			return err
		}
	}

	if err := u.articleRepo.DeleteAllArticleCategory(articleID); err != nil {
		return err
	}

	for _, wasteCategoryName := range article.WasteCategories {
		wasteCategoryID, err := u.articleRepo.FindCategoryByName(wasteCategoryName, "waste")
		if err != nil {
			return err
		}

		articleCategory := art.ArticleCategories{
			ArticleID:       articleID,
			WasteCategoryID: wasteCategoryID,
		}

		if err := u.articleRepo.CreateArticleCategory(articleCategory); err != nil {
			return err
		}
	}

	for _, contentCategoryName := range article.ContentCategories {
		contentCategoryID, err := u.articleRepo.FindCategoryByName(contentCategoryName, "content")
		if err != nil {
			return err
		}

		articleCategory := art.ArticleCategories{
			ArticleID:         articleID,
			ContentCategoryID: int(contentCategoryID),
		}

		if err := u.articleRepo.CreateArticleCategory(articleCategory); err != nil {
			return err
		}
	}

	return u.articleRepo.Update(articleToUpdate)
}

func (uc *articleUsecase) Delete(articleID string) error {
	articleFound, err := uc.articleRepo.FindByID(articleID)
	if err != nil {
		return err
	}

	if err := uc.articleRepo.Delete(articleFound.ID); err != nil {
		return err
	}

	if err := uc.articleRepo.DeleteAllSection(articleFound.ID); err != nil {
		return err
	}

	if err := uc.articleRepo.DeleteAllArticleCategory(articleFound.ID); err != nil {
		return err
	}

	if err := uc.articleRepo.DeleteAllArticleComment(articleFound.ID); err != nil {
		return err
	}

	return nil
}

func (uc *articleUsecase) GetArticleDetail(article art.Article) *art.ArticleDetail {
	adminDetail, _ := uc.GetDetailAuthor(article.AuthorID)
	wasteCategories, contentCategories, _ := uc.articleRepo.FindCategories(article.ID)
	comments, _ := uc.GetDetailComments(article.Comments)

	articleDetail := art.ArticleDetail{
		ID:                article.ID,
		Author:            *adminDetail,
		Title:             article.Title,
		Description:       article.Description,
		ThumbnailURL:      article.ThumbnailURL,
		CreatedAt:         article.CreatedAt,
		WasteCategories:   *wasteCategories,
		ContentCategories: *contentCategories,
		Sections:          article.Sections,
		Comments:          *comments,
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

func (uc *articleUsecase) NewArticleComment(comment art.CommentInput) error {
	articleFound, err := uc.GetArticleByID(comment.ArticleID)
	if err != nil {
		return pkg.ErrArticleNotFound
	}

	userFound, err := uc.userRepo.FindByID(comment.UserID)
	if err != nil {
		return pkg.ErrUserNotFound
	}

	newComment := art.ArticleComment{
		UserID:    userFound.ID,
		ArticleID: articleFound.ID,
		Comment:   comment.Comment,
	}

	if err := uc.articleRepo.CreateArticleComment(newComment); err != nil {
		return err
	}

	return nil
}

func (uc *articleUsecase) GetDetailUser(userID string) (*art.UserDetail, error) {
	userFound, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	userDetail := art.UserDetail{
		ID:       userFound.ID,
		Name:     userFound.Name,
		ImageURL: userFound.PictureURL,
	}

	return &userDetail, nil
}

func (uc *articleUsecase) GetDetailComments(comments []art.ArticleComment) (*[]art.CommentDetail, error) {
	commentDetails := make([]art.CommentDetail, len(comments))
	for i, comment := range comments {
		user, _ := uc.GetDetailUser(comment.UserID)
		detail := art.CommentDetail{
			ID:        comment.ID,
			User:      *user,
			ArticleID: comment.ArticleID,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt,
		}

		commentDetails[i] = detail
	}

	return &commentDetails, nil
}

func (uc *articleUsecase) GetAllCategories() (*art.CategoriesResponse, error) {
	wasteCategories, contentCategories, err := uc.articleRepo.FindAllCategories()
	if err != nil {
		return nil, err
	}

	return &art.CategoriesResponse{
		WasteCategories:   *wasteCategories,
		ContentCategories: *contentCategories,
	}, nil
}
