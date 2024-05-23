package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authHandler "github.com/sawalreverr/recything/internal/auth/handler"
	authUsecase "github.com/sawalreverr/recything/internal/auth/usecase"
	"github.com/sawalreverr/recything/internal/middleware"
	userRepo "github.com/sawalreverr/recything/internal/user/repository"
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

	// Example need user auth
	s.app.GET("/needUser", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}, middleware.UserMiddleware)
}

func (s *echoServer) authHttpHandler() {
	repository := userRepo.NewUserRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repository)
	handler := authHandler.NewAuthHandler(usecase)

	s.gr.POST("/register", handler.Register)
	s.gr.POST("/verify-otp", handler.VerifyOTP)
	s.gr.POST("/resend-otp", handler.ResendOTP)
	s.gr.POST("/login", handler.Login)
}
