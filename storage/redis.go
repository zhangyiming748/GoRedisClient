package storage

import (
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
func GetRedis() *redis.Client {
	return RedisClient
}
