package auth

import (
	"bootcamp/entity"
	"bootcamp/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	LoginService(username, password string) (entity.DataUserEntity, error)
}

type AuthService struct {
	repository repository.UserRepository
}

func (service AuthService) LoginService(username, password string) (entity.DataUserEntity, error) {
	user, isExist := service.repository.GetUserByUserName(username)

	if !isExist {
		return entity.DataUserEntity{}, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.DataUserEntity{}, err
	}
	return user, nil
}
