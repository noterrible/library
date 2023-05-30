package model

import (
	"context"
	"fmt"
)

func UserCheck(userName, password string) User {
	var user User
	sql := "select * from users where  user_name=? and password=?"
	err := Conn.Raw(sql, userName, password).Scan(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}
func GetUser(id int64) User {
	var user User
	sql := "select * from users where  id=?"
	err := Conn.Raw(sql, id).Scan(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}
func SearchUser(query string) []User {
	var users []User
	var users1 []User
	sql := "select * from users"
	err := Conn.Raw(sql).Scan(&users).Error
	if query != "" {
		sql = "select * from users where user_name like ? or name like ?"
		err = Conn.Raw(sql, query+"%", "%"+query+"%").Find(&users1).Error
		if err != nil {
			mysqlLogger.Error(context.Background(), err.Error())
			return nil
		}
		return users1
	}
	return users
}

func AddUser(user User) {
	sql := "insert into users(user_name,password,name,sex,phone) values (?,?,?,?,?)"
	err := Conn.Exec(sql, user.UserName, user.Password, user.Name, user.Sex, user.Phone).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
}
func UpdateUser(user User) {
	sql := "update users set user_name=? and password=? and phone=? where id=?"
	err := Conn.Exec(sql, user.UserName, user.Password, user.Phone, user.Id).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
}
func DeleteUser(id int64) {
	sql := "delete from user where id=?"
	err := Conn.Exec(sql, id).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
}
