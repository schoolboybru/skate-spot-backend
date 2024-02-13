package services

import "github.com/schoolboybru/skate-spot/internal/domain"

type PostService interface {
	GetPost(id string) (*domain.Post, error)
	AddPost(post *domain.Post) error
}

func (s *service) GetPost(id string) (*domain.Post, error) {
	return s.postgresRepo.GetPost(id)
}

func (s *service) AddPost(post *domain.Post) error {
	return s.postgresRepo.AddPost(post)
}
