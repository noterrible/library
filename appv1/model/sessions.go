package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 创建cookie储存器
var store = sessions.NewCookieStore([]byte("1234secret"))
var sessionName = "admin"

// 获取session值
func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Println("session:", session.Values)
	return session.Values
}

// 设置session的值
func SetSession(c *gin.Context, id int64, name string) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["id"] = id
	session.Values["name"] = name
	return session.Save(c.Request, c.Writer)
}

// 清除session的值
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["id"] = 0
	session.Values["name"] = ""
	return session.Save(c.Request, c.Writer)
}
