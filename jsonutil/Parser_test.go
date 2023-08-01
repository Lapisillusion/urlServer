package jsonutil

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	//f, _ := os.Open("./test.json")
	//bytes, _ := io.ReadAll(f)
	//
	//var today bean.TodayData
	//json.Unmarshal(bytes, &today)
	//fmt.Println(today.Data[0].TaskImages[0].ImageUrl)
	//
	//marshal, err := json.Marshal(today)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(string(marshal))
	//
	//today.Data = append(today.Data, &bean.TaskWithImg{TaskId: "123"})
	//fmt.Println(today.Data[0])
	//
	//for i := 0; i < 10; i++ {
	//	var t int
	//	fmt.Println(&t)
	//}

	//m := make(map[string][]*bean.Task)
	//var t1 = bean.Task{TaskId: "123"}
	//var t2 = bean.Task{TaskId: "456"}
	//var t3 = bean.Task{TaskId: "789"}
	//
	//m["1"] = append(m["1"], &t1)
	//m["1"] = append(m["1"], &t2)
	//m["1"] = append(m["1"], &t3)
	//
	//_, ok := m["1"]
	//
	//fmt.Println(ok)
	//
	//m["1"] = make([]*bean.Task, 0, 4)
	//_, ok = m["1"]
	//
	//fmt.Println(ok)

	//var s []int
	var m = make(map[string][]int)
	//fmt.Println(cap(s))
	//m["1"] = s
	_, ok := m["1"]
	fmt.Println(ok)

	m["1"] = append(m["1"], 1, 2)

}
