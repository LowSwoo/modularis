package service

import "modularis/models"

type Authorization interface {
	SignUp(user *models.User) (string, error)
	SignIn(login, password string) (string, error)
}

type Posting interface {
	CreatePost(post *models.Post) error
	RemovePost(id string) error
	GetPost(id string) (*models.Post, error)
}
