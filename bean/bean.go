package bean

type Report struct {
	TaskId     string   `json:"task_id"`
	TaskName   string   `json:"task_name"`
	TaskTime   int64    `json:"task_time"`
	TaskImages []string `json:"task_images"`
}

type Img struct {
	ImageId     int    `json:"image_id"`
	ImageUrl    string `json:"image_url"`
	ImageStatus string `json:"image_status"`
}

type TaskWithImg struct {
	TaskId     string `json:"task_id"`
	TaskName   string `json:"task_name"`
	TaskSlot   string `json:"task_slot"`
	TaskStatus string `json:"task_status"`
	TaskImages []*Img `json:"task_images"`
}

type TodayData struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []*TaskWithImg `json:"data"`
}

type Task struct {
	TaskId     string `json:"task_id"`
	TaskName   string `json:"task_name"`
	TaskStatus string `json:"task_status"`
}

type DateWithTask struct {
	DateTime string  `json:"date_time"`
	TaskList []*Task `json:"task_list"`
}

type TaskList struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []*DateWithTask `json:"data"`
}

//type NewStatus struct {
//	TaskId     string `json:"task_id"`
//	TaskName   string `json:"task_name"`
//	TaskSlot   string `json:"task_slot"`
//	TaskStatus string `json:"task_status"`
//	TaskImages []struct {
//		ImageId     int    `json:"image_id"`
//		ImageUrl    string `json:"image_url"`
//		ImageStatus string `json:"image_status"`
//	} `json:"task_images"`
//}
