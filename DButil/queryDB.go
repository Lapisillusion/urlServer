package DButil

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
	"urlServer/bean"
)

const (
	MAXIMGS          = 5
	TIMEFORMAT       = "2006-01-02"
	inserttask       = `insert into tasks values (null,"%s","%s","unaudited","%s","%s")`
	insertimg        = `insert into images values `
	selectimgs       = `select image_id,image_url,image_status from images where id = ?`
	selecttodaytask  = `select id,task_id,task_name,task_status,task_slot from tasks where date_time=?`
	selectrecenttask = `select task_id,task_name,task_status,date_time from tasks where DATEDIFF(?,date_time)<? AND task_status = "salacious"`
	updatetasksql    = `update tasks set task_status = ? where task_id=? and task_slot=? and date_time=?`
	updateimgsql     = `update images set image_status = "%s" where image_id in (`
)

// 任务5
func UpdateStatus(db *sql.DB, list []*bean.TaskWithImg) bool {

	ch := make(chan struct{}, 12)
	var wg sync.WaitGroup

	for _, taskimg := range list {
		wg.Add(1)
		ch <- struct{}{}
		go updateImg(db, taskimg.TaskImages, ch, &wg)
	}
	s := time.Now().Format(TIMEFORMAT)

	handleTaskList(db, list, s)
	wg.Wait()
	close(ch)
	return true
}

func updateImg(db *sql.DB, images []*bean.Img, ch chan struct{}, wg *sync.WaitGroup) {
	audlist := make([]*bean.Img, 0, 8)
	sallist := make([]*bean.Img, 0, 8)

	for _, img := range images {
		switch img.ImageStatus {
		case "audited":
			audlist = append(audlist, img)
		case "salacious":
			sallist = append(sallist, img)
		}
	}

	handleImgList(db, audlist, "audited")
	handleImgList(db, sallist, "salacious")
	<-ch
	wg.Done()
}

func handleTaskList(db *sql.DB, list []*bean.TaskWithImg, s string) {
	if len(list) == 0 {
		return
	}
	for _, task := range list {
		_, err := db.Exec(updatetasksql, task.TaskStatus, task.TaskId, task.TaskSlot, s)
		if err != nil {
			log.Println("handleTaskList error")
			return
		}
	}
}

func handleImgList(db *sql.DB, list []*bean.Img, s string) {
	if len(list) == 0 {
		return
	}
	var bs bytes.Buffer
	bs.WriteString(fmt.Sprintf(updateimgsql, s))
	for _, img := range list {
		bs.WriteString(fmt.Sprintf("%d,", img.ImageId))
	}
	str := bs.String()
	str = str[:len(str)-1]
	_, err := db.Exec(str + ")")
	if err != nil {
		log.Println("handleImgList error", err)
		return
	}
}

// 任务4
func GetRecentTask(db *sql.DB, diff int) *bean.TaskList {

	rows, err := db.Query(selectrecenttask, time.Now().Format(TIMEFORMAT), diff)
	defer rows.Close()
	if err != nil {
		return nil
	}
	var list bean.TaskList
	list.Code = 200
	list.Message = "ok"
	list.Data = make([]*bean.DateWithTask, 0, 16)
	taskMap := make(map[string][]*bean.Task)
	for rows.Next() {
		var tstr string
		var task bean.Task
		rows.Scan(&task.TaskId, &task.TaskName, &task.TaskStatus, &tstr)
		_, ok := taskMap[tstr]
		if !ok {
			taskMap[tstr] = make([]*bean.Task, 0, 8)
		}
		taskMap[tstr] = append(taskMap[tstr], &task)
	}

	for k, v := range taskMap {
		var date bean.DateWithTask
		date.DateTime = k
		date.TaskList = make([]*bean.Task, 0, 4)
		for _, t := range v {
			date.TaskList = append(date.TaskList, t)
		}
		list.Data = append(list.Data, &date)
	}

	return &list
}

// 任务3
func GetTodayData(db *sql.DB) *bean.TodayData {

	t := time.Now().Format(TIMEFORMAT)

	rows, err := db.Query(selecttodaytask, t)

	defer rows.Close()

	if err != nil {
		log.Println("query today's data failed")
		return nil
	}

	var today bean.TodayData
	today.Code = 200
	today.Message = "ok"
	today.Data = make([]*bean.TaskWithImg, 0, 16)
	ch := make(chan bool, 12)
	var wg sync.WaitGroup

	for rows.Next() {
		var taskimgs bean.TaskWithImg
		taskimgs.TaskImages = make([]*bean.Img, 0, 16)
		var id int
		rows.Scan(&id, &taskimgs.TaskId, &taskimgs.TaskName, &taskimgs.TaskSlot, &taskimgs.TaskStatus)
		today.Data = append(today.Data, &taskimgs)

		wg.Add(1)
		ch <- true
		go queryImg(db, &taskimgs, id, ch, &wg)
	}

	wg.Wait()
	close(ch)
	return &today
}

func queryImg(db *sql.DB, taskimgs *bean.TaskWithImg, id int, ch <-chan bool, wg *sync.WaitGroup) {
	defer func() {
		<-ch
		wg.Done()
	}()

	rows, err := db.Query(selectimgs, id)

	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var img bean.Img
		rows.Scan(&img.ImageId, &img.ImageUrl, &img.ImageStatus)
		taskimgs.TaskImages = append(taskimgs.TaskImages, &img)
	}
}

func getslot(h int) (string, error) {
	if h < 14 && h >= 0 {
		return "am", nil
	} else if h >= 14 && h <= 23 {
		return "pm", nil
	} else {
		return "", fmt.Errorf("wrong time format")
	}
}

// 任务2
func ReportTasks(db *sql.DB, report *bean.Report) bool {
	t := time.Unix(int64(report.TaskTime), 0)
	slot, err := getslot(t.Hour())
	if err != nil {
		log.Print(err)
		return false
	}

	result, err := db.Exec(fmt.Sprintf(inserttask, report.TaskId,
		report.TaskName,
		slot,
		t.Format(TIMEFORMAT)))

	if err != nil {
		log.Printf("insert %s task failed", report.TaskId)
		return false
	}

	imglen := len(report.TaskImages)
	if imglen == 0 {
		return true
	}
	tid, _ := result.LastInsertId()
	var i = 0
	for i = 0; i+MAXIMGS <= imglen; i += MAXIMGS {
		err := insertimgs(db, report.TaskImages, i, i+MAXIMGS, tid)
		if err != nil {
			log.Println("something wrong append when insert img url")
			continue
		}
	}
	if imglen%MAXIMGS != 0 {
		err := insertimgs(db, report.TaskImages, (imglen/MAXIMGS)*MAXIMGS, imglen, tid)
		if err != nil {
			log.Println("something wrong append when insert img url")
		}
	}
	return true
}

func insertimgs(db *sql.DB, imgs []string, l int, r int, tid int64) error {
	bs := bytes.Buffer{}
	bs.WriteString(insertimg)
	for i := l; i < r; i++ {
		bs.WriteString(fmt.Sprintf(`(null,"%s","unaudited","%d"),`, imgs[i], tid))
	}
	sqlstr := bs.String()
	_, err := db.Exec(sqlstr[:len(sqlstr)-1])
	return err
}
