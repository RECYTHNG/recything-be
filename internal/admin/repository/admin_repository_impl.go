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

func (repository *AdminRepositoryImpl) GetDataAdmin(id string) (*entity.Admin, error) {
	var admin entity.Admin
	if err := repository.DB.GetDB().Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) AddProfileImage(id string, imageUrl string) (*entity.Admin, error) {
	var admin entity.Admin
	if err := repository.DB.GetDB().Model(&admin).Where("id = ?", id).Update("image_url", imageUrl).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
