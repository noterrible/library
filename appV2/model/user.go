package model

import (
	"context"
	"errors"
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
	sql := "select * from users where id=?"
	err := Conn.Raw(sql, id).Scan(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}
func SearchUser(userName, name string) []User {
	var users []User
	sql := "select * from users where user_name like ? and name like ?"
	err := Conn.Raw(sql, userName+"%", "%"+name+"%").Find(&users).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		return nil
	}
	return users
}

func AddUser(user User) {
	sql := "insert into users(user_name,password,name,sex,phone,status) values (?,?,?,?,?,?)"
	err := Conn.Exec(sql, user.UserName, user.Password, user.Name, user.Sex, user.Phone, 0).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
}
func UpdateUser(user User) error {
	sql := "update users set user_name=?,password=?,phone=? where id=?"
	rows := Conn.Exec(sql, user.UserName, user.Password, user.Phone, user.Id).RowsAffected
	if rows <= 0 {
		return errors.New("没有此用户")
	}
	return nil
}
func DeleteUser(id int64) {
	sql := "delete from users where id=?"
	err := Conn.Exec(sql, id).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
	}
}
