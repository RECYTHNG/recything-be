package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type adminHandlerImpl struct {
	Usecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) *adminHandlerImpl {
	return &adminHandlerImpl{Usecase: adminUsecase}
}

func (handler *adminHandlerImpl) AddAdminHandler(c echo.Context) error {
	var request dto.AdminRequestCreate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	admin, errUc := handler.Usecase.AddAdminUsecase(request)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrEmailAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrEmailAlreadyExist.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	data := dto.AdminResponseRegister{
		Id:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusCreated, responseData)

}

func (handler *adminHandlerImpl) UpdateAdminHandler(c echo.Context) error {
	return nil
}
