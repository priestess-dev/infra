package redis

import (
	"context"
	"github.com/priestess-dev/infra/cache"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr         string `yaml:"address"` // address of redis server (default: localhost:6379)
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	Prefix       string `yaml:"prefix"`
	KeySeparator string `yaml:"key_separator"`
}

func NewCache(config Config) (cache.Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &accessor{
		db:     client,
		config: &config,
		kb: keyBuilder{
			localRedisKeyPrefix: config.Prefix,
			redisKeySeparator:   config.KeySeparator,
		},
	}, nil
}
