package model

import (
	"context"
	"errors"
	"fmt"
	"libraryManagementSystem/appV2/tools"
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
func GetUserByPhone(phone string) User {
	var user User
	sql := "select * from users where phone=?"
	err := Conn.Raw(sql, phone).Scan(&user).Error
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
func SearchUser(userName, name string, limit, offset int) ([]User, int64) {
	var users []User
	sql1 := "select * from users where user_name like ? and name like ? "
	count := Conn.Raw(sql1, userName+"%", "%"+name+"%").Find(&users).RowsAffected
	sql2 := "select * from users where user_name like ? and name like ? limit ? offset ?"
	err := Conn.Raw(sql2, userName+"%", "%"+name+"%", limit, offset).Find(&users).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		return nil, 0
	}
	return users, count
}

func AddUser(user User) (uId int64, err error) {
	if user.Password == "" { //用户密码为空，自动生成
		user.Password = tools.GenPass()
	}
	sql := "insert into users(user_name,password,name,sex,phone,status) values (?,?,?,?,?,?)"
	err = Conn.Exec(sql, user.UserName, user.Password, user.Name, user.Sex, user.Phone, 0).Last(&user).Error
	if err != nil {
		mysqlLogger.Error(context.Background(), err.Error())
		return user.Id, err
	}
	return user.Id, nil
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
