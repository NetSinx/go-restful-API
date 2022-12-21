package user

import (
	"html"
	"strings"
	"github.com/NetSinx/go-restful-API/model/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userrepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userrepository {
	return &userrepository{
		DB: db,
	}
}

func (repository *userrepository) Register(user domain.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	err = repository.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *userrepository) Delete(nameUser string) error {
	var user domain.User

	err := repository.DB.Where("username = ?", nameUser).First(&user).Error
	if err != nil {
		return err
	}

	repository.DB.Delete(&user)

	return nil
}

func (repository *userrepository) Login(users domain.Login) error {
	var u domain.User

	err := repository.DB.Where("username = ?", users.Username).Take(&u).Error
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(users.Password))
	if err != nil {
		return err
	}

	return nil
}