package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	"github.com/sawalreverr/recything/internal/task/manage_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ManageTaskHandlerImpl struct {
	Usecase usecase.ManageTaskUsecase
}

func NewManageTaskHandler(usecase usecase.ManageTaskUsecase) *ManageTaskHandlerImpl {
	return &ManageTaskHandlerImpl{Usecase: usecase}
}

func (handler *ManageTaskHandlerImpl) CreateTaskHandler(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	var request dto.CreateTaskResquest
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body, detail : "+err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	taskChallange, err := handler.Usecase.CreateTaskUsecase(&request, claims.UserID)

	if err != nil {
		if errors.Is(err, pkg.ErrTaskStepsNull) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTaskStepsNull.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	taskStep := []dto.TaskSteps{}

	data := dto.CreateTaskResponse{
		Id:          taskChallange.ID,
		Title:       taskChallange.Title,
		Description: taskChallange.Description,
		StartDate:   taskChallange.StartDate,
		EndDate:     taskChallange.EndDate,
		Steps:       taskStep,
	}
	for _, step := range taskChallange.TaskSteps {
		taskSteps := dto.TaskSteps{
			Title:       step.Title,
			Description: step.Description,
		}
		taskStep = append(taskStep, taskSteps)
	}
	data.Steps = taskStep

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}
