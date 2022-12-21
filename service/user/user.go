package user

import "github.com/NetSinx/go-restful-API/model/domain"

type UserService interface {
	Register(user domain.User) error
	Delete(nameUser string) error
	Login(users domain.Login) error
}