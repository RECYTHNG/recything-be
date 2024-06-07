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

func NewRecycleHandlerImpl(usecase usecase.RecycleUsecase) *RecycleHandlerImpl {
	return &RecycleHandlerImpl{
		RecycleUsecase: usecase,
	}
}

func (handler *RecycleHandlerImpl) GetHomeRecycle(c echo.Context) error {
	data, err := handler.RecycleUsecase.GetHomeRecycle()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)

	return c.JSON(http.StatusOK, responseData)

}
