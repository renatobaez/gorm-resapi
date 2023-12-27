package domain

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidUser = errors.New("invalid user")
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Tasks     []Task
}

func (u *User) Validate() error {
	if u.FirstName == "" || u.LastName == "" || u.Email == "" {
		return ErrInvalidUser
	}
	return nil
}
