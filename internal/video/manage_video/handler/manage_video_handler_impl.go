package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/video/manage_video/usecase"
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
