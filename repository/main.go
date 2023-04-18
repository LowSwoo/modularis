package repository

type Repository struct {
	User
	Post
}

func NewRepository(user User, post Post) *Repository {
	return &Repository{User: user, Post: post}
}
