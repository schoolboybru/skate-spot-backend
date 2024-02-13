package services

import (
	"github.com/schoolboybru/skate-spot/internal/repositories/cache"
	"github.com/schoolboybru/skate-spot/internal/repositories/postgres"
)

type service struct {
	postgresRepo postgres.DatabaseRepository
	cacheRepo    cache.CacheRepository
}

type Services interface {
	CommentService
	LocationService
	PostService
}

func New(repository postgres.DatabaseRepository) Services {
	return &service{
		postgresRepo: repository,
	}
}
