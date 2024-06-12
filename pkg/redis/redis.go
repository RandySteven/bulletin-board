package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"task_mission/pkg/config"
)

type RedisRepositories struct {
	client *redis.Client
}

func NewRedisRepositories(config *config.Config) *RedisRepositories {
	addr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	return &RedisRepositories{
		client: client,
	}
}
