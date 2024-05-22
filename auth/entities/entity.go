package entities

import "time"

type User struct {
	UID       int       `gorm:"primaryKey;not null" json:"user_id"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Tel       string    `gorm:"not null" json:"tel"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// request data
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
