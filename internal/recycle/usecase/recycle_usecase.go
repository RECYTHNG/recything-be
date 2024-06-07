package usecase

import (
	"github.com/sawalreverr/recything/internal/recycle/dto"
)

type RecycleUsecase interface {
	GetHomeRecycle() (*dto.RecycleHomeResponse, error)
}
