package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Port int
}

type Server struct {
	GinServer  *gin.Engine
	options    Options
}

func New(options Options) *Server {
	return &Server{
		GinServer: gin.Default(),
		options:   options,
	}
}

func (s *Server) Start() {
	s.GinServer.Run(fmt.Sprintf(":%v", s.options.Port))
}
