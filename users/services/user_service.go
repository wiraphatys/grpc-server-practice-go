package services

import (
	"user-services/entities"
	"user-services/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(user *entities.User) (*entities.User, error) {
	result, err := s.userRepository.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) GetUserByEmail(email string) (*entities.User, error) {
	result, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}
