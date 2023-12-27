package repository

import (
	"errors"

	"github.com/renatobaez/gorm-resapi/internal/domain"
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (domain.User, error)
	CreateUser(user domain.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) GetAllUsers() ([]domain.User, error) {
	var usersModel []models.User
	var usersDomain []domain.User
	ur.DB.Find(&usersModel)

	for _, userModel := range usersModel {
		usersDomain = append(usersDomain, userModel.ToDomain())
	}

	return usersDomain, nil
}

func (ur *userRepository) GetUserByID(id string) (domain.User, error) {
	var user models.User
	ur.DB.First(&user, id)
	if user.ID == 0 {
		return domain.User{}, errors.New("User not found")
	}
	ur.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	return user.ToDomain(), nil
}

func (ur *userRepository) CreateUser(user domain.User) error {
	userModel := models.NewUserModel(user.FirstName, user.LastName, user.Email)
	createdUser := ur.DB.Create(&userModel)
	if createdUser.Error != nil {
		return createdUser.Error
	}

	return nil
}

func (ur *userRepository) DeleteUser(id string) error {
	var user models.User
	ur.DB.First(&user, id)
	if user.ID == 0 {
		return errors.New("User not found")
	}
	userDeleted := ur.DB.Delete(&user)
	if userDeleted.Error != nil {
		return userDeleted.Error
	}
	return nil
}
