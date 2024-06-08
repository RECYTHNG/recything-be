package handler

import "github.com/labstack/echo/v4"

type RecycleHandler interface {
	GetHomeRecycleHandler(c echo.Context) error
	SearchVideoHandler(c echo.Context) error
	GetAllCategoryVideoHandler(c echo.Context) error
}
