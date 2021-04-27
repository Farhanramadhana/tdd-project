package repository

import (
	"bootcamp/entity"
	"container/list"
)

type UserRepositoryInterface interface {
	CreateUser(d entity.DataUserEntity) entity.DataUserEntity
	GetAllUser() ([]entity.DataUserEntity, bool)
	GetUserByID(id string) (entity.DataUserEntity, bool)
	GetUserByUserName(username string) (entity.DataUserEntity, bool)
	GetUserByEmail(email string) (entity.DataUserEntity, bool)
	GetUserByName(firstName, lastName string) []entity.DataUserEntity
	DeleteUserByID(id string) (entity.DataUserEntity, bool)
	UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity)
}

type UserRepository struct{}

var l = list.New()

// CreateUser is to save data to linkedlist
func (u UserRepository) CreateUser(d entity.DataUserEntity) entity.DataUserEntity {
	l.PushFront(d)
	return d
}

// GetAllUser return all user detail
func (u UserRepository) GetAllUser() ([]entity.DataUserEntity, bool) {
	var data []entity.DataUserEntity
	if l.Len() <= 0 {
		return data, false
	}

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value.(entity.DataUserEntity))
	}

	return data, true
}

// GetUserById return specific user detail by id
func (u UserRepository) GetUserByID(id string) (entity.DataUserEntity, bool) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			return data, true
		}
	}
	return entity.DataUserEntity{}, false
}

// GetUserByUserName return specific user detail by username
func (u UserRepository) GetUserByUserName(username string) (entity.DataUserEntity, bool) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if username == data.Username {
			return data, true
		}
	}
	return entity.DataUserEntity{}, false
}

// GetUserByEmail return specific user detail by email
func (u UserRepository) GetUserByEmail(email string) (entity.DataUserEntity, bool) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if email == data.Email {
			return data, true
		}
	}
	return entity.DataUserEntity{}, false
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

func (u UserRepository) DeleteUserByID(id string) (entity.DataUserEntity, bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			l.Remove(e)
			return e.Value.(entity.DataUserEntity), true
		}
	}

	return entity.DataUserEntity{}, false
}

func (u UserRepository) UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity) {
	for e := l.Front(); e != nil; e = e.Next() {
		data := e.Value.(entity.DataUserEntity)
		if id == data.ID {
			l.Remove(e)
			l.PushFront(d)
			return d
		}
	}

	return d
}
