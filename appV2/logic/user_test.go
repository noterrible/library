package logic

import (
	"context"
	"fmt"
	"libraryManagementSystem/appV2/model"
	"strconv"
	"testing"
)

func TestGetPhoneCode(t *testing.T) {
	countRedisConn := model.StopRestartRequestConn
	r, _ := countRedisConn.Get(context.Background(), "15537607006").Result()
	count, _ := strconv.Atoi(r)
	fmt.Printf("r.val:%v", count)
}
