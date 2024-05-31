package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type AchievementHandlerImpl struct {
	Usecase usecase.AchievementUseCase
}

func NewAchievementHandler(usecase usecase.AchievementUseCase) *AchievementHandlerImpl {
	return &AchievementHandlerImpl{Usecase: usecase}
}

func (h *AchievementHandlerImpl) AddAchievementHandler(c echo.Context) error {
	var request dto.AchievementRequestCreate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	achievement, err := h.Usecase.AddAchievement(request)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	response := dto.AchievementResponse{
		Id:          achievement.Id,
		Level:       achievement.Level,
		Lencana:     achievement.Lencana,
		TargetPoint: achievement.TargetPoint,
		CreatedAt:   achievement.CreatedAt,
		UpdatedAt:   achievement.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, helper.ResponseData(http.StatusCreated, "success", response))
}

func (h *AchievementHandlerImpl) GetAchievementsHandler(c echo.Context) error {
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
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}

	achievements, totalData, err := h.Usecase.GetAchievements(limitInt, pageInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	data := []dto.AchievementResponse{}
	for _, a := range achievements {
		data = append(data, dto.AchievementResponse{
			Id:          a.Id,
			Level:       a.Level,
			Lencana:     a.Lencana,
			TargetPoint: a.TargetPoint,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		})
	}

	totalPage := (totalData + limitInt - 1) / limitInt

	response := dto.AchievementResponseGetAll{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AchievementHandlerImpl) GetAchievementByIdHandler(c echo.Context) error {
	id := c.Param("achievementId")

	achievement, err := h.Usecase.GetAchievementById(id)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	response := dto.AchievementResponse{
		Id:          achievement.Id,
		Level:       achievement.Level,
		Lencana:     achievement.Lencana,
		TargetPoint: achievement.TargetPoint,
		CreatedAt:   achievement.CreatedAt,
		UpdatedAt:   achievement.UpdatedAt,
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", response))
}

func (h *AchievementHandlerImpl) UpdateAchievementHandler(c echo.Context) error {
	id := c.Param("achievementId")

	var request dto.AchievementRequestUpdate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	achievement, err := h.Usecase.UpdateAchievement(request, id)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	response := dto.AchievementResponse{
		Id:          achievement.Id,
		Level:       achievement.Level,
		Lencana:     achievement.Lencana,
		TargetPoint: achievement.TargetPoint,
		CreatedAt:   achievement.CreatedAt,
		UpdatedAt:   achievement.UpdatedAt,
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", response))
}

func (h *AchievementHandlerImpl) DeleteAchievementHandler(c echo.Context) error {
	id := c.Param("achievementId")

	err := h.Usecase.DeleteAchievement(id)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "achievement successfully deleted", nil))
}
