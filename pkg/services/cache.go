package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client redis.UniversalClient
	ttl time.Duration
}

const apqPrefix = "apq"

func NewCache(ctx context.Context, redisAddr string, ttl time.Duration) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("cache could not be created: %w", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

func (c *Cache) Add(ctx context.Context, key string, val interface{}){
	c.client.Set(ctx, apqPrefix+key, val, c.ttl).Result()
}

func (c *Cache) Get(ctx context.Context, key string,) (interface{}, bool){
	s, err := c.client.Get(ctx, apqPrefix+key).Result()
	if err != nil {
		return struct{}{}, false
	}

	return s, true
}

