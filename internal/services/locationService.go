package services

import (
	"github.com/schoolboybru/location-service/internal/domain"
)

type LocationService interface {
	AddLocation(location *domain.Location) error
	GetLocation(id string) (*domain.Location, error)
	GetLocationsByCity(id string) (*[]domain.Location, error)
	GetLocationsByCountry(id string) (*[]domain.Location, error)
}

func (s *service) AddLocation(l *domain.Location) error {
	return s.postgresRepo.AddLocation(l)
}

func (s *service) GetLocation(id string) (*domain.Location, error) {
	return s.postgresRepo.GetLocation(id)
}

func (s *service) GetLocationsByCity(id string) (*[]domain.Location, error) {
	return s.postgresRepo.GetLocationsByCity(id)
}

func (s *service) GetLocationsByCountry(id string) (*[]domain.Location, error) {
	return s.postgresRepo.GetLocationsByCountry(id)
}
