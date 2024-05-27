package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/handler"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/admin/usecase"
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
	userRepository := userRepo.NewUserRepository(s.db)
	adminRepository := repository.NewAdminRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(userRepository, adminRepository)
	handler := authHandler.NewAuthHandler(usecase)

	// Register User
	s.gr.POST("/register", handler.Register)

	// Verify OTP after Register
	s.gr.POST("/verify-otp", handler.VerifyOTP)

	// Resend OTP
	s.gr.POST("/resend-otp", handler.ResendOTP)

	// Login User
	s.gr.POST("/login", handler.LoginUser)

	// Login Admin
	s.gr.POST("/admin/login", handler.LoginAdmin)
}

func (s *echoServer) userHttpHandler() {
	repository := userRepo.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repository)
	handler := userHandler.NewUserHandler(usecase)

	// Profile user based on JWT user token
	s.gr.GET("/user/profile", handler.Profile, UserMiddleware)

	// Edit detail user based on JWT user token
	s.gr.PUT("/user/profile", handler.UpdateDetail, UserMiddleware)

	// Upload avatar user based on JWT user token
	s.gr.POST("/user/uploadAvatar", handler.UploadAvatar, UserMiddleware)

	// Find user data using param userId, doesnt need jwt
	s.gr.GET("/user/:userId", handler.FindUser, AllRoleMiddleware)

	// Find all user data with pagination, need JWT admin or superadmin token
	s.gr.GET("/users", handler.FindAllUser, SuperAdminOrAdminMiddleware)

	// Delete user data using param userId
	s.gr.DELETE("/user/:userId", handler.DeleteUser, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) supAdminHttpHandler() {
	repository := repository.NewAdminRepository(s.db)
	usecase := usecase.NewAdminUsecase(repository)
	handler := handler.NewAdminHandler(usecase)

	// register admin by super admin
	s.gr.POST("/admins", handler.AddAdminHandler, SuperAdminMiddleware)

	// get all admin by super admin
	s.gr.GET("/admins", handler.GetDataAllAdminHandler, SuperAdminMiddleware)

	// get data admin by id by super admin
	s.gr.GET("/admins/:adminId", handler.GetDataAdminByIdHandler, SuperAdminMiddleware)

	// update admin by super admin
	s.gr.PUT("/admins/:adminId", handler.UpdateAdminHandler, SuperAdminMiddleware)

	// delete admin by super admin
	s.gr.DELETE("/admins/:adminId", handler.DeleteAdminHandler, SuperAdminMiddleware)

	// get profile admin or super admin
	s.gr.GET("/profile", handler.GetProfileAdminHandler, SuperAdminOrAdminMiddleware)
}
