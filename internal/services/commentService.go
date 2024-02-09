package services

import (
    "github.com/schoolboybru/location-service/internal/domain"
)

type CommentService interface {
    AddComment(*domain.Comment) error
    GetComments(postId string) (*[]domain.Comment, error)
}

func (s *service) AddComment(comment *domain.Comment) error {
    return s.postgresRepo.AddComment(comment)
}

func (s *service) GetComments(postId string) (*[]domain.Comment, error) {
    return s.postgresRepo.GetComments(postId)
}
