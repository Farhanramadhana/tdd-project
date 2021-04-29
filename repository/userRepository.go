package repository

import (
	"bootcamp/entity"
	"container/list"
	"errors"
)

type UserRepositoryInterface interface {
	CreateUser(d entity.DataUserEntity) bool
	GetAllUser() ([]entity.DataUserEntity, error)
	GetUserByID(id string) (entity.DataUserEntity, error)
	GetUserByUserName(username string) (entity.DataUserEntity, error)
	GetUserByEmail(email string) (entity.DataUserEntity, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity, error)
}

type UserRepository struct{}

var l = list.New()

// CreateUser is to save data to linkedlist
func (u UserRepository) CreateUser(d entity.DataUserEntity) bool {
	l.PushFront(d)
	return true
}

// GetAllUser return all user detail
func (u UserRepository) GetAllUser() ([]entity.DataUserEntity, error) {
	var data []entity.DataUserEntity
	if l.Len() <= 0 {
		return data, errors.New("data doesn't exist")
	}

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value.(entity.DataUserEntity))
	}

	return data, nil
}

// GetUserById return specific user detail by id
func (u UserRepository) GetUserByID(id string) (entity.DataUserEntity, error) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			return data, nil
		}
	}
	return entity.DataUserEntity{}, errors.New("data doesn't exist")
}

// GetUserByUserName return specific user detail by username
func (u UserRepository) GetUserByUserName(username string) (entity.DataUserEntity, error) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if username == data.Username {
			return data, nil
		}
	}
	return entity.DataUserEntity{}, errors.New("user doesn't exist")
}

// GetUserByEmail return specific user detail by email
func (u UserRepository) GetUserByEmail(email string) (entity.DataUserEntity, error) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if email == data.Email {
			return data, nil
		}
	}
	return entity.DataUserEntity{}, errors.New("user doesn't exist")
}

// GetUserByName return specific user detail by name
func (u UserRepository) GetUserByName(firstName, lastName string) []entity.DataUserEntity {
	var user []entity.DataUserEntity

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if firstName == data.FirstName && lastName == data.LastName {
			user = append(user, data)
		}
	}
	return user
}

func (u UserRepository) DeleteUserByID(id string) (error) {
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			l.Remove(e)
			return nil
		}
	}

	return errors.New("user doesn't exist")
}

func (u UserRepository) UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity, error) {
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			l.Remove(e)
			l.PushFront(d)
			return d, nil
		}
	}

	return d, errors.New("user doesn't exist")
}
