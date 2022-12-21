package user

import (
	"errors"
	"github.com/NetSinx/go-restful-API/model/domain"
	repository "github.com/NetSinx/go-restful-API/repository/user"
)

type userservice struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userservice {
	return &userservice{
		UserRepository: userRepo,
	}
}

func (service *userservice) Register(user domain.User) error {
	err := service.UserRepository.Register(user)
	if err != nil {
		return errors.New("user cannot registered")
	}

	return nil
}

func (service *userservice) Delete(nameUser string) error {
	err := service.UserRepository.Delete(nameUser)
	if err != nil {
		return errors.New("user cannot deleted")
	}

	return nil
}

func (service *userservice) Login(users domain.Login) error {
	err := service.UserRepository.Login(users)
	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}