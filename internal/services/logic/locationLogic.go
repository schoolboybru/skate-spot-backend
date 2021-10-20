package logic

import (
	"github.com/schoolboybru/location-service/internal/domain"
	"github.com/schoolboybru/location-service/internal/repositories"
	"github.com/schoolboybru/location-service/internal/repositories/cache"
	"github.com/schoolboybru/location-service/internal/services"
)

type service struct {
	postgresRepo repositories.LocationRepository
	cacheRepo    cache.CacheRepository
}

func New(repository repositories.LocationRepository, cache cache.CacheRepository) services.LocationService {
	return &service{
		postgresRepo: repository,
		cacheRepo:    cache,
	}
}

func (s *service) AddLocation(l *domain.Location) error {
	s.cacheRepo.AddLocation(l)
	return s.postgresRepo.AddLocation(l)
}

func (s *service) GetLocation(id string) (*domain.Location, error) {

	l, err := s.cacheRepo.GetLocation(id)

	if err != nil {
		println(err)
	}

	if l != nil {
		return l, nil
	}

	return s.postgresRepo.GetLocation(id)
}

func (s *service) GetLocationsByCity(id string) (*[]domain.Location, error) {
	return s.postgresRepo.GetLocationsByCity(id)
}

func (s *service) GetLocationsByCountry(id string) (*[]domain.Location, error) {
	return s.postgresRepo.GetLocationsByCountry(id)
}
