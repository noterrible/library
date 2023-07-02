package model

import (
	"time"
)

type BookInfo struct {
	Id                 int64     `json:"id" form:"id"`
	BookName           string    `json:"bookName" form:"bookName"`
	Author             string    `json:"author" form:"author"`
	PublishingHouse    string    `json:"publishingHouse" form:"publishingHouse"`
	Translator         string    `json:"translator" form:"translator"`
	PublishDate        time.Time `json:"publishDate" form:"publishDate"`
	Pages              int       `json:"pages" form:"pages"`
	ISBN               string    `json:"ISBN" form:"ISBN"`
	Price              float64   `json:"price" form:"price"`
	BriefIntroduction  string    `json:"briefIntroduction" form:"briefIntroduction"`
	AuthorIntroduction string    `json:"authorIntroduction" form:"authorIntroduction"`
	imgUrl             string    `json:"imgUrl" form:"imgUrl"`
	delFlg             int       `json:"delFlg" form:"delFlg"` //默认0
	Count              int       `json:"count" form:"count"`
	CategoryId         int64     `json:"categoryId" form:"categoryId"`
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
type Message struct {
	Id         int64     `json:"id" form:"id"`
	UserId     int64     `json:"userId" form:"userId"`
	Message    string    `json:"message" form:"message" gorm:"type:varchar(100)"`
	Status     int       `json:"status" form:"status"` //0未读 1已读
	CreateTime time.Time `json:"createTime" form:"createTime"`
}
