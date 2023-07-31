package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"urlServer/initconfig"
)

func main() {

	flag.Parse()
	initconfig.FinishInit(flag.Args()[0])

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:         initconfig.Get("listenurl") + ":" + initconfig.Get("listenport"),
		Handler:      mux,
		ReadTimeout:  time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe() failed\n", err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintln(writer, request.RemoteAddr)
	if err != nil {
		log.Print(err)
		return
	}
}

/*
1、task_slot时间范围
2、今日任务计划时间范围
3、状态变更仅更新当日数据？
*/
