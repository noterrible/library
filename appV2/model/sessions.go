package model

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"time"
)

// 创建cookie储存器
var store, _ = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
var sessionName = "admin"

// 获取session值
func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Println("session:", session.Values)
	return session.Values
}

// 设置session的值
func SetSession(c *gin.Context, id int64, name string) error {
	store.Options(sessions.Options{MaxAge: int(24 * time.Hour)})
	session, _ := store.Get(c.Request, sessionName)
	session.Values["id"] = id
	session.Values["name"] = name
	return session.Save(c.Request, c.Writer)
}

// 清除session的值
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["id"] = int64(0)
	session.Values["name"] = ""
	return session.Save(c.Request, c.Writer)
}
