package model

import (
	"context"
	"fmt"
)

func GetAllRecordsByUserId(id int64) []Record {
	var records []Record
	sql := "select * from records where id=? order by start_time"
	err := Conn.Exec(sql, id).Find(&records).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
	return records
}
func GetStatusRecordsByUserId(id int64, status int) []Record {
	var records []Record
	sql := "select * from records where id=? and status=? order by start_time"
	err := Conn.Raw(sql, id, status).Scan(&records).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
	return records
}
func GetStatusRecords(status int) []Record {
	var records []Record
	sql := "select * from records where status=? order by start_time"
	err := Conn.Raw(sql, status).Scan(&records).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
	return records
}

func CreateRecord(record Record) int64 {
	B := Conn.Begin()
	sql := "insert into records (user_id,book_id,status,start_time,over_time) values(?,?,?,?,?)"
	err := B.Exec(sql, record.UserId, record.BookId, record.Status, record.StartTime, record.OverTime).Last(&record).Error
	fmt.Println("r.Id:", record.Id)
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		B.Rollback()
		return 0
	}
	sql = "update books set count=count-1 where id=?"
	err = Conn.Exec(sql, record.BookId).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		B.Rollback()
		return 0
	}
	B.Commit()
	return record.Id
}
func UpdateRecordAndBook(id int64) int64 {
	sql := "select * from records where id=?"
	var record Record
	err := Conn.Raw(sql, id).Scan(&record).Error
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	record.Status = 1
	B := Conn.Begin()
	//更新记录
	sql = "update records set status=1 where id=?"
	err = B.Exec(sql, id).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0
	}
	//更新书籍数量
	sql = "update books set count=count+1 where id=?"
	err = B.Exec(sql, record.BookId).Error
	if err != nil {
		fmt.Println(err.Error())
		B.Rollback()
		return 0
	}
	B.Commit()
	return record.Id
}
