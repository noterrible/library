package model

import (
	"errors"
	"fmt"
)

func GetAllRecordsByUserId(id int64) []Record {
	var records []Record
	sql := "select * from records where id=? order by start_time"
	err := Conn.Exec(sql, id).Find(&records).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return records
}
func GetStatusRecordsByUserId(userId int64, status int) []Record {
	var records []Record
	sql := "select * from records where user_id=? and status=? order by start_time"
	err := Conn.Raw(sql, userId, status).Scan(&records).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return records
}
func GetStatusRecords(status int) []Record {
	var records []Record
	sql := "select * from records where status=? order by start_time"
	err := Conn.Raw(sql, status).Scan(&records).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return records
}

func CreateRecord(record Record) (int64, error) {
	var book Book
	sql := "SELECT count FROM books WHERE id = ?"
	rows := Conn.Raw(sql, record.BookId).Scan(&book).RowsAffected
	if rows == 0 {
		return 0, errors.New("没有此书")
	}
	B := Conn.Begin()
	if book.Count <= 0 {
		B.Rollback()
		return 0, errors.New("书籍数量不够")
	}
	sql = "insert into records (user_id,book_id,status,start_time,over_time) values(?,?,?,?,?)"
	err := B.Exec(sql, record.UserId, record.BookId, record.Status, record.StartTime, record.OverTime).Last(&record).Error
	fmt.Println("r.Id:", record.Id)
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}

	sql = "update books set count=count-1 where id=?"
	err = Conn.Exec(sql, record.BookId).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	B.Commit()
	return record.Id, nil
}
func UpdateRecordAndBook(id int64) (int64, error) {
	sql := "select * from records where id=?"
	var record Record
	rows := Conn.Raw(sql, id).Scan(&record).RowsAffected
	if rows <= 0 {
		return 0, errors.New("你没有关于这本书的借书记录")
	}
	if record.Status == 1 {
		return 0, errors.New("你已经还过书了")
	}
	B := Conn.Begin()
	//更新记录
	sql = "update records set status=1 where id=?"
	err := B.Exec(sql, id).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	//更新书籍数量
	sql = "update books set count=count+1 where id=?"
	err = B.Exec(sql, record.BookId).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	B.Commit()
	return record.Id, nil
}
