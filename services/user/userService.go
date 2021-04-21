package user

import (
	"bootcamp/infra"
	"bootcamp/repository"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserPayload struct {
	FullName string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=10"`
	Role     string `validate:"required"`
}

// CreateUser func is used to create user data
func CreateUser(c *gin.Context) {
	var data UserPayload
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}

	var validate infra.Validator
	v := validate.NewValidator()
	err = v.Validate(data)

	if err != nil {
		err2 := strings.Split(err.Error(), ",")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err2,
		})
		return
	}

	name := SplitFullName(data.FullName)
	initialName := CreateInitialName(data.FullName)
	userName := GenerateUserName(name)
	ID := GenerateID()

	user := repository.DataUser{
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

	createUser := user.CreateUser()

	if createUser {
		c.JSON(http.StatusOK, gin.H{
			"status": "success create user data",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "failed create user data",
	})
}

// GetAllUser func is used to retrieve all data user
func GetAllUser(c *gin.Context) {
	userData := repository.GetAllUser()

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
		"data":   userData,
	})
	return
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
	data := repository.GetUserByName(name[0], name[2])
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
