package model

import (
	"fmt"
	"time"
)

func ListeningTask() {
	//查出所有距离截止时间一天的所有记录
	sql := "select distinct user_id from records where over_time >= DATE_SUB(NOW(), INTERVAL 1 DAY) and status=0"
	var userIds []int64
	Conn.Raw(sql).Scan(&userIds)
	fmt.Println("需要还书：", userIds)
	if len(userIds) == 0 {
		return
	}
	//生成所有未还书的用户消息
	for _, userId := range userIds {
		B := Conn.Begin()
		//再次检测该用户是否已还
		sql = "select * from records where status=0 and user_id=? for update"
		var users []User
		B.Raw(sql, userId).Scan(&users)
		//检测该用户最近30天是否已生成过未读消息
		sql = "select * from messages where status=0 and user_id=? and create_time>= DATE_SUB(CURDATE(), INTERVAL 30 DAY)"
		var messages []Message
		B.Raw(sql, userId).Scan(&messages)
		if len(users) != 0 && len(messages) == 0 {
			sql = "insert into messages(user_id,message,status,create_time) values(?,?,?,?)"
			err := B.Exec(sql, userId, "您有书需要归还", 0, time.Now()).Error
			if err != nil {
				fmt.Println(err.Error())
				B.Rollback()
				return
			}
		}
		B.Commit()
	}
}
