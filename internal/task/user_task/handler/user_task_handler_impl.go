package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	"github.com/sawalreverr/recything/internal/task/user_task/usecase"
)

type UserTaskHandlerImpl struct {
	Usecase usecase.UserTaskUsecase
}

func NewUserTaskHandler(usecase usecase.UserTaskUsecase) *UserTaskHandlerImpl {
	return &UserTaskHandlerImpl{Usecase: usecase}
}

func (handler *UserTaskHandlerImpl) GetAllTasksHandler(c echo.Context) error {
	userTask, err := handler.Usecase.GetAllTasksUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}

	var data []dto.DataUserTask

	for _, task := range userTask {
		var taskStep []dto.TaskSteps

		for _, step := range task.TaskSteps {
			taskStep = append(taskStep, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}
		data = append(data, dto.DataUserTask{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Thumbnail:   task.Thumbnail,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			Point:       task.Point,
			Status:      task.Status,
			TaskSteps:   taskStep,
		})
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}
