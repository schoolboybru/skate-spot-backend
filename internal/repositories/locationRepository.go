package repositories

import "github.com/schoolboybru/location-service/internal/domain"

type LocationRepository interface {
	AddLocation(location *domain.Location) error
	GetLocation(id string) (*domain.Location, error)
	GetLocationsByCity(id string) (*[]domain.Location, error)
	GetLocationsByCountry(id string) (*[]domain.Location, error)
}
