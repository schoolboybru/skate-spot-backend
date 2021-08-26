package cache

import (
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	*redis.Client
}

func NewCache() (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Cache{rdb}, nil

}
