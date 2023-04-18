package server

import (
	"github.com/gin-gonic/gin"
	"modularis/service"
)

type Server struct {
	router *gin.Engine
	svc    *service.Service
	cfg    struct {
		Host string
		Port string
	}
}

func NewServer(config struct {
	Host string
	Port string
}, services *service.Service) *Server {
	return &Server{
		router: gin.Default(),
		cfg:    config,
		svc:    services,
	}
}

func (s *Server) initHandlers() {
	s.router.GET("/", s.home)
	usrGroup := s.router.Group("/user")
	{
		usrGroup.POST("/sign-up", s.signUp)
		usrGroup.POST("/sign-in", s.signIn)
	}
}

func (s *Server) Run() {
	s.initHandlers()
	s.router.Run(s.cfg.Host + ":" + s.cfg.Port)
}
