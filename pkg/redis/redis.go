package redis

import (
	"context"
	"errors"
	"time"


	"github.com/redis/go-redis/v9"
)

var (
	ErrCacheMiss = errors.New("redis cache miss")
	ErrCacheFull = errors.New("redis cache full")
	ErrCacheSet  = errors.New("redis cache set error")
)

type RedisCacheAdapter struct {
	client *redis.Client
}

func NewRedisProvider() *RedisCacheAdapter {
	rdb := redis.NewClient(&redis.Options{
		Addr:       ":6379",
		Password:   "",
		DB:         0,
		ClientName: "",
	})

	return &RedisCacheAdapter{
		client: rdb,
	}
}

func (r *RedisCacheAdapter) Set(ctx context.Context, key string, ttl time.Duration, data []byte) error {
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCacheAdapter) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrCacheMiss
		}
		return nil, err
	}

	return data, nil
}

func (r *RedisCacheAdapter) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCacheAdapter) Exists(ctx context.Context, key string) (bool, error) {
	return r.client.Exists(ctx, key).Val() == 1, nil
}

func (r *RedisCacheAdapter) Purge(ctx context.Context){
	r.client.FlushDB(ctx)
}
