package logic

import (
	"github.com/schoolboybru/location-service/internal/domain"
	"github.com/schoolboybru/location-service/internal/repositories"
	"github.com/schoolboybru/location-service/internal/services"
)

type service struct {
	postgresRepo repositories.LocationRepository
	//cacheRepo    CacheRepository
}

func New(repository repositories.LocationRepository) services.LocationService {
	return &service{postgresRepo: repository}
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
