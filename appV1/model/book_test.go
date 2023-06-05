package model

import (
	"fmt"
	"testing"
)

func TestSearchBook(t *testing.T) {
	q := "BN"
	fmt.Println("搜索书籍"+q+"：", SearchBook(q))
}

func TestAddBook(t *testing.T) {
	book := Book{
		BN:          "BN00002",
		Name:        "三国演义",
		Description: "讲述了。。",
		Count:       100,
		CategoryId:  1,
	}
	AddBook(book)
}

func TestUpdateBook(t *testing.T) {
	book := Book{
		Id:          2,
		BN:          "BN00002",
		Name:        "三国演义",
		Description: "讲述了。。",
		Count:       101,
		CategoryId:  1,
	}
	UpdateBook(book)
}
