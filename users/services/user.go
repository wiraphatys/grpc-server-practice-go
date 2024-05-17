package services

import "user-services/entities"

type UserService interface {
	CreateUser(user *entities.User) (*entities.User, error)
}
