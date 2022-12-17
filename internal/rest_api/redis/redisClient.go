package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return rdb, nil
}
