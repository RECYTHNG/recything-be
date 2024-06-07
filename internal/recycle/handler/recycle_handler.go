package handler

import "github.com/labstack/echo/v4"

type RecycleHandler interface {
	GetHomeRecycle(c echo.Context) error
}
