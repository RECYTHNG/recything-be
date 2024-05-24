package repository

import (
	"github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/database"
)

type AdminRepositoryImpl struct {
	DB database.Database
}

func NewAdminRepository(db database.Database) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{DB: db}
}

func (repository *AdminRepositoryImpl) CreateDataAdmin(admin *entity.Admin) (*entity.Admin, error) {
	if err := repository.DB.GetDB().Create(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) UpdateDataAdmin(admin *entity.Admin, id string) (*entity.Admin, error) {
	if err := repository.DB.GetDB().Save(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) FindAdminByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin
	if err := repository.DB.GetDB().Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindAdminByID(id string) (*entity.Admin, error) {
	var admin entity.Admin
	if err := repository.DB.GetDB().Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) GetDataAllAdmin(limit int) ([]entity.Admin, error) {
	var admins []entity.Admin
	if err := repository.DB.GetDB().Order("id desc").Limit(limit).Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (repository *AdminRepositoryImpl) FindLastIdAdmin() (string, error) {
	var admin entity.Admin
	if err := repository.DB.GetDB().Order("id desc").First(&admin).Error; err != nil {
		return "AD0000", err
	}
	return admin.ID, nil
}
