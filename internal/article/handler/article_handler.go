package handler

import "github.com/labstack/echo/v4"

type ArticleHandler interface {
	CreateArticleHandler(echo.Context) error
	GetDataAllArticleHandler(echo.Context) error
	GetDataArticleByIdHandler(echo.Context) error
	UpdateArticleHandler(echo.Context) error
	DeleteArticleHandler(echo.Context) error
}
