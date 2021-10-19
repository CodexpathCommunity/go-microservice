package kv

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	ctx   context.Context
	redis *redis.Client
}

func NewRedisClient(ctx context.Context, opts *redis.Options) (KV, error) {
	redis := redis.NewClient(opts)
	if _, err := redis.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &redisClient{ctx: ctx, redis: redis}, nil
}

func (r *redisClient) Get(key string) (string, error) {
	return r.redis.Get(r.ctx, key).Result()
}

func (r *redisClient) Set(key string, value interface{}) error {
	return r.redis.Set(r.ctx, key, value, 0).Err()
}

func (r *redisClient) SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	return r.redis.Set(r.ctx, key, value, ttl).Err()
}

func (r *redisClient) Delete(key string) error {
	return r.redis.Del(r.ctx, key).Err()
}
