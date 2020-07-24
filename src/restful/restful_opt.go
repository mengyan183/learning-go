package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello", hello)
	router.POST("/hello/:name", hello)
	_ = http.ListenAndServe(":8080", router)
}
