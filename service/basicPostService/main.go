package basicPostService

import (
	"modularis/models"
	"modularis/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

func (s *Service) RemovePost(id string) error {
	return s.repo.RemovePost(id)
}

func (s *Service) GetPost(id string) (*models.Post, error) {
	return s.repo.GetPost(id)
}
