package user

import (
	"bootcamp/config"
	"bootcamp/entity"
	"bootcamp/repository"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(data entity.RegistrationUserEntity) (entity.DataUserEntity, error)
	GetAllUser() ([]entity.DataUserEntity, bool)
	GetUserByUsername(username string) (entity.DataUserEntity, bool)
	GetUserById(id string) (entity.DataUserEntity, bool)
	GetUserByEmail(email string) (entity.DataUserEntity, bool)
	DeleteUserById(id string) (entity.DataUserEntity, bool)
}

type UserService struct {
	repository repository.UserRepository
}

// CreateUser func is used to create user data
func (service UserService) CreateUser(data entity.RegistrationUserEntity) (entity.DataUserEntity, error) {
	var validate config.Validator

	v := validate.NewValidator()
	err := v.Validate(data)

	if err != nil {
		return entity.DataUserEntity{}, err
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

	dataUser := service.repository.CreateUser(user)
	return dataUser, nil
}

// GetAllUser func is used to retrieve all data user
func (service UserService) GetAllUser() ([]entity.DataUserEntity, bool) {
	userData, isExist := service.repository.GetAllUser()
	return userData, isExist
}

func (service UserService) GetUserByUsername(username string) (entity.DataUserEntity, bool) {
	userData, isExist := service.repository.GetUserByUserName(username)
	return userData, isExist
}

func (service UserService) GetUserById(id string) (entity.DataUserEntity, bool) {
	userData, isExist := service.repository.GetUserById(id)
	return userData, isExist
}

func (service UserService) GetUserByEmail(email string) (entity.DataUserEntity, bool) {
	userData, isExist := service.repository.GetUserByEmail(email)
	return userData, isExist
}

func (service UserService) DeleteUserById(id string) (entity.DataUserEntity, bool) {
	userData, isExist := service.repository.DeleteUserById(id)
	return userData, isExist
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
	var service repository.UserRepositoryInterface = repository.UserRepository{}
	data := service.GetUserByName(name[0], name[2])
	totalData := len(data)

	var userName string
	if totalData >= 1 && name[2] != "" {
		index := totalData + 1
		userName = name[0] + "." + name[2] + strconv.Itoa(index)
	} else if totalData >= 1 && name[2] == "" {
		index := totalData + 1
		userName = name[0] + strconv.Itoa(index)
	} else if totalData == 0 && name[2] != "" {
		userName = name[0] + "." + name[2]
	} else {
		userName = name[0]
	}

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

func EncryptPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		fmt.Print(err.Error())
	}

	return (string(hash))
}
