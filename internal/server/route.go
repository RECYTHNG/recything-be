package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	achievementHandler "github.com/sawalreverr/recything/internal/achievements/manage_achievements/handler"
	achievementRepo "github.com/sawalreverr/recything/internal/achievements/manage_achievements/repository"
	achievementUsecase "github.com/sawalreverr/recything/internal/achievements/manage_achievements/usecase"
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
	approvalTaskHandler "github.com/sawalreverr/recything/internal/task/approval_task/handler"
	approvalTaskRepo "github.com/sawalreverr/recything/internal/task/approval_task/repository"
	approvalTaskUsecase "github.com/sawalreverr/recything/internal/task/approval_task/usecase"
	taskHandler "github.com/sawalreverr/recything/internal/task/manage_task/handler"
	taskRepo "github.com/sawalreverr/recything/internal/task/manage_task/repository"
	taskUsecase "github.com/sawalreverr/recything/internal/task/manage_task/usecase"
	userTaskHandler "github.com/sawalreverr/recything/internal/task/user_task/handler"
	userTaskRepo "github.com/sawalreverr/recything/internal/task/user_task/repository"
	userTaskUsecase "github.com/sawalreverr/recything/internal/task/user_task/usecase"
	userHandler "github.com/sawalreverr/recything/internal/user/handler"
	userRepo "github.com/sawalreverr/recything/internal/user/repository"
	userUsecase "github.com/sawalreverr/recything/internal/user/usecase"
	videoHandler "github.com/sawalreverr/recything/internal/video/manage_video/handler"
	videoRepo "github.com/sawalreverr/recything/internal/video/manage_video/repository"
	videoUsecase "github.com/sawalreverr/recything/internal/video/manage_video/usecase"
	userVideoHandler "github.com/sawalreverr/recything/internal/video/user_video/handler"
	userVideoRepo "github.com/sawalreverr/recything/internal/video/user_video/repository"
	userVideoUsecase "github.com/sawalreverr/recything/internal/video/user_video/usecase"
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

func (s *echoServer) manageTask() {
	repository := taskRepo.NewManageTaskRepository(s.db)
	usecase := taskUsecase.NewManageTaskUsecase(repository)
	handler := taskHandler.NewManageTaskHandler(usecase)

	// upload thumbnail task
	s.gr.POST("/tasks/thumbnail", handler.UploadThumbnailHandler, SuperAdminOrAdminMiddleware)

	// create task by admin or super admin
	s.gr.POST("/tasks", handler.CreateTaskHandler, SuperAdminOrAdminMiddleware)

	// get task challenge by pagination
	s.gr.GET("/tasks", handler.GetTaskChallengePaginationHandler, SuperAdminOrAdminMiddleware)

	// get task challenge by id
	s.gr.GET("/tasks/:taskId", handler.GetTaskByIdHandler, SuperAdminOrAdminMiddleware)

	// update task challenge
	s.gr.PUT("/tasks/:taskId", handler.UpdateTaskHandler, SuperAdminOrAdminMiddleware)

	// delete task challenge
	s.gr.DELETE("/tasks/:taskId", handler.DeleteTaskHandler, SuperAdminOrAdminMiddleware)

}

func (s *echoServer) userTask() {
	repository := userTaskRepo.NewUserTaskRepository(s.db)
	usecase := userTaskUsecase.NewUserTaskUsecase(repository)
	handler := userTaskHandler.NewUserTaskHandler(usecase)

	// get all tasks
	s.gr.GET("/user/tasks", handler.GetAllTasksHandler, UserMiddleware)

	// get task by id
	s.gr.GET("/user/tasks/:taskId", handler.GetTaskByIdHandler, UserMiddleware)

	// create task by user or start task
	s.gr.POST("/user/tasks/:taskChallengeId", handler.CreateUserTaskHandler, UserMiddleware)

	// get task in progress by user current
	s.gr.GET("/user_current/tasks/in-progress", handler.GetUserTaskByUserIdHandler, UserMiddleware)

	// send task done by user current
	s.gr.POST("/user_current/tasks/:userTaskId", handler.UploadImageTaskHandler, UserMiddleware)

	// get task done by user current
	s.gr.GET("/user_current/tasks/done", handler.GetUserTaskDoneByUserIdHandler, UserMiddleware)

	// update user task if reject
	s.gr.PUT("/user_current/tasks/:userTaskId", handler.UpdateUserTaskHandler, UserMiddleware)

	// get user task details if repair
	s.gr.GET("/user_current/tasks/:userTaskId", handler.GetUserTaskDetailsHandler, UserMiddleware)
}

