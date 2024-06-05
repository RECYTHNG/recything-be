package handler

import (
	"errors"
	"net/http"

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
