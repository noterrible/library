package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client

func init() {
	// 创建 Redis 客户端连接
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "qq74263827", // Redis 未设置密码时为空
		DB:       1,            // 使用默认数据库
	})
	// 测试连接是否成功
	_, err := RedisConn.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v", err)
		return
	}
	fmt.Println("Connected to Redis")
}
