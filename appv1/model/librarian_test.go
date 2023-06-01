package model

import (
	"fmt"
	"testing"
)

func TestIsAdmin(t *testing.T) {
	fmt.Println("校验管理员：", GetAdmin("admin", "admin"))
}
