package handler

import "github.com/labstack/echo/v4"

type ManageTaskHandler interface {
	CreateTaskHandler(c echo.Context) error
	GetTaskChallangePagginationHandler(c echo.Context) error
	UploadThumbnailHandler(c echo.Context) error
}
