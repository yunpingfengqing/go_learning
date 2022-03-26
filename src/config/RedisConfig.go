package config

import (
	"blog/src/util"
	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func init() {
	util.ProtectRun(func() {
		Redisdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := Redisdb.Ping().Result()

		if err != nil {
			panic("failed to connect redis")
		}
	})
}
