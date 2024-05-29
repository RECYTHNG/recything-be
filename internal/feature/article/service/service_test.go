package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteArticleSuccess(t *testing.T) {
	repoData := new(mocks.ArticleRepositoryInterface)
	articleService := service.NewArticleService(repoData)

	articleID := "12345abc"

	repoData.On("DeleteArticle", articleID).Return(nil)
	err := articleService.DeleteArticle(articleID)

	assert.NoError(t, err)
}
