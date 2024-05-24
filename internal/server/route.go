package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authHandler "github.com/sawalreverr/recything/internal/auth/handler"
	authUsecase "github.com/sawalreverr/recything/internal/auth/usecase"
	"github.com/sawalreverr/recything/internal/middleware"
	userHandler "github.com/sawalreverr/recything/internal/user/handler"
	userRepo "github.com/sawalreverr/recything/internal/user/repository"
	userUsecase "github.com/sawalreverr/recything/internal/user/usecase"
)

var (
	SuperAdminMiddleware        = middleware.RoleBasedMiddleware("superadmin")
	SuperAdminOrAdminMiddleware = middleware.RoleBasedMiddleware("superadmin", "admin")
	UserMiddleware              = middleware.RoleBasedMiddleware("user")
	AllRoleMiddleware           = middleware.RoleBasedMiddleware("superadmin", "admin", "user")
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
	}, UserMiddleware)
}

func (s *echoServer) authHttpHandler() {
	repository := userRepo.NewUserRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repository)
	handler := authHandler.NewAuthHandler(usecase)

	// Register User
	s.gr.POST("/register", handler.Register)

	// Verify OTP after Register
	s.gr.POST("/verify-otp", handler.VerifyOTP)

	// Resend OTP
	s.gr.POST("/resend-otp", handler.ResendOTP)

	// Login User
	s.gr.POST("/login", handler.Login)
}

func (s *echoServer) userHttpHandler() {
	repository := userRepo.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repository)
	handler := userHandler.NewUserHandler(usecase)

	// Profile user based on JWT user token
	s.gr.GET("/user/profile", handler.Profile, middleware.UserMiddleware)

	// Edit detail user based on JWT user token
	s.gr.PUT("/user/profile", handler.UpdateDetail, middleware.UserMiddleware)

	// Upload avatar user based on JWT user token
	s.gr.POST("/user/uploadAvatar", handler.UploadAvatar, middleware.UserMiddleware)

	// Find user data using param userId, doesnt need jwt
	s.gr.GET("/user/:userId", handler.FindUser)

	// Find all user data with pagination, need JWT admin or superadmin token
	s.gr.GET("/users", handler.FindAllUser, middleware.AdminMiddleware, middleware.SuperAdminMiddleware)

	// Delete user data using param userId
	s.gr.DELETE("/user/:userId", handler.DeleteUser, middleware.AdminMiddleware, middleware.SuperAdminMiddleware)
}
