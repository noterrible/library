package model

import (
	"fmt"
	"github.com/gin-contrib/sessions/redis"
	"strconv"

	//"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//如何设置key到redis
//sessionid是怎么加密的
//

// 创建cookie储存器
var store, _ = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
var sessionName = "admin"

// 获取session值
func GetSession(c *gin.Context) map[interface{}]interface{} {
	idStr, _ := c.Cookie("admin_id")
	fmt.Println("admin_id:", idStr)
	session, _ := store.Get(c.Request, sessionName+idStr)
	fmt.Println("session:", session.Values)
	return session.Values
}

// 设置session的值
func SetSession(c *gin.Context, id int64, name string) error {
	idStr := strconv.FormatInt(id, 10)
	fmt.Println("admin_id:", idStr)
	session, _ := store.Get(c.Request, sessionName+idStr)
	session.Options.MaxAge = 86400
	session.Values["id"] = id
	session.Values["name"] = name
	return session.Save(c.Request, c.Writer)
}

// 清除session的值
func FlushSession(c *gin.Context) error {
	idStr, _ := c.Cookie("id")
	session, _ := store.Get(c.Request, sessionName+idStr)
	session.Values["id"] = int64(0)
	session.Values["name"] = ""
	return session.Save(c.Request, c.Writer)
}
