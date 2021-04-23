package repository

import (
	"bootcamp/entity"
)

type UserRepositoryInterface interface {
	CreateUser(d entity.DataUserEntity) bool
	GetAllUser() []entity.DataUserEntity
	GetUserById(id string) entity.DataUserEntity
	GetUserByName(firstName, lastName string) []entity.DataUserEntity
}

type UserRepository struct{}

var data []entity.DataUserEntity

//CreateUser is to save data to database
func (u UserRepository) CreateUser(d entity.DataUserEntity) bool {
	data = append(data, d)
	return true
}

// FindByID return specific user detail by id
func (u UserRepository) GetAllUser() []entity.DataUserEntity {
	return data
}

// FindByID return specific user detail by id
func (u UserRepository) GetUserById(id string) entity.DataUserEntity {
	return entity.DataUserEntity{
		ID:          "123",
		FirstName:   "farhan",
		MiddleName:  "stona",
		LastName:    "ramadhana",
		Username:    "frans",
		Role:        "admin",
		InitialName: "FSR",
		Email:       "frans@kata.ai",
		Password:    "pass",
		UpdatedAt:   "2020-02-02",
	}
}

func (u UserRepository) GetUserByName(firstName, lastName string) []entity.DataUserEntity {
	var user []entity.DataUserEntity
	for _, v := range data {
		if firstName == v.FirstName && lastName == v.LastName {
			user = append(user, v)
		}
	}

	return user
}
