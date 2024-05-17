package entities

import "time"

type User struct {
	UID       int       `gorm:"primaryKey;not null" json:"user_id"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Tel       string    `gorm:"not null" json:"tel"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
