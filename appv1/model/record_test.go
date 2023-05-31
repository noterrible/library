package model

import (
	"fmt"
	"libraryManagementSystem/appv1/tools"
	"os"
	"testing"
	"time"
)

func TestGetAllRecordsByUserId(t *testing.T) {
	records := GetAllRecordsByUserId(1)
	fmt.Println(records)
}
func TestMain(m *testing.M) {
	New()
	m.Run()
	os.Exit(0)
}

func TestGetStatusRecordsByUserId(t *testing.T) {
	records := GetStatusRecordsByUserId(1, 1)
	fmt.Println(records)
}

func TestCreateRecord(t *testing.T) {
	record := Record{
		UserId:    1,
		BookId:    1,
		Status:    0,
		StartTime: time.Now(),
		OverTime:  time.Now().Add(tools.T),
	}
	CreateRecord(record)
}

func TestUpdateRecordAndBook(t *testing.T) {
	var id int64 = 2
	UpdateRecordAndBook(id)
}
