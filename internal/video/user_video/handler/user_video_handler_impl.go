package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
	"github.com/sawalreverr/recything/internal/video/user_video/usecase"
)

type UserVideoHandlerImpl struct {
	Usecase usecase.UserVideoUsecase
}

func NewUserVideoHandler(usecase usecase.UserVideoUsecase) *UserVideoHandlerImpl {
	return &UserVideoHandlerImpl{Usecase: usecase}
}

func (handler *UserVideoHandlerImpl) GetAllVideoHandler(c echo.Context) error {
	videos, err := handler.Usecase.GetAllVideo()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	var dataVideo []dto.DataVideo
	data := dto.GetAllVideoResponse{
		DataVideo: dataVideo,
	}

	for _, video := range *videos {
		dataVideo = append(dataVideo, dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
			LinkVideo:    video.Link,
			Viewer:       video.Viewer,
		})
	}
	data.DataVideo = dataVideo
	responseData := helper.ResponseData(http.StatusOK, "success get all video", data.DataVideo)

	return c.JSON(http.StatusOK, responseData)
}
