package test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // 没有密码，默认值
	DB:       0,  // 默认DB 0
})

func TestRedisConnection(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Error(err)
	}
	value, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}
