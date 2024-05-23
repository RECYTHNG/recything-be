package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/handler"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/admin/usecase"
)

func (s *echoServer) publicHttpHandler() {
	// Healthy Check
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Swagger
	s.app.Static("/assets", "web/assets")
	s.app.Static("/docs", "docs")
	s.app.GET("/", func(c echo.Context) error {
		return c.File("web/index.html")
	})

	// Admin
	adminRepo := repository.NewAdminRepository(s.db)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	s.app.POST("/admin/register", adminHandler.AddAdminHandler)
}
