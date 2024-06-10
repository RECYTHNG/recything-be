package article

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type articleHandler struct {
	usecase art.ArticleUsecase
}

func NewArticleHandler(uc art.ArticleUsecase) art.ArticleHandler {
	return &articleHandler{usecase: uc}
}

func (h *articleHandler) GetArticleByID(c echo.Context) error {
	articleId := c.Param("articleId")

	articleFound, err := h.usecase.GetArticleByID(articleId)
	if err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", articleFound)
}
