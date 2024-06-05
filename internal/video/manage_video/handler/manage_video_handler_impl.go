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
	return nil
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
