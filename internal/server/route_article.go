package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/feature/article/handler"
	"github.com/sawalreverr/recything/internal/feature/article/repository"
	"github.com/sawalreverr/recything/internal/feature/article/service"
	"gorm.io/gorm"
)

func RouteArticle(e *echo.Group, db *gorm.DB) {

	// manage article
	trashRepo := trashRepository.NewTrashCategoryRepository(db)
	articleRepo := repository.NewArticleRepository(db, trashRepo)
	articleServ := service.NewArticleService(articleRepo)
	articleHand := handler.NewArticleHandler(articleServ)

	admin := e.Group("/admins/manage/articles", jwt.JWTMiddleware())
	admin.POST("", articleHand.CreateArticle)
	admin.GET("", articleHand.GetAllArticle)
	admin.GET("/:id", articleHand.GetSpecificArticle)
	admin.PUT("/:id", articleHand.UpdateArticle)
	admin.DELETE("/:id", articleHand.DeleteArticle)

	user := e.Group("/articles", jwt.JWTMiddleware())
	user.GET("", articleHand.GetAllArticleUser)
	user.GET("/:id", articleHand.GetSpecificArticle)
	user.GET("/popular", articleHand.GetPopularArticle)
}
