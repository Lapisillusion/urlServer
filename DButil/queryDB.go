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
	MAXIMGS         = 5
	TIMEFORMAT      = "2006-01-02"
	inserttask      = `insert into tasks values (null,"%s","%s","unaudited","%s","%s")`
	insertimg       = `insert into images values `
	selectimgs      = `select image_id,image_url,image_status from images where id = ?`
	selecttodaytask = `select id,task_id,task_name,task_status,task_slot from tasks where data_time=?`
)

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

	ch := make(chan bool, 12)
	var wg sync.WaitGroup

	for rows.Next() {
		var taskimgs bean.TaskWithImg
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

const Universal_ok = `{
  "code": 200,
  "message": "ok",
  "data": true
}`

const Universal_failed = `{
  "code": 503,
  "message": "failed",
  "data": false
}`

func getslot(h int) (string, error) {
	if h < 14 && h >= 0 {
		return "am", nil
	} else if h >= 14 && h <= 23 {
		return "pm", nil
	} else {
		return "", fmt.Errorf("wrong time format")
	}
}

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
