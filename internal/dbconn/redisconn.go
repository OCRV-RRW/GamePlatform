package dbconn

import (
	"context"

	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisConnection struct {
	RedisClient *redis.Client
}

func NewRedisConnection(ctx context.Context, redis_host string, redis_port string) RedisConnection {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redis_host + ":" + redis_port,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := redisClient.Set(ctx, "test", "test", 0).Err()
	if err != nil {
		panic(err)
	}

	slog.Info("✅ Redis client connected successfully...")
	return RedisConnection{
		RedisClient: redisClient,
	}
}

func (c RedisConnection) Close() {
	c.RedisClient.Close()
}
