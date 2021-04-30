package repository

import (
	"bootcamp/entity"
	"errors"
	"sort"
)

type UserRepositoryInterface interface {
	CreateUser(d entity.DataUserEntity) (entity.DataUserEntity, error)
	GetAllUser() ([]entity.DataUserEntity, error)
	GetUserByID(id string) (entity.DataUserEntity, error)
	GetUserByUserName(username string) (entity.DataUserEntity, error)
	GetUserByEmail(email string) (entity.DataUserEntity, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity, error)
}

type UserRepository struct{}

var list = make(map[string]map[string]interface{})

func init() {
	var superAdmin = entity.DataUserEntity{
		ID:          "VGU0pGrzLkqK",
		FirstName:   "super",
		MiddleName:  "",
		LastName:    "admin",
		Username:    "super.admin",
		Role:        "admin",
		InitialName: "SA",
		Email:       "superadmin@admin.com",
		Password:    "$2a$04$9RuasgZY3gxe6xGJBep/cO1SjWWa84nzhx5MzQVs/e0IjdXB/Qqxi",
		UpdatedAt:   "2021-04-30T10:12:13+07:00",
	}

	list["ID"] = make(map[string]interface{})
	list["UserName"] = make(map[string]interface{})
	list["Email"] = make(map[string]interface{})

	list["ID"][superAdmin.ID] = superAdmin
	list["UserName"][superAdmin.Username] = superAdmin
	list["Email"][superAdmin.Email] = superAdmin
}

// CreateUser is to save data to linkedlist
func (u UserRepository) CreateUser(d entity.DataUserEntity) (entity.DataUserEntity, error) {
	list["ID"][d.ID] = d
	list["UserName"][d.Username] = d
	list["Email"][d.Email] = d

	return d, nil
}

// GetAllUser return all user detail
func (u UserRepository) GetAllUser() ([]entity.DataUserEntity, error) {
	var data []entity.DataUserEntity

	for _, v := range list["ID"] {
		data = append(data, v.(entity.DataUserEntity))
	}
	
	sort.SliceStable(data, func(i,j int) bool {
		return data[i].UpdatedAt > data[j].UpdatedAt
	})
	return data, nil
}

// GetUserById return specific user detail by id
func (u UserRepository) GetUserByID(id string) (entity.DataUserEntity, error) {
	if list["ID"][id] == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}

	return list["ID"][id].(entity.DataUserEntity), nil
}

// GetUserByUserName return specific user detail by username
func (u UserRepository) GetUserByUserName(username string) (entity.DataUserEntity, error) {
	if list["UserName"][username] == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}

	return list["UserName"][username].(entity.DataUserEntity), nil
}

// GetUserByEmail return specific user detail by email
func (u UserRepository) GetUserByEmail(email string) (entity.DataUserEntity, error) {
	if list["Email"][email] == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}

	return list["Email"][email].(entity.DataUserEntity), nil
}

func (u UserRepository) DeleteUserByID(id string) error {
	data := list["ID"][id].(entity.DataUserEntity)
	delete(list["ID"], id)
	delete(list["Email"], data.Email)
	delete(list["UserName"], data.Username)
	return nil
}

func (u UserRepository) UpdateUserByID(id string, d entity.DataUserEntity) (entity.DataUserEntity, error) {
	list["ID"][d.ID] = d
	list["UserName"][d.Username] = d
	list["Email"][d.Email] = d

	return d, nil
}
