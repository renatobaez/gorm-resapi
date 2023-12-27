package application

import (
	"github.com/renatobaez/gorm-resapi/internal/domain"
	"github.com/renatobaez/gorm-resapi/internal/infrastructure/repository"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUser(id string) (domain.User, error)
	CreateUser(domain.User) error
	DeleteUser(id string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (us *userService) GetAllUsers() ([]domain.User, error) {
	res, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (us *userService) GetUser(id string) (domain.User, error) {
	res, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}

func (us *userService) CreateUser(user domain.User) error {
	if err := us.userRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *userService) DeleteUser(id string) error {
	if err := us.userRepository.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
