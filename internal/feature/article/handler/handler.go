package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/request"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/article/dto/response"
	"github.com/sawalreverr/recything/internal/article/entity"

	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type articleHandler struct {
	articleService entity.ArticleServiceInterface
}

func NewArticleHandler(article entity.ArticleServiceInterface) *articleHandler {
	return &articleHandler{
		articleService: article,
	}
}

func (a *articleHandler) CreateArticle(e echo.Context) error {
	Id, role, _ := helper.JwtCustomClaims(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))

	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}
	if role == "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("role tidak sesuai, akses di tolak"))
	}

	newArticle := request.ArticleRequest{}
	err := e.Bind(&newArticle)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	thumbnail, err := e.FormFile("thumbnail")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse("tidak ada file yang di upload"))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	articleInput := request.ArticleRequestToArticleCore(newArticle)
	_, errCreate := a.articleService.CreateArticle(articleInput, thumbnail)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), pkg.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse("kategori tidak ditemukan"))

		}
		if strings.Contains(errCreate.Error(), pkg.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan artikel"))

}

func (a *articleHandler) GetAllArticle(e echo.Context) error {
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	Id, _, err := helper.JwtCustomClaims(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(pkg.ERROR_ID_INVALID))
	}

	articleData, paginationInfo, count, err := a.articleService.GetAllArticle(page, limit, search, "")
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan artikel"))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan semua article", articleResponse, paginationInfo, count))
}

func (article *articleHandler) GetAllArticleUser(e echo.Context) error {
	filter := e.QueryParam("filter")
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	Id, _, err := helper.JwtCustomClaims(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(pkg.ERROR_ID_INVALID))
	}

	articleData, paginationInfo, count, err := article.articleService.GetAllArticle(page, limit, search, filter)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if len(articleData) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(pkg.SUCCESS_NULL))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan semua article", articleResponse, paginationInfo, count))
}

func (article *articleHandler) UpdateArticle(e echo.Context) error {
	Id, role, _ := helper.JwtCustomClaims(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idParams := e.Param("id")

	updatedData := request.ArticleRequest{}
	errBind := e.Bind(&updatedData)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	image, _ := e.FormFile("image")

	articleInput := request.ArticleRequestToArticleCore(updatedData)
	updateArticle, errUpdate := article.articleService.UpdateArticle(idParams, articleInput, image)
	if errUpdate != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errUpdate.Error()))
	}

	articleResponse := response.ArticleCoreToArticleResponse(updateArticle)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil", articleResponse))
}

func (article *articleHandler) DeleteArticle(e echo.Context) error {
	Id, role, _ := helper.JwtCustomClaims(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idParams := e.Param("id")

	errDelete := article.articleService.DeleteArticle(idParams)
	if errDelete != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errDelete.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus artikel"))
}
