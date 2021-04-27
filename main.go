package main

import (
	"bootcamp/config"
	"bootcamp/entity"
	"bootcamp/services/auth"
	"bootcamp/services/user"
	"fmt"
)

func main() {
	s := config.New(config.Options{Port: 50001})

	// predifines admin user
	superAdmin := entity.RegistrationUserEntity{
		FullName: "super admin",
		Email:    "superadmin@admin.com",
		Password: "admin12345",
		Role:     "admin",
	}

	var userService user.UserServiceInterface = user.UserService{}
	createUser, err := userService.CreateUser(superAdmin)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(createUser)
	authorized := s.GinServer.Group("/")
	authorized.Use(auth.LoginHandler())
	{
		// admin only
		authorized.POST("/user", user.CreateUserHandler())
		authorized.DELETE("/user/:id", user.DeleteUserByID())
		authorized.PATCH("/user/:id", user.UpdateUserByID())

		// all user type
		authorized.GET("/user", user.GetAllUserHandler())
		authorized.GET("/user/username/:username", user.GetUserByUsernameHandler())
		authorized.GET("/user/email/:email", user.GetUserByEmailHandler())
		authorized.GET("/user/id/:id", user.GetUserByIDHandler())
	}
	s.Start()
}
