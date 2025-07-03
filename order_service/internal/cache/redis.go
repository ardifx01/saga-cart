package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &Redis{
		rdb: rdb,
	}
}

func (c *Redis) Set(ctx context.Context, key string, value interface{}, minute int) error {
	err := c.rdb.Set(ctx, key, value, time.Duration(minute)*time.Minute)
	if err != nil {
		return fmt.Errorf("[Redis] error while set value: %v", err.Err())
	}
	return nil
}

func (c *Redis) Get(ctx context.Context, key string) (string, error) {
	value, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("[Redis] error while get value: %v", err.Error())
	}
	return value, nil
}
