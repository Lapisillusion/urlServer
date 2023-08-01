package jsonutil

import (
	"fmt"
	"testing"
	"urlServer/bean"
)

func TestParser(t *testing.T) {
	//f, _ := os.Open("./test.json")
	//bytes, _ := io.ReadAll(f)
	//
	var today bean.TodayData
	//json.Unmarshal(bytes, &today)
	//fmt.Println(today.Data[0].TaskImages[0].ImageUrl)
	//
	//marshal, err := json.Marshal(today)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(string(marshal))

	today.Data = append(today.Data, &bean.TaskWithImg{TaskId: "123"})
	fmt.Println(today.Data[0])

	for i := 0; i < 10; i++ {
		var t int
		fmt.Println(&t)
	}
}
