package usecase

import (
	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/entity"
)

type AdminUsecase interface {
	AddAdminUsecase(request dto.AdminRequestCreate) (*entity.Admin, error)
	UpdateAdminUsecase(request dto.AdminUpdateRequest) (*entity.Admin, error)
}
