package handler

import "github.com/labstack/echo/v4"

type ManageVideoHandler interface {
	CreateDataVideoHandler(c *echo.Context) error
}
