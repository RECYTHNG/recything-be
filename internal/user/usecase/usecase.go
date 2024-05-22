package user

import (
	u "github.com/sawalreverr/recything/internal/user"
)

type userUsecase struct {
	userRepository u.UserRepository
}

func NewUserUsecase(userRepo u.UserRepository) u.UserUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) UpdateUserDetail(userID string, user u.UserDetail) error {
	// userFound, _ := u.userRepository.()

	return nil
}

func (u *userUsecase) UpdateUserPicture(userID string, picture_url string) error {

	return nil
}

func (u *userUsecase) FindUserByID(userID string) (*u.User, error) {

	return nil, nil
}

func (u *userUsecase) FindAllUser(page int, limit int, sortBy string, sortType string) (*[]u.User, error) {

	return nil, nil
}

func (u *userUsecase) DeleteUser(userID string) error {

	return nil
}
