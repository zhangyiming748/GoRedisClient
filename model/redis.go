package model

import (
	"GoRedisClient/storage"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func RDB() {
	storage.GetRedis().BgSave(ctx)
}

/*
时间段过期
*/
func Expire(key string, t time.Duration) (bool, error) {
	return storage.GetRedis().Expire(ctx, key, t).Result()
}

/*
时间点过期
*/
func ExpireAt(key string, t time.Time) (bool, error) {
	return storage.GetRedis().ExpireAt(ctx, key, t).Result()
}

func Exists(pattern string) (int64, error) {
	return storage.GetRedis().Exists(ctx, pattern).Result()
}

func keys(pattern string) ([]string, error) {
	return storage.GetRedis().Keys(ctx, pattern).Result()
}

func MGet(key []string) (interface{}, error) {
	return storage.GetRedis().MGet(ctx, key...).Result()
}

func Get(key string) (string, error) {
	return storage.GetRedis().Get(ctx, key).Result()
}

func Set(key, value string) (string, error) {
	return storage.GetRedis().Set(ctx, key, value, 0).Result()
}

func HMGet(key string, fields []string) (interface{}, error) {
	return storage.GetRedis().HMGet(ctx, key, fields...).Result()
}

func HGetAll(key string) (map[string]string, error) {
	return storage.GetRedis().HGetAll(ctx, key).Result()
}

func LRange(key string, start, stop int64) ([]string, error) {
	return storage.GetRedis().LRange(ctx, key, start, stop).Result()
}

/*
返回元素在集合中的升序排名情况
ZRevRank()倒序
*/
func ZRank(key, field string) (int64, error) {
	return storage.GetRedis().ZRank(ctx, key, field).Result()
}

/*
获取元素的score
*/
func ZScore(key, field string) (float64, error) {
	return storage.GetRedis().ZScore(ctx, key, field).Result()
}

/*
根据分值排序后的，升序的列表获取
降序ZRevRange()
*/
func ZRange(key string, min, max int64) ([]string, error) {
	return storage.GetRedis().ZRange(ctx, key, min, max).Result()
}

func ZAdd(key string, score float64, fields []interface{}) (int64, error) {
	z := new(redis.Z)
	z.Score = score
	z.Member = fields
	return storage.GetRedis().ZAdd(ctx, key, *z).Result()
}

/*
获取链表元素个数
*/
func LLen(name string) (int64, error) {
	return storage.GetRedis().LLen(ctx, name).Result()
}

func RPop(name string) (string, error) {
	return storage.GetRedis().RPop(ctx, name).Result()
}

func LPush(name string, val interface{}) (int64, error) {
	return storage.GetRedis().LPush(ctx, name, val).Result()
}

func RPush(name string, val interface{}) (int64, error) {
	return storage.GetRedis().RPush(ctx, name, val).Result()
}

func HIncrBy(key, field string, incr int64) (int64, error) {
	return storage.GetRedis().HIncrBy(ctx, key, field, incr).Result()
}

func HMSet(key string, fields interface{}) (interface{}, error) {
	return storage.GetRedis().HMSet(ctx, key, fields).Result()
}
