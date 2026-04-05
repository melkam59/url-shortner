package services

import (
	shortner "github.com/melkam59/url-shortner/internal/services/shortner"
	redis "github.com/melkam59/url-shortner/internal/storage"
)

type URLShortner struct {
	ShortnerService shortner.IShortnerService
	RedisService    redis.IRedisService
}

func NewURLShortner() (URLShortner, error) {
	app := URLShortner{}
	app.ShortnerService = shortner.NewShortnerService()
	app.RedisService = redis.NewRedisService("your-redis-host", "your-redis-port", "your-redis-password")

	return app, nil
}
