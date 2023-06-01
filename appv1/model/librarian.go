package model

import "fmt"

func GetAdmin(userName, password string) Librarian {
	var admin Librarian
	sql := "select * from librarians where user_name=? and password=?"
	err := Conn.Raw(sql, userName, password).Scan(&admin).Error
	if err != nil {
		fmt.Println(err)
		return admin
	}
	if admin.Id > 0 {
		return admin
	}
	return admin
}
func AdminCheck(id int64, name string) bool {
	sql := "select * from librarians where id=? and user_name=?"
	var admin Librarian
	err := Conn.Raw(sql, id, name).First(&admin).Error
	if err != nil {
		return false
	}
	if admin.Id > 0 {
		return true
	}
	return false
}
