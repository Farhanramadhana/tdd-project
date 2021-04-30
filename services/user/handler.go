package user

import (
	"bootcamp/entity"
	"bootcamp/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler func is used to create user data
func CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		var data entity.RegistrationUserEntity
		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}

		_ = c.ShouldBindJSON(&data)

		userData, err := userService.CreateUser(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"data": userData,
		})
	}
}

// GetAllUser func is used to retrieve all data user
func GetAllUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" && role != "user" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		userData, err := userService.GetAllUser()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   "null",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}

func GetUserByUsernameHandler() gin.HandlerFunc {
	fmt.Print("lalalal")
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" && role != "user" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		userData, err := userService.GetUserByUserName(c.Param("username"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   "null",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}

func GetUserByEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" && role != "user" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		userData, err := userService.GetUserByEmail(c.Param("email"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   "null",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}

func GetUserByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" && role != "user" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		userData, err := userService.GetUserByID(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   "null",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}

func DeleteUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		err := userService.DeleteUserByID(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"message":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func UpdateUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("Role")

		if role != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "forbidden",
			})
			return
		}

		var data entity.UpdateUserEntity
		err := c.ShouldBindJSON(&data)
		if err != nil {
			fmt.Println(err)
		}

		repositoryService := repository.UserRepository{}
		userService := UserService{repositoryService}
		userData, err := userService.UpdateUserByID(c.Param("id"), data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userData,
		})
	}
}
