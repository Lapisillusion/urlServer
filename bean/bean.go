package bean

type UrlEntry struct {
	Url    string `json:"url"`
	Status int32  `json:"status"`
}

type UrlList []UrlEntry

type TaskEntry struct {
	Task_id   string  `json:"task_id"`
	Task_name string  `json:"task_name"`
	Url_list  UrlList `json:"url_list"`
}

type DataEntry struct {
	App_id    string      `json:"app_id"`
	Task_list []TaskEntry `json:"task_list"`
}

type Pack struct {
	Data []DataEntry `json:"data"`
	Msg  string      `json:"msg"`
}
