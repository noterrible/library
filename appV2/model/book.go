package model

import (
	"errors"
	"fmt"
)

func GetBook(id int64) BookInfo {
	var book BookInfo
	sql := "select * from book_info where id=?"
	err := Conn.Raw(sql, id).Scan(&book).Error
	if err != nil {
		fmt.Println(err.Error())
		return book
	}
	return book
}
func SearchBook(ISBN, bookName string) []BookInfo {
	var books []BookInfo
	var books1 []BookInfo
	sql := "select * from book_info"
	err := Conn.Raw(sql).Scan(&books).Error
	if ISBN != "" || bookName != "" {
		sql = "select * from book_info where ISBN like ? and book_name like ?"
		err = Conn.Raw(sql, ISBN+"%", "%"+bookName+"%").Find(&books1).Error
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return books1
	}
	return books
}
func AddBook(book BookInfo) {
	sql := "insert into book_info(book_name,author,publishing_house,translator,publish_date,pages,ISBN,price,brief_description,author_description,img_url,count,category_id) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	err := Conn.Exec(sql, book.BookName, book.Author, book.PublishingHouse, book.Translator, book.PublishDate, book.Pages, book.ISBN, book.Price, book.BriefIntroduction, book.AuthorIntroduction, book.imgUrl, book.Count, book.CategoryId).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
func UpdateBook(book BookInfo) error {
	sql := "update book_info set book_name=?,author=?,publishing_house=?,translator=?,publish_date=?,pages=?,ISBN=?,price=?,brief_description=?,author_description=?,img_url=?,count=?,category_id=? where id=?"
	rows := Conn.Exec(sql, book.BookName, book.Author, book.PublishingHouse, book.Translator, book.PublishDate, book.Pages, book.ISBN, book.Price, book.BriefIntroduction, book.AuthorIntroduction, book.imgUrl, book.delFlg, book.delFlg, book.Count, book.CategoryId).RowsAffected
	if rows <= 0 {
		fmt.Println(errors.New("记录不存在或者没有改动"))
		return errors.New("记录不存在或者没有改动")
	}
	return nil
}
func DeleteBook(id int64) error {
	sql := "delete from book_info where id=?"
	count := Conn.Exec(sql, id).RowsAffected
	if count == 0 {
		return errors.New("记录不存在")
	}
	return nil
}
