package handler

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type ManageAchievementHandlerImpl struct {
	usecae usecase.ManageAchievementUsecase
}

func NewManageAchievementHandler(usecae usecase.ManageAchievementUsecase) ManageAchievementHandler {
	return &ManageAchievementHandlerImpl{usecae: usecae}
}

func (handler ManageAchievementHandlerImpl) GetAllAchievementHandler(c echo.Context) error {
	achievements, err := handler.usecae.GetAllArchievementUsecase()
	if err != nil {
		return helper.ErrorHandler(c, 500, "internal server error, details: "+err.Error())
	}
	var data []*dto.DataAchievement
	for _, achievement := range achievements {
		data = append(data, &dto.DataAchievement{
			Id:          achievement.ID,
			Level:       achievement.Level,
			TargetPoint: achievement.TargetPoint,
			BadgeUrl:    achievement.BadgeUrl,
		})
	}
	responseData := &dto.GetAllAchievementResponse{
		Data: data,
	}
	return helper.ResponseHandler(c, http.StatusOK, "Success", responseData.Data)
}

func (handler ManageAchievementHandlerImpl) GetAchievementByIdHandler(c echo.Context) error {
	achievementId := c.Param("achievementId")
	achievementIdInt, errConvert := strconv.Atoi(achievementId)

	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "Invalid request param, details: "+errConvert.Error())
	}

	achievement, err := handler.usecae.GetAchievementByIdUsecase(achievementIdInt)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrAchievementNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	responseData := &dto.DataAchievement{
		Id:          achievement.ID,
		Level:       achievement.Level,
		TargetPoint: achievement.TargetPoint,
		BadgeUrl:    achievement.BadgeUrl,
	}
	return helper.ResponseHandler(c, http.StatusOK, "Success", responseData)
}

func (handler ManageAchievementHandlerImpl) UpdateAchievementHandler(c echo.Context) error {
	achievementId := c.Param("achievementId")
	achievementIdInt, errConvert := strconv.Atoi(achievementId)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "Invalid request param, details: "+errConvert.Error())
	}

	request := dto.UpdateAchievementRequest{}
	json_data := c.FormValue("json_data")
	if err := json.Unmarshal([]byte(json_data), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	var badge []*multipart.FileHeader
	if form != nil {
		badge = form.File["badge"]
	}
	err := handler.usecae.UpdateAchievementUsecase(&request, badge, achievementIdInt)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrAchievementNotFound.Error())
		}
		if errors.Is(err, pkg.ErrBadgeMaximum) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrBadgeMaximum.Error())
		}
		if errors.Is(err, errors.New("upload image size must less than 2MB")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB")
		}
		if errors.Is(err, errors.New("only image allowed")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed")
		}
		if errors.Is(err, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	return helper.ResponseHandler(c, http.StatusOK, "achievement updated", nil)
}

func (handler ManageAchievementHandlerImpl) DeleteAchievementHandler(c echo.Context) error {
	achievementId := c.Param("achievementId")
	achievementIdInt, errConvert := strconv.Atoi(achievementId)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "Invalid request param, details: "+errConvert.Error())
	}
	err := handler.usecae.DeleteAchievementUsecase(achievementIdInt)
	if err != nil {
		if errors.Is(err, pkg.ErrAchievementNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrAchievementNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	return helper.ResponseHandler(c, http.StatusOK, "achievement deleted", nil)
}
