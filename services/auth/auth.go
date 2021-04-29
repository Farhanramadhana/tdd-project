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
	Repository repository.UserRepositoryInterface
}

func (service AuthService) LoginService(username, password string) (entity.DataUserEntity, error) {
	user, err := service.Repository.GetUserByUserName(username)

	if err != nil {
		return entity.DataUserEntity{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.DataUserEntity{}, err
	}
	return user, nil
}
