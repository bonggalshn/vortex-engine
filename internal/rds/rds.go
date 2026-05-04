package rds

import (
	"context"
	"os"
	"vortex-engine/logger"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func EstablishRedisConnection() *redis.Client {
	logger.Info.Println("Establish Redis connection.")
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"), // Redis server address
		DB:   0,                                                       // Default DB index
	})

	logger.Info.Println("Redis connection established." + rdb.NodeAddress())

	return rdb
}
