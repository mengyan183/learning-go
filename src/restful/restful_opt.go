package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	_ "net/http/pprof" // 引入http pprof
)

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello", hello)
	router.POST("/hello/:name", hello)
	// 通过设置pprof handler修改router 不使用httprouter的规则
	router.Handler(http.MethodGet, "/debug/pprof/*item", http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", router))
}
