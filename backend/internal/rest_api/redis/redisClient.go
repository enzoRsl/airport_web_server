package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var ctx = context.Background()

func GetRedisClient() (*redis.Client, error) {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       db,                      // use default DB
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		println(err.Error())
		return nil, err
	}
	return rdb, nil
}
