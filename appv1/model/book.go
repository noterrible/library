package model

import "fmt"

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
func UpdateBook(book Book) {
	sql := "update books set bn=?,name=?,description=?,count=?,category_id=? where id=?"
	err := Conn.Exec(sql, book.BN, book.Name, book.Description, book.Count, book.CategoryId, book.Id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
func DeleteBook(id int64) {
	sql := "delete from books where id=?"
	err := Conn.Exec(sql, id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
