package handler

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	"github.com/sawalreverr/recything/internal/video/manage_video/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ManageVideoHandlerImpl struct {
	ManageVideoUsecase usecase.ManageVideoUsecase
}

func NewManageVideoHandlerImpl(manageVideoUsecase usecase.ManageVideoUsecase) *ManageVideoHandlerImpl {
	return &ManageVideoHandlerImpl{
		ManageVideoUsecase: manageVideoUsecase,
	}
}

func (handler *ManageVideoHandlerImpl) CreateDataVideoHandler(c echo.Context) error {
	var request dto.CreateDataVideoRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body, detail+"+err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := handler.ManageVideoUsecase.CreateDataVideoUseCase(&request); err != nil {
		if errors.Is(err, pkg.ErrVideoTitleAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoTitleAlreadyExist.Error())
		}
		if errors.Is(err, pkg.ErrVideoCategoryNotFound) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoCategoryNotFound.Error())
		}
		if errors.Is(err, pkg.ErrNoVideoIdFoundOnUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrNoVideoIdFoundOnUrl.Error())
		}
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoNotFound.Error())
		}
		if errors.Is(err, pkg.ErrVideoService) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrVideoService.Error())
		}
		if errors.Is(err, pkg.ErrApiYouTube) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrApiYouTube.Error())
		}
		if errors.Is(err, pkg.ErrParsingUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrParsingUrl.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	return helper.ResponseHandler(c, http.StatusCreated, "success create data video", nil)
}

func (handler *ManageVideoHandlerImpl) UploadThumbnailVideoHandler(c echo.Context) error {
	file, errFile := c.FormFile("thumbnail")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "thumbnail is required")
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

	imageUrl, err := helper.UploadToCloudinary(src, "video_thumbnail")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
	}

	data := dto.UploadThumbnailResponse{UrlThumbnail: imageUrl}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *ManageVideoHandlerImpl) CreateCategoryVideoHandler(c echo.Context) error {
	var request dto.CreateCategoryVideoRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body, detail+"+err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := handler.ManageVideoUsecase.CreateCategoryVideoUseCase(&request); err != nil {
		if errors.Is(err, pkg.ErrVideoCategoryNameAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoCategoryNameAlreadyExist.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	return helper.ResponseHandler(c, http.StatusCreated, "success create category video", nil)
}

func (handler *ManageVideoHandlerImpl) GetAllCategoryVideoHandler(c echo.Context) error {
	categories, err := handler.ManageVideoUsecase.GetAllCategoryVideoUseCase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	var dataCategories []*dto.DataCategory
	data := &dto.GetAllCategoryVideoResponse{
		Data: []*dto.DataCategory{},
	}

	for _, category := range categories {
		dataCategories = append(dataCategories, &dto.DataCategory{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	data.Data = dataCategories
	responseData := helper.ResponseData(http.StatusOK, "success", data.Data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageVideoHandlerImpl) GetAllDataVideoPaginationHandler(c echo.Context) error {
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
	videos, totalData, err := handler.ManageVideoUsecase.GetAllDataVideoPaginationUseCase(limitInt, pageInt)

	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var dataVideos []*dto.DataVideo
	for _, video := range videos {
		dataVideos = append(dataVideos, &dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
		})
	}

	data := &dto.GetAllDataVideoPaginationResponse{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      dataVideos,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: int(math.Ceil(float64(totalData) / float64(limitInt))),
	}

	return c.JSON(http.StatusOK, data)
}

func (handler *ManageVideoHandlerImpl) GetDetailsDataVideoByIdHandler(c echo.Context) error {
	id := c.Param("videoId")
	idInt, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}
	video, err := handler.ManageVideoUsecase.GetDetailsDataVideoByIdUseCase(idInt)
	if err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	var dataVideo *dto.GetDetailsDataVideoByIdResponse
	dataVideo = &dto.GetDetailsDataVideoByIdResponse{
		Id:           video.ID,
		Title:        video.Title,
		Description:  video.Description,
		UrlThumbnail: video.Thumbnail,
		LinkVideo:    video.Link,
		Viewer:       video.View,
		Category:     dto.DataCategory{Id: video.Category.ID, Name: video.Category.Name},
	}
	responseData := helper.ResponseData(http.StatusOK, "success", dataVideo)
	return c.JSON(http.StatusOK, responseData)
}
