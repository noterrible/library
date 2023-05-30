package model

import "context"

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

func CreateRecord(record Record) {
	B := Conn.Begin()
	sql := "insert into records (user_id,book_id,status,start_time,over_time) values(?,?,?,?,?)"
	err := B.Exec(sql, record.UserId, record.BookId, record.Status, record.StartTime, record.OverTime).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		B.Rollback()
		return
	}
	sql = "update books set count=count-1 where id=?"
	err = Conn.Exec(sql, record.BookId).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		B.Rollback()
		return
	}
	B.Commit()
}
