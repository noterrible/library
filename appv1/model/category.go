package model

import "fmt"

func SearchCategory(query string) []Category {
	var categories []Category
	var categories1 []Category
	sql := "select * from categories"
	err := Conn.Raw(sql).Scan(&categories).Error
	if query != "" {
		sql = "select * from categories where name like ?"
		err = Conn.Raw(sql, "%"+query+"%").Find(&categories1).Error
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return categories1
	}
	return categories
}
func GetCategory(id int64) Category {
	var category Category
	sql := "select * from users where id=?"
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
func UpdateCategory(category Category) {
	sql := "update categories set name=? where id=?"
	err := Conn.Exec(sql, category.Name, category.Id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
func DeleteCategory(id int64) {
	sql := "delete from categories where id=?"
	err := Conn.Exec(sql, id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
