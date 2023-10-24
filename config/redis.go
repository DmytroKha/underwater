package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var Redis *redis.Client

func CreateRedisClient(ctx context.Context, conf Configuration) {

	RedisDatabase, err := strconv.Atoi(conf.RedisDatabase)
	if err != nil {
		log.Println(err)
	}

	rds := redis.NewClient(&redis.Options{
		Addr: conf.RedisHost,
		//Password: conf.RedisPassword,
		DB: RedisDatabase,
	})

	_, err = rds.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	Redis = rds
}
