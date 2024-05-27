package response

import "github.com/sawalreverr/recything/internal/article/entity"

func CategoryCoreToCategoryResponse(category entity.ArticleTrashCategoryCore) TrashCategoryResponse {
	return TrashCategoryResponse{
		Category: category.Category,
	}
}

func ListCategoryCoreToCategoryResponse(categories []entity.ArticleTrashCategoryCore) []TrashCategoryResponse {
	ResponseCategory := []TrashCategoryResponse{}
	for _, v := range categories {
		category := CategoryCoreToCategoryResponse(v)
		ResponseCategory = append(ResponseCategory, category)
	}
	return ResponseCategory
}

func ArticleCoreToArticleResponse(article entity.ArticleCore) ArticleCreateResponse {
	articleResp := ArticleCreateResponse{
		Id:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Category_id: article.Category_id,
		Thumbnail:   article.Thumbnail,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
	category := ListCategoryCoreToCategoryResponse(article.Categories)
	articleResp.Categories = category
	return articleResp
}

func ListArticleCoreToListArticleResponse(articles []entity.ArticleCore) []ArticleCreateResponse {
	articleResp := []ArticleCreateResponse{}
	for _, article := range articles {
		articlesData := ArticleCoreToArticleResponse(article)
		articleResp = append(articleResp, articlesData)
	}
	return articleResp
}
