package model

import (
	"errors"
	"fmt"
)

func GetBook(id int64) Book {
	var book Book
	sql := "select * from books where id=?"
	err := Conn.Raw(sql, id).Scan(&book).Error
	if err != nil {
		fmt.Println(err.Error())
		return book
	}
	return book
}
func SearchBook(query string) []Book {
	var books []Book
	var books1 []Book
	sql := "select * from books"
	err := Conn.Raw(sql).Scan(&books).Error
	if query != "" {
		sql = "select * from books where bn like ? or name like ?"
		err = Conn.Raw(sql, query+"%", "%"+query+"%").Find(&books1).Error
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return books1
	}
	return books
}
func AddBook(book Book) {
	sql := "insert into books(bn,name,description,count,category_id) values (?,?,?,?,?)"
	err := Conn.Exec(sql, book.BN, book.Name, book.Description, book.Count, book.CategoryId).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
func UpdateBook(book Book) error {
	sql := "update books set bn=?,name=?,description=?,count=?,category_id=? where id=?"
	rows := Conn.Exec(sql, book.BN, book.Name, book.Description, book.Count, book.CategoryId, book.Id).RowsAffected
	if rows <= 0 {
		fmt.Println(errors.New("记录不存在或者没有改动"))
		return errors.New("记录不存在或者没有改动")
	}
	return nil
}
func DeleteBook(id int64) error {
	sql := "delete from books where id=?"
	count := Conn.Exec(sql, id).RowsAffected
	if count == 0 {
		return errors.New("记录不存在")
	}
	return nil
}
