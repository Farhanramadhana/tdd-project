package main

import (
	"bootcamp/infra"
	"bootcamp/handler"
)

func main() {
	s := infra.New(infra.Options{Port: 50001})
	
	s.Router("POST", "/user", handler.CreateUserHandler())
	s.Router("GET", "/user", handler.GetAllUserHandler())

	s.Start()
}
