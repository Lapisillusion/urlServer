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

	json2 = `[
{
"task_id":"11590",
"task_name":"热讯⻛采",
"task_slot":"am",
"task_status":"salacious",
"task_images":[
{
"image_id":1,
"image_url":"https://wm.cnnfootballclub.com/tzzq1.png",
"image_status":"audited"
},
{
"image_id":3,
"image_url":"https://wm.cnnfootballclub.com/tzzq2.png",
"image_status":"audited"
},
{
"image_id":4,
"image_url":"https://wm.cnnfootballclub.com/tzzq2.png",
"image_status":"salacious"
}
]
},
{
"task_id":"11590",
"task_name":"热讯⻛采",
"task_slot":"pm",
"task_status":"audited",
"task_images":[
{
"image_id":5,
"image_url":"https://wm.cnnfootballclub.com/tzzq1.png",
"image_status":"audited"
},
{
"image_id":6,
"image_url":"https://wm.cnnfootballclub.com/tzzq2.png",
"image_status":"audited"
},
{
"image_id":7,
"image_url":"https://wm.cnnfootballclub.com/tzzq2.png",
"image_status":"audited"
}
]
}
]`
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

func TestGetRecentTask(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	r := GetRecentTask(db, 7)
	fmt.Println(r)
}

func TestUpdateStatus(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	var list []*bean.TaskWithImg

	err := json.Unmarshal([]byte(json2), &list)

	if err != nil {
		log.Print("json format error")
		return
	}

	r := UpdateStatus(db, list)
	fmt.Println(r)
}
