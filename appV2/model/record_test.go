package model

import (
	"fmt"
	"libraryManagementSystem/appV2/tools"
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
	_, err := CreateRecord(record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestUpdateRecordAndBook(t *testing.T) {
	var id int64 = 2
	UpdateRecordAndBook(id)
}

func TestListeningTask(t *testing.T) {
	ListeningTask()

}
