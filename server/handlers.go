package server

import (
	"github.com/gin-gonic/gin"
	"modularis/models"
	"net/http"
)

func (s *Server) home(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (s *Server) signUp(c *gin.Context) {
	answer, err := s.svc.SignUp(models.NewUser(c.PostForm("login"), c.PostForm("password")))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, answer)
	}

}

func (s *Server) signIn(c *gin.Context) {
	answer, err := s.svc.SignIn(c.PostForm("login"), c.PostForm("password"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, answer)
	}

}
