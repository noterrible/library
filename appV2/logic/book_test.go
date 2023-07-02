package logic

import (
	"fmt"
	"libraryManagementSystem/appV2/model"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestPageBook(t *testing.T) {
	//查第一页数据
	//books := model.PageBooks(0, 10)
	////查出所有数据后分页函数进行分页
	//currentPage, _ := strconv.Atoi("1")
	//total := model.TotalBooks()
	//page := tools.NewPage[model.BookInfo](currentPage, 10, int(total), books)
	//fmt.Println(page)
}
func TestMain(m *testing.M) {
	model.New()
	m.Run()
	os.Exit(0)
}
func TestRandInt(t *testing.T) {
	for i := 0; i < 50; i++ {
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(43) + 20
		fmt.Println(randInt)
	}

}
