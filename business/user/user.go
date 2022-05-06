package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Modify(newName, newEmail, newPassword string) User {
	return User{
		Name:     newName,
		Email:    newEmail,
		Password: newPassword,
	}
}
