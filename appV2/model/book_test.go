package model

import (
	"testing"
	"time"
)

func TestSearchBook(t *testing.T) {
	//q := "BN"

	//fmt.Println("搜索书籍"+q+"：", SearchBook(q, ""))
}

func TestAddBook(t *testing.T) {
	book := BookInfo{
		BookName:           "test",
		Author:             "testAuthor",
		PublishingHouse:    "V2",
		Translator:         "1234",
		PublishDate:        time.Now(),
		Pages:              100,
		ISBN:               "BN001",
		Price:              20,
		BriefIntroduction:  "test",
		AuthorIntroduction: "test",
		imgUrl:             "",
		delFlg:             0,
		Count:              100,
		CategoryId:         1,
	}
	AddBook(book)
}

func TestUpdateBook(t *testing.T) {
	book := BookInfo{
		BookName:           "test",
		Author:             "testAuthor1",
		PublishingHouse:    "V2.1",
		Translator:         "12341",
		PublishDate:        time.Now(),
		Pages:              100,
		ISBN:               "BN001",
		Price:              20,
		BriefIntroduction:  "test1",
		AuthorIntroduction: "test1",
		imgUrl:             "",
		delFlg:             0,
		Count:              90,
		CategoryId:         1,
	}
	UpdateBook(book)
}

//
//func TestTotalBooks(t *testing.T) {
//	fmt.Println(TotalBooks())
//}
//
//func TestPageBooks(t *testing.T) {
//	for _, v := range PageBooks(0, 10) {
//		fmt.Println(v.Id)
//
//	}
//}
