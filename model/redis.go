package model

import (
	"GoRedisClient/storage"
	"context"
)

var ctx = context.Background()

func ExistsNum(pattern string) (int64, error) {
	return storage.GetRedis().Exists(ctx, pattern).Result()
}
func keys(pattern string) ([]string, error) {
	return storage.GetRedis().Keys(ctx, pattern).Result()
}
func RDB() {
	storage.GetRedis().BgSave(ctx)
}
