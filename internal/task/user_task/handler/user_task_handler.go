package handler

import "github.com/labstack/echo/v4"

type UserTaskHandler interface {
	GetAllTasksHandler(c echo.Context) error
	GetTaskByIdHandler(c echo.Context) error
	CreateUserTaskHandler(c echo.Context) error
}
