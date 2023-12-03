package storage

import (
	"context"
	"fmt"
	"testing"
)

func TestMaster(t *testing.T) {
	SetRedis()
	ctx := context.Background()
	GetRedis().Set(ctx, "name", "zhen", 0)
	result, _ := GetRedis().Get(ctx, "name").Result()
	fmt.Println(result)

}
