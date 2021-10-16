package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/schoolboybru/location-service/internal/domain"
)

type CacheRepository interface {
	AddLocation(location *domain.Location) error
	GetLocation(id string) (*domain.Location, error)
}

func NewCache() (CacheRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	repo := &LocationCache{client: rdb}

	return repo, nil

}
