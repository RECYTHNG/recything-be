package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/handler"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	authHandler "github.com/sawalreverr/recything/internal/auth/handler"
	authUsecase "github.com/sawalreverr/recything/internal/auth/usecase"
	faqHandler "github.com/sawalreverr/recything/internal/faq/handler"
	faqRepo "github.com/sawalreverr/recything/internal/faq/repository"
	faqUsecase "github.com/sawalreverr/recything/internal/faq/usecase"
	"github.com/sawalreverr/recything/internal/middleware"
	reportHandler "github.com/sawalreverr/recything/internal/report/handler"
	reportRepo "github.com/sawalreverr/recything/internal/report/repository"
	reportUsecase "github.com/sawalreverr/recything/internal/report/usecase"
	userHandler "github.com/sawalreverr/recything/internal/user/handler"
	userRepo "github.com/sawalreverr/recything/internal/user/repository"
	userUsecase "github.com/sawalreverr/recything/internal/user/usecase"
)

var (
	SuperAdminMiddleware        = middleware.RoleBasedMiddleware("super admin")
	SuperAdminOrAdminMiddleware = middleware.RoleBasedMiddleware("super admin", "admin")
	UserMiddleware              = middleware.RoleBasedMiddleware("user")
	AllRoleMiddleware           = middleware.RoleBasedMiddleware("super admin", "admin", "user")
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
	s.gr.POST("/admin", handler.AddAdminHandler, SuperAdminMiddleware)

	// get all admin by super admin
	s.gr.GET("/admins", handler.GetDataAllAdminHandler, SuperAdminMiddleware)

	// get data admin by id by super admin
	s.gr.GET("/admin/:adminId", handler.GetDataAdminByIdHandler, SuperAdminMiddleware)

	// update admin by super admin
	s.gr.PUT("/admin/:adminId", handler.UpdateAdminHandler, SuperAdminMiddleware)

	// delete admin by super admin
	s.gr.DELETE("/admin/:adminId", handler.DeleteAdminHandler, SuperAdminMiddleware)

	// get profile admin or super admin
	s.gr.GET("/admin/profile", handler.GetProfileAdminHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) reportHttpHandler() {
	repository := reportRepo.NewReportRepository(s.db)
	usecase := reportUsecase.NewReportUsecase(repository)
	handler := reportHandler.NewReportHandler(usecase)

	// User create new report
	s.gr.POST("/report", handler.NewReport, UserMiddleware)

	// User get all history reports
	s.gr.GET("/report", handler.GetHistoryUserReports, UserMiddleware)

	// Admin update status approved or reject
	s.gr.PUT("/report/:reportId", handler.UpdateStatus, SuperAdminOrAdminMiddleware)

	// Admin get all with pagination and filter
	s.gr.GET("/reports", handler.GetAllReports, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) faqHttpHandler() {
	repository := faqRepo.NewFaqRepository(s.db)
	usecase := faqUsecase.NewFaqUsecase(repository)
	handler := faqHandler.NewFaqHandler(usecase)

	// User get all faqs
	s.gr.GET("/faqs", handler.GetAllFaqs, UserMiddleware)

	// User get all faqs by category
	s.gr.GET("/faqs/category", handler.GetFaqsByCategory, UserMiddleware)

	// User get all faqs by keyword
	s.gr.GET("/faqs/search", handler.GetFaqsByKeyword, UserMiddleware)
}
