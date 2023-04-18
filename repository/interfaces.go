package repository

import "modularis/models"

type User interface {
	CreateUser(user *models.User) error
	RemoveUSer(login string) error
	GetUserInfo(login string) (*models.User, error)
}

type Post interface {
	CreatePost(post *models.Post) error
	RemovePost(id string) error
	GetPost(id string) (*models.Post, error)
}
