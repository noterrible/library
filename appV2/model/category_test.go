package model

import (
	"fmt"
	"libraryManagementSystem/appV2/tools"
	"testing"
	"time"
)

func TestSearchCategory(t *testing.T) {
	q := "小"
	fmt.Println("搜索分类"+q+"：", SearchCategory(q))
}
func TestGzip(t *testing.T) {
	pageInfo := PageInfo{
		Page:  1,
		Limit: 3,
		Sort:  "",
	}
	var data ListResponse[BookInfo]
	err := Pages(&data, pageInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 100; i++ {
		go func() {
			dataToGzip, errTG := tools.ToGzip(data)
			if errTG != nil {
				fmt.Println(errTG.Error())
			} else {
				fmt.Println("压缩前数据", data.List[0].ISBN)
			}
			var dataNew ListResponse[BookInfo]
			errGT := tools.GzipTo(&dataNew, dataToGzip)
			if errGT != nil {
				fmt.Println(errGT.Error())
			} else {
				fmt.Println("压缩后解压数据", dataNew.List[0].ISBN)
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
func TestGetGzipTo(t *testing.T) {
	//key := fmt.Sprintf("books_%v_%v_%v_%v", currentPageString, pageSizeString, ISBN, bookName)

	//InfoCacheRedisConn.Get(context.Background(), key)

}
