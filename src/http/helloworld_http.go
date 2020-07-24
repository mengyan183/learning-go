package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "helloWorld")
	})
	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		fmt.Fprint(writer, timeStr, "  leafNode")
	})
	http.HandleFunc("/time/", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		fmt.Fprint(writer, timeStr, "  leafTree")
	})

	http.HandleFunc("/sub/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"  leafTree sub")
	})

	http.HandleFunc("/sub/node", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"  leafNode sub")
	})
	// 启动并监听8080端口,默认访问域名为localhost
	http.ListenAndServe(":8080", nil)
}
