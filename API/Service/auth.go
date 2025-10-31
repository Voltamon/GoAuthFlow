package auth

import (
	"fmt"

	data "github.com/Voltamon/GoAuthFlow/API/Data"
)

type AuthService struct {
	userRepo *data.UserRepo
}

func NewAuthService(userRepo *data.UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) RegisterUser(user *data.User) (*data.User, error) {
	existingUsers, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, existingUser := range existingUsers {
		if existingUser.Hash == user.Hash {
			return nil, fmt.Errorf("user already exists")
		}
	}

	err = s.userRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) LoginUser(user *data.User) (*data.User, error) {
	existingUser, err := s.userRepo.GetUserbyHash(user)
	if err != nil {
		return nil, err
	}

	existingUser.Hash = ""
	return existingUser, nil
}
