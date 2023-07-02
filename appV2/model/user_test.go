package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := User{UserName: "test", Phone: "13424563463"}
	id, err := AddUser(user)
	fmt.Println("id", id, "err", err)
}
