package model

import (
	"errors"
	"fmt"
)

func SearchCategory(query string) []Category {
	var categories []Category
	sql := "select * from categories where name like ?"
	err := Conn.Raw(sql, "%"+query+"%").Find(&categories).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return categories
}
func GetCategory(id int64) Category {
	var category Category
	sql := "select * from categories where id=?"
	err := Conn.Raw(sql, id).Scan(&category).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return category
}
func AddCategory(category Category) {
	sql := "insert into categories(name) values (?)"
	err := Conn.Exec(sql, category.Name).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
func UpdateCategory(category Category) error {
	sql := "update categories set name=? where id=?"
	rows := Conn.Exec(sql, category.Name, category.Id).RowsAffected
	if rows <= 0 {
		return errors.New("没有此分类或者没有改动")
	}
	return nil
}
func DeleteCategory(id int64) error {
	sql := "delete from categories where id=?"
	rows := Conn.Exec(sql, id).RowsAffected
	if rows <= 0 {
		return errors.New("没有此分类")
	}
	return nil
}
