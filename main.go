package main

import (
	"bootcamp/config"
	"bootcamp/services/user"

)

func main() {
	s := config.New(config.Options{Port: 50001})

	s.Router("POST", "/user", user.CreateUserHandler())
	s.Router("GET", "/user", user.GetAllUserHandler())

	s.Start()
}
