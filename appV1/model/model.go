package model

import (
	"time"
)

type Book struct {
	Id          int64  `json:"id" form:"id"`
	BN          string `json:"bn" form:"bn" gorm:"type:varchar(100)"`
	Name        string `json:"name" form:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" form:"description" gorm:"type:varchar(100)"`
	Count       int    `json:"count" form:"count"` //默认1
	CategoryId  int64  `json:"categoryId" form:"categoryId"`
}
type Category struct {
	Id   int64  `json:"id" form:"name"`
	Name string `json:"name" form:"name" gorm:"type:varchar(100)"`
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
	Id       int64  `json:"id" form:"id"`
	UserName string `json:"userName" form:"userName" gorm:"type:varchar(100)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(100)"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(100)"`
	Sex      string `json:"sex" form:"sex" gorm:"type:varchar(100)"`
	Phone    string `json:"phone" form:"phone" gorm:"type:varchar(100)"`
}
type Record struct {
	Id         int64     `json:"id" form:"id"`
	UserId     int64     `json:"userId" form:"userId"`
	BookId     int64     `json:"bookId" form:"bookId"`
	Status     int       `json:"status" form:"status"` //已归还1 未归还0
	StartTime  time.Time `json:"startTime" form:"startTime"`
	OverTime   time.Time `json:"overTime" form:"overTime"`
	ReturnTime time.Time `json:"returnTime" form:"returnTime"`
}
