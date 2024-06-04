package handler

import "github.com/labstack/echo/v4"

type ManageAchievementHandler interface {
	UploadBadgeHandler(c echo.Context) error
	CreateAchievementHandler(c echo.Context) error
}
