package user

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// struct
type User struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"-"`
	Point       uint      `json:"point"`
	Gender      string    `json:"gender" gorm:"type:enum('laki-laki', 'perempuan', '-');default:-"`
	BirthDate   time.Time `json:"birth_date" gorm:"type:datetime"`
	Address     string    `json:"address"`
	PictureURL  string    `json:"picture_url"`
	OTP         uint      `json:"otp"`
	IsVerified  bool      `json:"is_verified"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// interface
type UserRepository interface {
	Create(user User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByPhoneNumber(phoneNumber string) (*User, error)
	FindByID(userID string) (*User, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*[]User, error)
	FindLastID() (string, error)
	Update(user User) error
	Delete(userID string) error
}

type UserUsecase interface {
	UpdateUserDetail(userID string, user UserDetail) error
	UpdateUserPicture(userID string, picture_url string) error

	FindUserByID(userID string) (*User, error)
	FindAllUser(page int, limit int, sortBy string, sortType string) (*[]User, error)
	DeleteUser(userID string) error
}

type UserHandler interface {
	// User
	Profile(c echo.Context) error
	UpdateDetail(c echo.Context) error
	UploadAvatar(c echo.Context) error

	// Manage User
	FindAllUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
