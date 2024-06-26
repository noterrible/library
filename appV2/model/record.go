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
	var book BookInfo
	B := Conn.Begin()
	sql := "SELECT * FROM book_info WHERE id = ? for update"
	B.Raw(sql, record.BookId).Scan(&book)
	if book.Id <= 0 {
		B.Rollback()
		return 0, errors.New("没有此书")
	}
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

	sql = "update book_info set count=count-1 where id=?"
	err = B.Exec(sql, record.BookId).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	B.Commit()
	return record.Id, nil
}
func UpdateRecordAndBook(id int64) (int64, error) {
	B := Conn.Begin()
	sql := "select * from records where id=? for update"
	var record Record
	B.Raw(sql, id).Scan(&record)
	if record.Id <= 0 {
		B.Rollback()
		return 0, errors.New("你没有关于这本书的借书记录")
	}
	if record.Status == 1 {
		B.Rollback()
		return 0, errors.New("你已经还过书了")
	}
	//更新记录
	sql = "update records set status=1 where id=?"
	err := B.Exec(sql, id).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	//更新书籍数量
	sql = "update book_info set count=count+1 where id=?"
	err = B.Exec(sql, record.BookId).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0, err
	}
	//还书更新消息系统，把未读的还书提示状态置为已读
	//查有没有要还的书
	sql = "select * from records where over_time >= DATE_SUB(NOW(), INTERVAL 1 DAY) and status=0 and user_id"
	var records []Record
	B.Raw(sql, record.UserId).Scan(&records)
	//没有要还的书，消息则置为已读
	if len(records) == 0 {
		sql = "update messages set status==1 where user_id=?"
		err = B.Exec(sql, record.UserId).Error
		if err != nil {
			fmt.Println(err.Error())
			B.Rollback()
			return 0, err
		}
	}
	B.Commit()
	return record.Id, nil
}
