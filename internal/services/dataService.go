package services

import (
	"github.com/schoolboybru/location-service/internal/repositories/cache"
	"github.com/schoolboybru/location-service/internal/repositories/postgres"
)

type service struct {
	postgresRepo postgres.DatabaseRepository
	cacheRepo    cache.CacheRepository
}

type Services interface {
    CommentService
    LocationService
}

func New(repository postgres.DatabaseRepository) Services {
	return &service{
		postgresRepo: repository,
	}
}
