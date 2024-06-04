package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/dto"
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type ManageAchievementHandlerImpl struct {
	usecae usecase.ManageAchievementUsecase
}

func NewManageAchievementHandler(usecae usecase.ManageAchievementUsecase) *ManageAchievementHandlerImpl {
	return &ManageAchievementHandlerImpl{usecae: usecae}
}

func (handler ManageAchievementHandlerImpl) UploadBadgeHandler(c echo.Context) error {

	badge, errFile := c.FormFile("badge")

	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "Invalid request body, details: "+errFile.Error())
	}
	if badge.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(badge.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
	}
	src, errFileOpen := badge.Open()
	if errFileOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+errFileOpen.Error())
	}
	defer src.Close()
	imageUrl, err := helper.UploadToCloudinary(src, "achievement_badge")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, details: "+err.Error())
	}
	responseData := &dto.UploadBadgeResponse{
		BadgeUrl: imageUrl,
	}

	return helper.ResponseHandler(c, http.StatusCreated, "Success", responseData)
}

func (handler ManageAchievementHandlerImpl) CreateAchievementHandler(c echo.Context) error {

	request := &dto.CreateArchievementRequest{}
	if err := c.Bind(request); err != nil {
		return helper.ErrorHandler(c, 400, "Invalid request body, details: "+err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ErrorHandler(c, 400, "Invalid request body, details: "+err.Error())
	}

	archievement, err := handler.usecae.CreateArchievementUsecase(request)
	if err != nil {
		if errors.Is(err, pkg.ErrArchievementLevelAlreadyExist) {
			return helper.ErrorHandler(c, 400, err.Error())
		}
		return helper.ErrorHandler(c, 500, "internal server error, details: "+err.Error())
	}
	responseData := &dto.CreateArchievementResponse{
		Level:       archievement.Level,
		TargetPoint: archievement.TargetPoint,
		BadgeUrl:    archievement.BadgeUrl,
	}
	return helper.ResponseHandler(c, 200, "Success", responseData)
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
	return helper.ResponseHandler(c, 200, "Success", responseData.Data)
}
