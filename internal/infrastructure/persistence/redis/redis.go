package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func CoonectToRedis() {

	dsn, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		// do something later
		fmt.Println("Failed to env")
	}

	var dbIndex int = 1
	cachedb := os.Getenv("REDIS_INDEX_ADMIN")
	if cachedb == "" {
		panic("REDIS_INDEX_ADMIN is not set")
	} else {
		// convert cachedb to int
		dbIndex, err = strconv.Atoi(cachedb)
		if err != nil {
			panic("REDIS_INDEX_ADMIN is not set between 1-15")
		}

	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbIndex,
	})

	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}

	RedisClient = rdb

}
