package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, minute int) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Redis struct {
	rdb *redis.Client
}

func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// ensure redis cache flush, whenever start the app
	rdb.FlushDB(context.Background())
	return &Redis{
		rdb: rdb,
	}
}

func (c *Redis) Set(ctx context.Context, key string, value interface{}, minute int) error {
	err := c.rdb.Set(ctx, key, value, time.Duration(minute)*time.Millisecond)
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

func (c *Redis) Delete(ctx context.Context, key string) error {
	_, err := c.rdb.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("[Redis] error delete given key: %v", err.Error())
	}
	return nil
}
