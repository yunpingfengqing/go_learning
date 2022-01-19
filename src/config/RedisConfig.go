package config

import (
	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,  // use default DB
	})

	_, err := Redisdb.Ping().Result()

	if err != nil {
		panic("failed to connect redis")
	}
}
