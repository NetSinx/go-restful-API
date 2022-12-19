package user

import "github.com/NetSinx/go-restful-API/model/domain"

type UserRepository interface {
	Register(user domain.Register) error
	Delete(nameUser string) error
	Login(users domain.Login) error
}