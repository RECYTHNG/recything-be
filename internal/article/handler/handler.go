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

func (h *articleHandler) NewArticle(c echo.Context) error {
	var request art.ArticleInput

	authorID := c.Get("user").(*helper.JwtCustomClaims).UserID

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	response, err := h.usecase.NewArticle(request, authorID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "ok", response)
}

func (h *articleHandler) UpdateArticle(c echo.Context) error {
	var request art.ArticleInput
	articleID := c.Param("articleId")

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Update(articleID, request); err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", nil)
}

func (h *articleHandler) DeleteArticle(c echo.Context) error {
	articleID := c.Param("articleId")

	if err := h.usecase.Delete(articleID); err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", nil)
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

	sortBy := c.QueryParam("sort_by")
	sortType := c.QueryParam("sort_type")

	if sortBy == "" {
		sortBy = "created_at"
		sortType = "asc"
	}

	if sortType == "" {
		sortType = "asc"
	}

	response, err := h.usecase.GetAllArticle(page, limit, sortBy, sortType)
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
