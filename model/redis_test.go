package model

import (
	"GoRedisClient/storage"
	"testing"
)

func init() {
	storage.SetRedis()
}
func TestExists(t *testing.T) {
	exists, err := ExistsNum("sh")
	if err != nil {
		println(err)
	}
	t.Logf("存在%d个值\n", exists)
}
func TestKeys(t *testing.T) {
	val, err := keys("*sh")
	if err != nil {
		println(err)
	}
	for i, v := range val {
		t.Logf("找到第%d个值:%s\n", i+1, v)
	}
}
