package models

import (
	"github.com/renatobaez/gorm-resapi/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null; unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}

func NewUserModel(firstName, lastName, email string) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
