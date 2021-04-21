package main

import (
	"bootcamp/infra"
	"bootcamp/services/user"
)

func main() {
	s := infra.New(infra.Options{Port: 50001})
	s.Router("POST", "/user", user.CreateUser)
	s.Router("GET", "/user", user.GetAllUser)

	s.Start()
}
