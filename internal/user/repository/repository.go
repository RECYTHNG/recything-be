package user

import (
	"errors"
	"fmt"

	"github.com/sawalreverr/recything/internal/database"
	u "github.com/sawalreverr/recything/internal/user"
	"gorm.io/gorm"
)

type userRepository struct {
	DB database.Database
}

func NewUserRepository(db database.Database) u.UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user u.User) (*u.User, error) {
	if err := r.DB.GetDB().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*u.User, error) {
	var user u.User
	if err := r.DB.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*u.User, error) {
	var user u.User
	if err := r.DB.GetDB().Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(userID string) (*u.User, error) {
	var user u.User
	if err := r.DB.GetDB().Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll(page int, limit int, sortBy string, sortType string) (*[]u.User, error) {
	var users []u.User

	db := r.DB.GetDB()
	offset := (page - 1) * limit

	if sortBy != "" {
		sort := fmt.Sprintf("%s %s", sortBy, sortType)
		db = db.Order(sort)
	}

	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) FindLastID() (string, error) {
	var user u.User
	if err := r.DB.GetDB().Where("deleted_at IS NULL").Order("id DESC").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "USR0000", nil
		}
		return "", err
	}

	return user.ID, nil
}

func (r *userRepository) Update(user u.User) error {
	if err := r.DB.GetDB().Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(userID string) error {
	var user u.User
	if err := r.DB.GetDB().Where("id = ?", userID).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CountAllUser() (int, error) {
	var totalCount int64

	if err := r.DB.GetDB().Model(&u.User{}).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}
