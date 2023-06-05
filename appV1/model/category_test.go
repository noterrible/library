package model

import (
	"fmt"
	"testing"
)

func TestSearchCategory(t *testing.T) {
	q := "小"
	fmt.Println("搜索分类"+q+"：", SearchCategory(q))
}
