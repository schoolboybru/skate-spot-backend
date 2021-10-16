package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/schoolboybru/location-service/internal/domain"
)

var ctx = context.Background()

type LocationCache struct {
	client *redis.Client
}

func (l *LocationCache) GetLocation(name string) (*domain.Location, error) {
	location := domain.Location{}

	c := l.client

	err := c.Get(ctx, name).Scan(&location)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (l *LocationCache) AddLocation(location *domain.Location) error {
	c := l.client

	err := c.Set(ctx, location.Name, location, 0).Err()

	if err != nil {
		return err
	}

	return nil

}
