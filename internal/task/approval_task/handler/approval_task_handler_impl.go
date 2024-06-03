package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/approval_task/dto"
	"github.com/sawalreverr/recything/internal/task/approval_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ApprovalTaskHandlerImpl struct {
	usecase usecase.ApprovalTaskUsecase
}

func NewApprovalTaskHandler(usecase usecase.ApprovalTaskUsecase) *ApprovalTaskHandlerImpl {
	return &ApprovalTaskHandlerImpl{usecase: usecase}
}

func (handler *ApprovalTaskHandlerImpl) GetAllApprovalTaskPaginationHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	userTask, total, err := handler.usecase.GetAllApprovalTaskPagination(limitInt, (pageInt-1)*limitInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	var data []dto.DataUserTask
	for _, task := range userTask {
		data = append(data, dto.DataUserTask{
			Id:           task.ID,
			StatusAccept: task.StatusAccept,
			Point:        task.Point,
			TaskChallenge: dto.DataTasks{
				Id:        task.TaskChallenge.ID,
				Title:     task.TaskChallenge.Title,
				StartDate: task.TaskChallenge.StartDate,
				EndDate:   task.TaskChallenge.EndDate,
			},
			User: dto.DataUser{
				Id:   task.User.ID,
				Name: task.User.Name,
			},
		})
	}

	totalPage := total / limitInt
	if total%limitInt != 0 {
		totalPage++
	}
	responseDataPagination := dto.GetUserTaskPagination{
		Code:      http.StatusOK,
		Message:   "success get all task",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: total,
		TotalPage: totalPage,
	}
	responseData := helper.ResponseData(http.StatusOK, "success get all task", responseDataPagination)

	return c.JSON(http.StatusOK, responseData)

}

func (handler *ApprovalTaskHandlerImpl) ApproveUserTaskHandler(c echo.Context) error {
	userTaskId := c.Param("userTaskId")
	if err := handler.usecase.ApproveUserTask(userTaskId); err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "success approve user task", nil)

	return c.JSON(http.StatusOK, responseData)
}
