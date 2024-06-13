package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
	"github.com/sawalreverr/recything/internal/video/user_video/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type UserVideoHandlerImpl struct {
	Usecase usecase.UserVideoUsecase
}

func NewUserVideoHandler(usecase usecase.UserVideoUsecase) UserVideoHandler {
	return &UserVideoHandlerImpl{Usecase: usecase}
}

func (handler *UserVideoHandlerImpl) GetAllVideoHandler(c echo.Context) error {
	videos, err := handler.Usecase.GetAllVideoUsecase()
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

func (handler *UserVideoHandlerImpl) SearchVideoByTitleHandler(c echo.Context) error {
	title := c.QueryParam("title")
	videos, err := handler.Usecase.SearchVideoByTitleUsecase(title)
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

func (handler *UserVideoHandlerImpl) GetVideoDetailHandler(c echo.Context) error {
	id := c.Param("videoId")
	intId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}

	video, comments, err := handler.Usecase.GetVideoDetailUsecase(intId)
	if err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	var dataComments []dto.DataComment
	data := dto.GetDetailsDataVideoByIdResponse{
		DataVideo: &dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
			LinkVideo:    video.Link,
			Viewer:       video.Viewer,
		},
		Comments: &dataComments,
	}

	for _, comment := range *comments {
		dataComments = append(dataComments, dto.DataComment{
			Id:        comment.ID,
			Comment:   comment.Comment,
			UserID:    comment.UserID,
			UserName:  comment.User.Name,
			CreatedAt: comment.CreatedAt,
		})
	}
	data.Comments = &dataComments
	responseData := helper.ResponseData(http.StatusOK, "success", data)

	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserVideoHandlerImpl) AddCommentHandler(c echo.Context) error {
	var request dto.AddCommentRequest
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := handler.Usecase.AddCommentUsecase(&request, userId); err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	responseData := helper.ResponseData(http.StatusOK, "success add comment", nil)
	return c.JSON(http.StatusCreated, responseData)
}
