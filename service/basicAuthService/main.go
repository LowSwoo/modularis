package basicAuthService

import (
	"modularis/models"
	"modularis/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) SignUp(user *models.User) (string, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return "successfully", nil
	} else {
		return "", err
	}

}

func (s Service) SignIn(login, password string) (string, error) {
	usr, err := s.repo.GetUserInfo(login)
	if err != nil {
		return "", err
	}
	if usr.Password == password {
		return "password correct", nil
	} else {
		return "password incorrect", nil
	}

}
