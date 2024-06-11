package article

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *articleHandler) GetAllArticle(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	response, err := h.usecase.GetAllArticle(page, limit)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}

func (h *articleHandler) GetArticleByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	response, err := h.usecase.GetArticleByKeyword(keyword)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}

func (h *articleHandler) GetArticleByCategory(c echo.Context) error {
	categoryType := c.QueryParam("type")
	categoryName := c.QueryParam("name")

	response, err := h.usecase.GetArticleByCategory(categoryName, categoryType)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
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
