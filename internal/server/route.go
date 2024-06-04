package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/handler"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	achievementHandler "github.com/sawalreverr/recything/internal/archievements/manage_archievements/handler"
	achievementRepo "github.com/sawalreverr/recything/internal/archievements/manage_archievements/repository"
	achievementUsecase "github.com/sawalreverr/recything/internal/archievements/manage_archievements/usecase"
	authHandler "github.com/sawalreverr/recything/internal/auth/handler"
	authUsecase "github.com/sawalreverr/recything/internal/auth/usecase"
	"github.com/sawalreverr/recything/internal/middleware"
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
}

func (s *echoServer) approvalTask() {
	repository := approvalTaskRepo.NewApprovalTaskRepository(s.db)
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
}
