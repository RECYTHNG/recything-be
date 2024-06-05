package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type ArticleHandlerImpl struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleHandler(articleUsecase usecase.ArticleUsecase) *ArticleHandlerImpl {
	return &ArticleHandlerImpl{Usecase: articleUsecase}
}

// Create Article
func (handler *ArticleHandlerImpl) CreateArticleHandler(c echo.Context) error {
	var request dto.ArticleRequestCreate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "image is required")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "failed to open file: "+errOpen.Error())
	}
	defer src.Close()

	article, errUc := handler.Usecase.AddArticleUsecase(request, src)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrTitleAlreadyExists) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTitleAlreadyExists.Error())
		}

		if errors.Is(errUc, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+errUc.Error())
	}

	data := dto.ArticleResponseRegister{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Image:   article.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusCreated, responseData)
}

// Get All Articles
func (handler *ArticleHandlerImpl) GetDataAllArticleHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}

	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil || limitInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}

	articles, totalData, err := handler.Usecase.GetDataAllArticleUsecase(limitInt, pageInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	data := []dto.ArticleDataGetAll{}
	for _, article := range articles {
		data = append(data, dto.ArticleDataGetAll{
			Id:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Image:   article.ImageUrl,
		})
	}

	totalPage := totalData / limitInt
	if totalData%limitInt != 0 {
		totalPage++
	}

	dataRes := dto.ArticleResponseGetDataAll{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return c.JSON(http.StatusOK, dataRes)
}

// Get Article by ID
func (handler *ArticleHandlerImpl) GetDataArticleByIdHandler(c echo.Context) error {
	id := c.Param("articleId")

	article, err := handler.Usecase.GetDataArticleByIdUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrArticleNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	data := dto.ArticleResponseGetDataById{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Image:   article.ImageUrl,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

// Update Article
func (handler *ArticleHandlerImpl) UpdateArticleHandler(c echo.Context) error {
	id := c.Param("articleId")

	var request dto.ArticleUpdateRequest
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "image is required")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "failed to open file: "+errOpen.Error())
	}
	defer src.Close()

	article, errUc := handler.Usecase.UpdateArticleUsecase(request, id, src)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrArticleNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrArticleNotFound.Error())
		}

		if errors.Is(errUc, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+errUc.Error())
	}

	data := dto.ArticleResponseUpdate{
		Id:      id,
		Title:   article.Title,
		Content: article.Content,
		Image:   article.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusOK, "data successfully updated", data)
	return c.JSON(http.StatusOK, responseData)
}

// Delete Article
func (handler *ArticleHandlerImpl) DeleteArticleHandler(c echo.Context) error {
	id := c.Param("articleId")

	err := handler.Usecase.DeleteArticleUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrArticleNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	responseData := helper.ResponseData(http.StatusOK, "data successfully deleted", nil)
	return c.JSON(http.StatusOK, responseData)
}
