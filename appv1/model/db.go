package model

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB
var mysqlLogger logger.Interface

func New() {
	mysqlLogger = logger.Default.LogMode(logger.Info)
	mysqlLogger.Info(context.Background(), "连接数据库···")
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "qq74263827", "127.0.0.1:3306", "library")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	Conn = conn
	Conn.AutoMigrate(&Book{}, &Category{}, &User{}, &Librarian{}, &Record{})
	MysqlLogger()
}

func MysqlLogger() {
	Conn = Conn.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
}
