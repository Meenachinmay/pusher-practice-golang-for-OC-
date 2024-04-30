package config

import (
	"github.com/pusher/pusher-http-go/v5"
	"github.com/redis/go-redis/v9"
	"pusher-practice/internal/database"
)

type ApiConfig struct {
	DB           *database.Queries
	RedisClient  *redis.Client
	PusherClient *pusher.Client
}
