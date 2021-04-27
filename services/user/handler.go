package user

import (
	"bootcamp/entity"
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
		var userService UserServiceInterface = UserService{}

		err := c.ShouldBindJSON(&data)
		if err != nil {
			fmt.Println(err)
		}

		dataUser, err := userService.CreateUser(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   dataUser,
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

		var userService UserServiceInterface = UserService{}
		userData, isExist := userService.GetAllUser()

		if !isExist {
			c.JSON(http.StatusBadRequest, gin.H{
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

		var userService UserServiceInterface = UserService{}
		userData, isExist := userService.GetUserByUsername(c.Param("username"))

		if !isExist {
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

		var userService UserServiceInterface = UserService{}
		userData, isExist := userService.GetUserByEmail(c.Param("email"))

		if !isExist {
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

		var userService UserServiceInterface = UserService{}
		userData, isExist := userService.GetUserByID(c.Param("id"))

		if !isExist {
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

		var userService UserServiceInterface = UserService{}
		userData, isExist := userService.DeleteUserByID(c.Param("id"))

		if !isExist {
			c.JSON(http.StatusBadRequest, gin.H{
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

		var userService UserServiceInterface = UserService{}
		userData, err := userService.UpdateUserByID(c.Param("id"), data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
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
