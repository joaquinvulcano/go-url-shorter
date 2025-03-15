package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string, password string, db int) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisStore{client: client}, nil
}

func (s *RedisStore) SaveURL(ctx context.Context, shortURL, longURL string) error {
	return s.client.Set(ctx, shortURL, longURL, 0).Err()
}

func (s *RedisStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	longURL, err := s.client.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("URL not found")
	}
	return longURL, err
}
