package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// 验证码
var RedisConn *redis.Client

// 防止用户重复发送请求，存用户id、访问路径
var StopRestartRequestConn *redis.Client

// 缓存查询的数据
var InfoCacheRedisConn *redis.Client

func init() {
	// 创建 Redis 客户端连接
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // Redis 未设置密码时为空
		DB:       0,  // 使用默认数据库0
	})
	// 测试连接是否成功
	_, err := RedisConn.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis for Captcha: %v", err)
		return
	}
	StopRestartRequestConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // Redis 未设置密码时为空
		DB:       1,  // 使用默认数据库1
	})

	// 测试连接是否成功
	_, err = StopRestartRequestConn.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis for StopRequest: %v", err)
		return
	}
	InfoCacheRedisConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // Redis 未设置密码时为空
		DB:       2,  // 使用默认数据库0
	})

	// 测试连接是否成功
	_, err = InfoCacheRedisConn.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis for InfoCache: %v", err)
		return
	}

	fmt.Println("Connected to Redis")
}
