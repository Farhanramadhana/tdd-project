package main

import (
	"bootcamp/config"

	"bootcamp/services/auth"
	"bootcamp/services/user"
)

func main() {
	s := config.New(config.Options{Port: 50001})

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
