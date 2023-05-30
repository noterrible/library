package model

import (
	"time"
)

type Book struct {
	Id          int64
	BN          string `gorm:"type:varchar(100)"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(100)"`
	Count       int    //默认1
	CategoryId  int64
}
type Category struct {
	Id   int64
	Name string `gorm:"type:varchar(100)"`
}
type User struct {
	Id       int64  `json:"id" form:"id"`
	UserName string `json:"userName" form:"userName" gorm:"type:varchar(100)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(100)"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(100)"`
	Sex      string `json:"sex" form:"sex" gorm:"type:varchar(100)"`
	Phone    string `json:"phone" form:"phone" gorm:"type:varchar(100)"`
	Status   int    `json:"status" form:"status"` //`json:""默认正常0 封禁1
}
type Librarian struct {
	Id       int64
	UserName string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	Name     string `gorm:"type:varchar(100)"`
	Sex      string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100)"`
}
type Record struct {
	Id         int64
	UserId     int64
	BookId     int64
	Status     int //已归还1 未归还0
	StartTime  time.Time
	OverTime   time.Time
	ReturnTime time.Time
}
