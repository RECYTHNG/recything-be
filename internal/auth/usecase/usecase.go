package auth

import (
	a "github.com/sawalreverr/recything/internal/auth"
	"github.com/sawalreverr/recything/internal/helper"
	u "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type authUsecase struct {
	userRepository u.UserRepository
}

func NewAuthUsecase(userRepo u.UserRepository) a.AuthUsecase {
	return &authUsecase{userRepository: userRepo}
}

func (uc *authUsecase) RegisterUser(user a.Register) (*u.User, error) {
	emailFound, _ := uc.userRepository.FindByEmail(user.Email)
	if emailFound != nil {
		return nil, pkg.ErrEmailAlreadyExists
	}

	phoneFound, _ := uc.userRepository.FindByPhoneNumber(user.PhoneNumber)
	if phoneFound != nil {
		return nil, pkg.ErrPhoneNumberAlreadyExists
	}

	lastID, _ := uc.userRepository.FindLastID()
	newID := helper.GenerateCustomID(lastID, "USR")

	hashedPass, _ := helper.GenerateHash(user.Password)

	newUser := u.User{
		ID:          newID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    hashedPass,
		OTP:         helper.GenerateOTP(),
	}

	createdUser, err := uc.userRepository.Create(newUser)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	return createdUser, nil
}

func (uc *authUsecase) LoginUser(user a.Login) (string, error) {
	userFound, err := uc.userRepository.FindByEmail(user.Email)
	if err != nil {
		return "", pkg.ErrUserNotFound
	}

	ok := helper.ComparePassword(userFound.Password, user.Password)
	if !ok {
		return "", pkg.ErrPasswordInvalid
	}

	token, err := helper.GenerateTokenJWT(userFound.ID, "user")

	return token, err
}

func (uc *authUsecase) VerifyOTP(user a.OTPRequest) error {
	// userFound, err := uc.userRepository.FindByEmail(user.Email)
	// if err != nil {
	// 	return pkg.ErrUserNotFound
	// }

	return nil
}

func (uc *authUsecase) UpdateOTP(phoneNumber string) error {

	return nil
}
