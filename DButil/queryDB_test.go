package DButil

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"urlServer/bean"
	"urlServer/initconfig"
)

const (
	json1 = `{
			"task_id":"11590",
			"task_name":"热讯⻛采",
			"task_time":1690791332,
			"task_images":[
			"https://wm.cnnfootballclub.com/tzzq1.png",
			"https://wm.cnnfootballclub.com/tzzq2.jpg",
			"https://wm.cnnfootballclub.com/tzzq3.png"
			]
}`
)

func TestInsert(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	var r bean.Report
	err := json.Unmarshal([]byte(json1), &r)
	if err != nil {
		log.Fatal("json format error")
		return
	}
	ReportTasks(db, &r)
}

func TestGetTodayData(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	r := GetTodayData(db)
	fmt.Println(r)
}
