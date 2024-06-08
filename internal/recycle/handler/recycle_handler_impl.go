package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/recycle/usecase"
)

type RecycleHandlerImpl struct {
	RecycleUsecase usecase.RecycleUsecase
}

func NewRecycleHandlerImpl(usecase usecase.RecycleUsecase) RecycleHandler {
	return &RecycleHandlerImpl{
		RecycleUsecase: usecase,
	}
}

func (handler *RecycleHandlerImpl) GetHomeRecycleHandler(c echo.Context) error {
	data, err := handler.RecycleUsecase.GetHomeRecycleUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)

	return c.JSON(http.StatusOK, responseData)

}

func (handler *RecycleHandlerImpl) SearchVideoHandler(c echo.Context) error {
	title := c.QueryParam("title")
	category := c.QueryParam("category")
	data, err := handler.RecycleUsecase.SearchVideoUsecase(title, category)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data.DataVideo)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *RecycleHandlerImpl) GetAllCategoryVideoHandler(c echo.Context) error {
	data, err := handler.RecycleUsecase.GetAllCategoryVideoUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data.DataCategory)
	return c.JSON(http.StatusOK, responseData)
}
