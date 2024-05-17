package repositories

import "user.services/entities"

type UserRepository interface {
	FindUserByEmail(email string) (*entities.User, error)
	InsertUser(user *entities.User) (*entities.User, error)
}
