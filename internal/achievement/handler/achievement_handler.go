package handler

import "github.com/labstack/echo/v4"

type AchievementHandler interface {
	AddAchievementHandler(echo.Context) error
	GetAchievementsHandler(echo.Context) error
	GetAchievementByIdHandler(echo.Context) error
	UpdateAchievementHandler(echo.Context) error
	DeleteAchievementHandler(echo.Context) error
}
