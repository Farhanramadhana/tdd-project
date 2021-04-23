package user

import (
	"bootcamp/entity"
	"bootcamp/infra"
	"bootcamp/repository"
	"fmt"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type UserServiceInterface interface {
	CreateUser(data entity.RegistrationUserEntity) (bool, error)
	GetAllUser() []entity.DataUserEntity
}

type UserService struct {
	repository repository.UserRepository
}

// CreateUser func is used to create user data
func (service UserService) CreateUser(data entity.RegistrationUserEntity) (bool, error) {
	var validate infra.Validator

	v := validate.NewValidator()
	err := v.Validate(data)

	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"status":  "Error",
		// 	"message": err2,
		// })
		return false, err
	}

	name := SplitFullName(data.FullName)
	initialName := CreateInitialName(data.FullName)
	userName := GenerateUserName(name)
	ID := GenerateID()

	user := entity.DataUserEntity{
		ID:          ID,
		FirstName:   name[0],
		MiddleName:  name[1],
		LastName:    name[2],
		Username:    userName,
		Role:        data.Role,
		InitialName: initialName,
		Email:       data.Email,
		Password:    data.Password,
		UpdatedAt:   "2020-02-02",
	}

	fmt.Print(user)
	createUser := service.repository.CreateUser(user)
	if createUser {
		return true, nil
	}

	return false, errors.New("failed create user data")
}

// GetAllUser func is used to retrieve all data user
func (service UserService) GetAllUser() []entity.DataUserEntity {
	// var repo repository.UserRepositoryInterface = repository.UserRepository{}
	userData := service.repository.GetAllUser()
	return userData
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
