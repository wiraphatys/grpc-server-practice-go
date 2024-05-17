package repositories

import "user-services/entities"

type UserRepository interface {
	InsertUser(user *entities.User) (*entities.User, error)
}
