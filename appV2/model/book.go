package model

import (
	"errors"
	"fmt"
)

func GetBook(id int64) BookInfo {
	var book BookInfo
	sql := "select * from book_infos where id=?"
	err := Conn.Raw(sql, id).Scan(&book).Error
	if err != nil {
		fmt.Println(err.Error())
		return book
	}
	return book
}

func SearchBook(ISBN, bookName string, limit, offset int) ([]BookInfo, int64) {
	var books []BookInfo
	sql1 := "select * from book_infos where ISBN like ? and book_name like ?"
	count := Conn.Raw(sql1, ISBN+"%", "%"+bookName+"%").Find(&books).RowsAffected
	sql2 := "select * from book_infos where ISBN like ? and book_name like ? limit ? offset ?"
	err := Conn.Raw(sql2, ISBN+"%", "%"+bookName+"%", limit, offset).Find(&books).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}
	return books, count
}

func AddBook(book BookInfo) error {
	sql := "insert into book_infos(book_name,author,publishing_house,translator,publish_date,pages,ISBN,price,brief_introduction,author_introduction,img_url,count,category_id) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	err := Conn.Exec(sql, book.BookName, book.Author, book.PublishingHouse, book.Translator, book.PublishDate, book.Pages, book.ISBN, book.Price, book.BriefIntroduction, book.AuthorIntroduction, book.imgUrl, book.Count, book.CategoryId).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
func UpdateBook(book BookInfo) error {
	sql := "update book_infos set book_name=?,author=?,publishing_house=?,translator=?,publish_date=?,pages=?,ISBN=?,price=?,brief_description=?,author_description=?,img_url=?,count=?,category_id=? where id=?"
	rows := Conn.Exec(sql, book.BookName, book.Author, book.PublishingHouse, book.Translator, book.PublishDate, book.Pages, book.ISBN, book.Price, book.BriefIntroduction, book.AuthorIntroduction, book.imgUrl, book.delFlg, book.delFlg, book.Count, book.CategoryId).RowsAffected
	if rows <= 0 {
		fmt.Println(errors.New("记录不存在或者没有改动"))
		return errors.New("记录不存在或者没有改动")
	}
	return nil
}
func DeleteBook(id int64) error {
	sql := "delete from book_infos where id=?"
	count := Conn.Exec(sql, id).RowsAffected
	if count == 0 {
		return errors.New("记录不存在")
	}
	return nil
}
