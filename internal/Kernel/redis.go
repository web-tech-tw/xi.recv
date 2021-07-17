package Kernel

import (
	"github.com/go-redis/redis/v8"
	"github.com/star-inc/xi.recv/internal/Kernel/Config"
	"os"
	"strconv"
)

func NewRedisClient() *redis.Client {
	database, err := strconv.Atoi(os.Getenv(Config.RedisDB))
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv(Config.RedisAddr),
		Password: os.Getenv(Config.RedisPassword),
		DB:       database,
	})
}
