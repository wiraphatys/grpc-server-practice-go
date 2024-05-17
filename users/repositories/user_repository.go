package repositories

import (
	"user.services/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) InsertUser(user *entities.User) (*entities.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*entities.User, error) {
	var existedUser *entities.User
	err := r.db.First(&existedUser, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}
