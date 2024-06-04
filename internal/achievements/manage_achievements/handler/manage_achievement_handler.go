package handler

import "github.com/labstack/echo/v4"

type ManageAchievementHandler interface {
	UploadBadgeHandler(c echo.Context) error
	CreateAchievementHandler(c echo.Context) error
	GetAllAchievementHandler(c echo.Context) error
	GetAchievementByIdHandler(c echo.Context) error
	UpdateAchievementHandler(c echo.Context) error
	UpdateBadgeHandler(c echo.Context) error
}
