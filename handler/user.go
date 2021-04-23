package handler

import (
	"bootcamp/entity"
	"bootcamp/services/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser func is used to create user data
func CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.RegistrationUserEntity
		var userService user.UserServiceInterface = user.UserService{}

		err := c.ShouldBindJSON(&data)
		if err != nil {
			fmt.Println(err)
		}

		createUser, err := userService.CreateUser(data)
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
}

// GetAllUser func is used to retrieve all data user
func GetAllUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userService *user.UserService = new(user.UserService)
		userData := userService.GetAllUser()
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}
