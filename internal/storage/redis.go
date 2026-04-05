package storage

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	CACHE_DURATION = 1 * time.Hour
)

type IRedisService interface {
	SetKey(ctx context.Context, originalUrl, shortUrl string) error
	GetValue(ctx context.Context, key string) (string, error)
}

type RedisClientService struct {
	client *redis.Client
}

func NewRedisService(host, post, password string) IRedisService {
	redisAddress := fmt.Sprintf("%s:%s", host, post)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: password,
	})

	return &RedisClientService{
		client: client,
	}
}

func (r *RedisClientService) SetKey(ctx context.Context, originalUrl, shortUrl string) error {
	err := r.client.Set(ctx, shortUrl, originalUrl, CACHE_DURATION).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClientService) GetValue(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil

}
