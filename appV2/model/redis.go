package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client
var StopRestartRequestConn *redis.Client

func init() {
	// 创建 Redis 客户端连接
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // Redis 未设置密码时为空
		DB:       0,  // 使用默认数据库0
	})
	StopRestartRequestConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // Redis 未设置密码时为空
		DB:       1,  // 使用默认数据库1
	})
	// 测试连接是否成功
	_, err := RedisConn.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v", err)
		return
	}
	fmt.Println("Connected to Redis")
}
