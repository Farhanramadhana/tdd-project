package user

import (
	"bootcamp/config"
	"bootcamp/entity"
	"bootcamp/repository"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(data entity.RegistrationUserEntity) (entity.DataUserEntity, error)
	GetAllUser() ([]entity.DataUserEntity, error)
	GetUserByUserName(username string) (entity.DataUserEntity, bool)
	GetUserByID(id string) (entity.DataUserEntity, error)
	GetUserByEmail(email string) (entity.DataUserEntity, bool)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, data entity.UpdateUserEntity) (entity.DataUserEntity, error)
}

type UserService struct {
	Repository repository.UserRepositoryInterface
}

// CreateUser func is used to create user data
func (service UserService) CreateUser(data entity.RegistrationUserEntity) (entity.DataUserEntity, error) {
	var validate config.Validator

	v := validate.NewValidator()
	err := v.Validate(data)

	if err != nil {
		return entity.DataUserEntity{}, err
	}

	_, err = service.Repository.GetUserByEmail(data.Email)
	if err == nil {
		return entity.DataUserEntity{}, errors.New("email already exist")
	}

	name := SplitFullName(data.FullName)
	initialName := CreateInitialName(data.FullName)
	userName := GenerateUserName(name)
	ID := GenerateID()
	password := EncryptPassword(data.Password)

	user := entity.DataUserEntity{
		ID:          ID,
		FirstName:   name[0],
		MiddleName:  name[1],
		LastName:    name[2],
		Username:    userName,
		Role:        data.Role,
		InitialName: initialName,
		Email:       data.Email,
		Password:    password,
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	userData, err := service.Repository.CreateUser(user)
	return userData, err
}

// GetAllUser func is used to retrieve all user data
func (service UserService) GetAllUser() ([]entity.DataUserEntity, error) {
	userData, isExist := service.Repository.GetAllUser()
	return userData, isExist
}

// GetUserByUserName func is used to retrieve specific user data by username
func (service UserService) GetUserByUserName(username string) (entity.DataUserEntity, error) {
	userData, err := service.Repository.GetUserByUserName(username)
	return userData, err
}

// GetUserById func is used to retrieve specific user data by user id
func (service UserService) GetUserByID(id string) (entity.DataUserEntity, error) {
	userData, err := service.Repository.GetUserByID(id)
	return userData, err
}

// GetUserByEmail func is used to retrieve specific user data by email
func (service UserService) GetUserByEmail(email string) (entity.DataUserEntity, error) {
	userData, err := service.Repository.GetUserByEmail(email)
	return userData, err
}

// DeleteUserByID func is used to delete specific user data by id
func (service UserService) DeleteUserByID(id string) error {
	_, err := service.Repository.GetUserByID(id)
	if err != nil {
		return err
	}

	_ = service.Repository.DeleteUserByID(id)
	return nil
}

// UpdateUserByID func is used to update specific user data by id
func (service UserService) UpdateUserByID(id string, data entity.UpdateUserEntity) (entity.DataUserEntity, error) {
	var validate config.Validator
	var user entity.DataUserEntity

	v := validate.NewValidator()
	err := v.Validate(data)

	if err != nil {
		return entity.DataUserEntity{}, err
	}

	existingData, err := service.Repository.GetUserByID(id)

	if err != nil {
		return entity.DataUserEntity{}, err
	}

	user.ID = existingData.ID
	if data.FullName != "" {
		var initialName = existingData.InitialName
		var userName = existingData.Username

		name := SplitFullName(data.FullName)
		initialName = CreateInitialName(data.FullName)

		if existingData.FirstName != name[0] || existingData.LastName != name[2] {
			userName = GenerateUserName(name)
		}

		user.FirstName = name[0]
		user.MiddleName = name[1]
		user.LastName = name[2]
		user.InitialName = initialName
		user.Username = userName
	}

	data.Password = existingData.Password
	if data.Password != "" {
		password := EncryptPassword(data.Password)
		user.Password = password
	}

	user.Role = existingData.Role
	if data.Role != "" && existingData.Role != data.Role {
		user.Role = data.Role
	}

	user.Email = existingData.Email
	if data.Email != "" && existingData.Email != data.Email {
		user.Email = data.Email
	}

	user.UpdatedAt = time.Now().Format(time.RFC3339)

	userData, err := service.Repository.UpdateUserByID(id, user)
	return userData, err
}

// SplitFullName is func to split fullname to first, middle, and last of Name
func SplitFullName(fullName string) []string {
	s := strings.Split(fullName, " ")
	length := len(s)

	var firstName, middleName, lastName string
	firstName = s[0]

	if length == 1 {
		return []string{firstName, middleName, lastName}
	}

	var middle []string
	if length >= 3 {
		middle = s[1 : length-1]
		middleName = strings.Join(middle, " ")
	}

	lastName = s[length-1]
	return []string{firstName, middleName, lastName}
}

// CreateInitialName is func to split create initial name
func CreateInitialName(fullName string) string {
	s := strings.Split(fullName, " ")

	var initialName []byte
	for _, v := range s {
		initialName = append(initialName, v[0])
	}
	return strings.ToUpper(string(initialName))
}

// GenerateUserName is func to split create initial name
func GenerateUserName(name []string) string {
	// find userName, if already exist add index number
	//name[0] is firstname, name[1] is middlename, name[2] is lastname

	var service repository.UserRepository
	userName := name[0]

	if name[2] != "" {
		userName += "." + name[2]
	}

	_, err := service.GetUserByUserName(userName)

	if err != nil {
		return userName
	}

	i := 1
	var userName2 string
	for err == nil {
		userName2 = userName + strconv.Itoa(i)
		_, err = service.GetUserByUserName(userName2)
		i++
	}
	userName = userName2
	return userName
}

// GenerateID is to generate random id with 12 alphanumeric
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 12
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	id := b.String()
	return id
}

// EncryptPassword is func to encryp password using bcrypt
func EncryptPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		fmt.Print(err.Error())
	}

	return (string(hash))
}