func (s *echoServer) approvalTask() {
	repository := approvalTaskRepo.NewApprovalTaskRepositoryImpl(s.db)
	usecase := approvalTaskUsecase.NewApprovalTaskUsecase(repository)
	handler := approvalTaskHandler.NewApprovalTaskHandler(usecase)

	// get all pagination user task
	s.gr.GET("/approve_tasks", handler.GetAllApprovalTaskPaginationHandler, SuperAdminOrAdminMiddleware)

	// approve user task
	s.gr.PUT("/approve_tasks/:userTaskId", handler.ApproveUserTaskHandler, SuperAdminOrAdminMiddleware)

	// reject user task
	s.gr.PUT("/reject_tasks/:userTaskId", handler.RejectUserTaskHandler, SuperAdminOrAdminMiddleware)

	// get user task details
	s.gr.GET("/user_task/:userTaskId", handler.GetUserTaskDetailsHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) manageAchievement() {
	repository := achievementRepo.NewManageAchievementRepository(s.db)
	usecase := achievementUsecase.NewManageAchievementUsecase(repository)
	handler := achievementHandler.NewManageAchievementHandler(usecase)

	// upload badge achievement
	s.gr.POST("/achievements/badge", handler.UploadBadgeHandler, SuperAdminOrAdminMiddleware)

	// create achievement
	s.gr.POST("/achievements", handler.CreateAchievementHandler, SuperAdminOrAdminMiddleware)

	// get all achievement
	s.gr.GET("/achievements", handler.GetAllAchievementHandler, SuperAdminOrAdminMiddleware)

	// get achievement by id
	s.gr.GET("/achievements/:achievementId", handler.GetAchievementByIdHandler, SuperAdminOrAdminMiddleware)

	// update badge achievement
	s.gr.PUT("/achievements/badge", handler.UpdateBadgeHandler, SuperAdminOrAdminMiddleware)

	// update achievement
	s.gr.PUT("/achievements/:achievementId", handler.UpdateAchievementHandler, SuperAdminOrAdminMiddleware)

	// delete achievement
	s.gr.DELETE("/achievements/:achievementId", handler.DeleteAchievementHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) manageVideo() {
	repository := videoRepo.NewManageVideoRepository(s.db)
	usecase := videoUsecase.NewManageVideoUsecaseImpl(repository)
	handler := videoHandler.NewManageVideoHandlerImpl(usecase)

	// upload thumbnail video
	s.gr.POST("/videos/thumbnail", handler.UploadThumbnailVideoHandler, SuperAdminOrAdminMiddleware)

	// create data video
	s.gr.POST("/videos/data", handler.CreateDataVideoHandler, SuperAdminOrAdminMiddleware)

	// create category video
	s.gr.POST("/videos/categories", handler.CreateCategoryVideoHandler, SuperAdminOrAdminMiddleware)

	// get all category video
	s.gr.GET("/videos/categories", handler.GetAllCategoryVideoHandler, SuperAdminOrAdminMiddleware)

	// get all data video pagination
	s.gr.GET("/videos/data", handler.GetAllDataVideoPaginationHandler, SuperAdminOrAdminMiddleware)

	// get details data video by id
	s.gr.GET("/videos/data/:videoId", handler.GetDetailsDataVideoByIdHandler, SuperAdminOrAdminMiddleware)

	// update data video
	s.gr.PUT("/videos/data/:videoId", handler.UpdateDataVideoHandler, SuperAdminOrAdminMiddleware)

	// delete data video
	s.gr.DELETE("/videos/data/:videoId", handler.DeleteDataVideoHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) userVideo() {
	repository := userVideoRepo.NewUserVideoRepository(s.db)
	usecase := userVideoUsecase.NewUserVideoUsecase(repository)
	handler := userVideoHandler.NewUserVideoHandler(usecase)

	// get all video
	s.gr.GET("/videos", handler.GetAllVideoHandler, UserMiddleware)
}
