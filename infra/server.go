package infra

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Port int
}

type Server struct {
	ginServer *gin.Engine
	options Options
}

func (s *Server) Router(method string, path string, handler func(*gin.Context)) {
	if method == "GET" {
		s.ginServer.GET(path, handler)
	} else if method == "POST" {
		s.ginServer.POST(path, handler)
	} else if method == "PUT" {
		s.ginServer.PUT(path, handler)
	}else {
		fmt.Sprintf("method don't allowed")
	}
}

func New(options Options) *Server {
	return &Server{
		ginServer: gin.Default(),
		options:   options,
	}
}

func (s *Server) Start() {
	s.ginServer.Run(fmt.Sprintf(":%v", s.options.Port))
}