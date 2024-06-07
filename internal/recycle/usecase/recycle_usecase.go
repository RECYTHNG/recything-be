package usecase

import (
	"github.com/sawalreverr/recything/internal/recycle/dto"
)

type RecycleUsecase interface {
	GetHomeRecycleUsecase() (*dto.RecycleHomeResponse, error)
	SearchVideoUsecase(title string, category string) (*dto.SearchVideoResponse, error)
}
