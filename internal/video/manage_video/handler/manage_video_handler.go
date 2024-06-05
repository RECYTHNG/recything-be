package handler

import "github.com/labstack/echo/v4"

type ManageVideoHandler interface {
	CreateDataVideoHandler(c *echo.Context) error
	CreateCategoryVideoHandler(c *echo.Context) error
	GetAllCategoryVideoHandler(c *echo.Context) error
	UploadThumbnailVideoHandler(c *echo.Context) error
}
